package id_single

import (
	"context"
	"fmt"
	"sync"

	"github.com/team-ide/framework/db"
	"github.com/team-ide/modules/module_id"
	"github.com/team-ide/modules/module_id/id_storage"
)

func NewIdSingleHandler(dbService db.IService) (res *IdSingleHandler) {
	res = &IdSingleHandler{}
	res.idLockerCache = make(map[module_id.IdType]sync.Locker)
	res.idLockerCacheLocker = &sync.Mutex{}
	res.recoverLockerCache = make(map[module_id.IdType]sync.Locker)
	res.recoverLockerCacheLocker = &sync.Mutex{}
	res.persistLockerCache = make(map[module_id.IdType]sync.Locker)
	res.persistLockerCacheLocker = &sync.Mutex{}
	res.IdStorage = &id_storage.IdStorage{
		DbService: dbService,
	}
	return
}

type IdSingleHandler struct {
	*id_storage.IdStorage
	idLockerCache            map[module_id.IdType]sync.Locker
	idLockerCacheLocker      sync.Locker
	recoverLockerCache       map[module_id.IdType]sync.Locker
	recoverLockerCacheLocker sync.Locker
	persistLockerCache       map[module_id.IdType]sync.Locker
	persistLockerCacheLocker sync.Locker
}

func (this_ *IdSingleHandler) ShouldCheckAndPersist() bool {
	return false
}
func (this_ *IdSingleHandler) Increment(ctx context.Context, idType module_id.IdType, size int64) (int64, error) {
	locker := this_.getIdLocker(idType)
	locker.Lock()
	defer locker.Unlock()

	queryId, err := this_.StorageQueryId(ctx, idType)
	if err != nil {
		return 0, err
	}
	if queryId == 0 {
		return 0, nil
	}

	value := queryId + size
	err = this_.StorageUpdate(ctx, idType, value)
	if err != nil {
		return 0, err
	}
	return value, nil
}
func (this_ *IdSingleHandler) CacheId(ctx context.Context, idType module_id.IdType, id int64) error {
	if err := this_.StorageUpdate(ctx, idType, id); err != nil {
		return fmt.Errorf("storage insert id failed: %w", err)
	}
	return nil
}
func (this_ *IdSingleHandler) RecoverLock(ctx context.Context, idType module_id.IdType) (unlocker module_id.Unlocker, err error) {
	locker := this_.getRecoverLocker(idType)
	locker.Lock()
	unlocker = locker
	return
}
func (this_ *IdSingleHandler) PersistLock(ctx context.Context, idType module_id.IdType) (unlocker module_id.Unlocker, err error) {
	locker := this_.getPersistLocker(idType)
	locker.Lock()
	unlocker = locker
	return
}

func (this_ *IdSingleHandler) getIdLocker(idType module_id.IdType) sync.Locker {
	this_.idLockerCacheLocker.Lock()
	defer this_.idLockerCacheLocker.Unlock()
	find, ok := this_.idLockerCache[idType]
	if !ok {
		find = &sync.Mutex{}
		this_.idLockerCache[idType] = find
	}
	return find
}

func (this_ *IdSingleHandler) getRecoverLocker(idType module_id.IdType) sync.Locker {
	this_.recoverLockerCacheLocker.Lock()
	defer this_.recoverLockerCacheLocker.Unlock()
	find, ok := this_.recoverLockerCache[idType]
	if !ok {
		find = &sync.Mutex{}
		this_.recoverLockerCache[idType] = find
	}
	return find
}
func (this_ *IdSingleHandler) getPersistLocker(idType module_id.IdType) sync.Locker {
	this_.persistLockerCacheLocker.Lock()
	defer this_.persistLockerCacheLocker.Unlock()
	find, ok := this_.persistLockerCache[idType]
	if !ok {
		find = &sync.Mutex{}
		this_.persistLockerCache[idType] = find
	}
	return find
}
