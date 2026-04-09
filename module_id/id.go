package module_id

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

var (
	// IdWorker 全局的ID生成器 请在 程序初始化 时候赋值
	IdWorker IdFace
)

type IdType string

type IdFace interface {
	GetId(ctx context.Context, idType IdType) (int64, error)
	GetIds(ctx context.Context, idType IdType, size int64) ([]int64, error)
}

type IdHandler interface {
	ShouldCheckAndPersist() bool
	// Increment 增长，使用 redis 的 incr 指令实现原子操作
	// 如果缓存不存在则返回的 id 为 0
	Increment(ctx context.Context, idType IdType, size int64) (int64, error)
	// CacheId 缓存 id，缓存中不存在则通过 QueryId 查询恢复
	CacheId(ctx context.Context, idType IdType, id int64) error
	// StorageQueryId 存储 查询 id，缓存中不存在时进行查询
	StorageQueryId(ctx context.Context, idType IdType) (int64, error)
	// StorageInsert 存储 新增 id
	StorageInsert(ctx context.Context, idType IdType, id int64) error
	// StorageUpdate 存储 更新 id
	StorageUpdate(ctx context.Context, idType IdType, id int64) error
	// RecoverLock id 恢复锁，通常在存储时使用
	RecoverLock(ctx context.Context, idType IdType) (Unlocker, error)
	// PersistLock id 入库落地锁，通常在存储时使用
	PersistLock(ctx context.Context, idType IdType) (Unlocker, error)
}

type Unlocker interface {
	Unlock()
}

var (
	DefaultIdStorageSpan int64 = 500
	DefaultStartId       int64 = 10000
)

type IdGenerator struct {
	handler       IdHandler
	idStorageSpan int64
	startId       int64
	startIds      map[IdType]int64
	mu            sync.RWMutex // 用于保护 startIds 的并发访问
}

func NewIdGenerator(handler IdHandler) *IdGenerator {
	return &IdGenerator{
		handler:       handler,
		idStorageSpan: DefaultIdStorageSpan,
		startId:       DefaultStartId,
		startIds:      make(map[IdType]int64),
	}
}

// WithIdStorageSpan 设置ID存储跨度
func (this_ *IdGenerator) WithIdStorageSpan(span int64) *IdGenerator {
	this_.idStorageSpan = span
	return this_
}

// WithStartId 设置默认起始ID
func (this_ *IdGenerator) WithStartId(startId int64) *IdGenerator {
	this_.startId = startId
	return this_
}

// SetStartIdForType 为特定类型设置起始ID
func (this_ *IdGenerator) SetStartIdForType(idType IdType, startId int64) {
	this_.mu.Lock()
	defer this_.mu.Unlock()
	this_.startIds[idType] = startId
}

func (this_ *IdGenerator) GetId(ctx context.Context, idType IdType) (int64, error) {
	ids, err := this_.GetIds(ctx, idType, 1)
	if err != nil {
		return 0, fmt.Errorf("get id failed: %w", err)
	}
	return ids[0], nil
}

func (this_ *IdGenerator) GetIds(ctx context.Context, idType IdType, size int64) ([]int64, error) {
	if size < 0 {
		return nil, errors.New("GetIds size must be > 0")
	}

	currentId, err := this_.handler.Increment(ctx, idType, size)
	if err != nil {
		return nil, fmt.Errorf("increment id failed: %w", err)
	}

	// 缓存中没有则通过数据库恢复
	if currentId == 0 {
		currentId, err = this_.recoverId(ctx, idType, size)
		if err != nil {
			return nil, fmt.Errorf("recover id failed: %w", err)
		}
	} else {
		// 检查是否需要持久化
		if this_.handler.ShouldCheckAndPersist() {
			if err = this_.checkAndPersist(ctx, idType, currentId, size); err != nil {
				return nil, fmt.Errorf("persist id failed: %w", err)
			}
		}
	}

	return this_.toIds(currentId, size), nil
}

func (this_ *IdGenerator) toIds(currentId int64, size int64) []int64 {
	ids := make([]int64, size)
	for i := size - 1; i >= 0; i-- {
		ids[i] = currentId - i
	}
	return ids
}

func (this_ *IdGenerator) checkAndPersist(ctx context.Context, idType IdType, currentId int64, size int64) error {
	oldId := currentId - size
	oldSpan := oldId / this_.idStorageSpan
	currentSpan := currentId / this_.idStorageSpan

	// 跨度不同时需要持久化
	if oldSpan < currentSpan {
		if err := this_.persistId(ctx, idType, currentId); err != nil {
			return fmt.Errorf("persist id failed: %w", err)
		}
	}
	return nil
}

func (this_ *IdGenerator) recoverId(ctx context.Context, idType IdType, size int64) (int64, error) {
	unlocker, err := this_.handler.RecoverLock(ctx, idType)
	if err != nil {
		return 0, fmt.Errorf("recoverId RecoverLock failed: %w", err)
	}
	defer unlocker.Unlock()

	id, err := this_.handler.StorageQueryId(ctx, idType)
	if err != nil {
		return 0, fmt.Errorf("recoverId QueryId failed: %w", err)
	}
	var find = id != 0
	if !find {
		this_.mu.RLock()
		startId, exists := this_.startIds[idType]
		this_.mu.RUnlock()

		if !exists {
			startId = this_.startId
		}
		id = startId
	}

	storageId := id + this_.idStorageSpan
	currentId := storageId + size
	if find {
		if err = this_.handler.StorageUpdate(ctx, idType, storageId); err != nil {
			return 0, fmt.Errorf("storage update id failed: %w", err)
		}
	} else {
		if err = this_.handler.StorageInsert(ctx, idType, storageId); err != nil {
			return 0, fmt.Errorf("storage insert id failed: %w", err)
		}
	}

	if err = this_.handler.CacheId(ctx, idType, currentId); err != nil {
		return 0, fmt.Errorf("cache id failed: %w", err)
	}

	return currentId, nil
}

func (this_ *IdGenerator) persistId(ctx context.Context, idType IdType, currentId int64) error {
	unlocker, err := this_.handler.PersistLock(ctx, idType)
	if err != nil {
		return fmt.Errorf("persistId PersistLock failed: %w", err)
	}
	defer unlocker.Unlock()

	storageId := (currentId/this_.idStorageSpan + 1) * this_.idStorageSpan
	if err = this_.handler.StorageUpdate(ctx, idType, storageId); err != nil {
		return fmt.Errorf("persistId StorageId failed: %w", err)
	}
	return nil
}
