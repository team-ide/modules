package id_cluster

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/team-ide/framework"
	"github.com/team-ide/framework/db"
	"github.com/team-ide/framework/redis"
	"github.com/team-ide/modules/module_id"
	"github.com/team-ide/modules/module_id/id_storage"
)

func NewIdClusterHandler(dbService db.IService, redisService redis.IService) (res *IdClusterHandler) {
	res = &IdClusterHandler{}
	res.redisService = redisService
	res.redisLocker = redis.NewLocker(redisService)
	res.idKeyFormat = DefaultIdKeyFormat
	res.lockTimeout = DefaultIdLockTimeout
	res.IdStorage = &id_storage.IdStorage{
		DbService: dbService,
	}
	return
}

var (
	DefaultIdKeyFormat   = "id:type:%s"
	DefaultIdLockTimeout = time.Second * 3
)

type IdClusterHandler struct {
	*id_storage.IdStorage
	redisService redis.IService
	idKeyFormat  string
	lockTimeout  time.Duration
	redisLocker  *redis.Locker
}

func (this_ *IdClusterHandler) ShouldCheckAndPersist() bool {
	return true
}
func (this_ *IdClusterHandler) WithIdKey(idKeyFormat string) *IdClusterHandler {
	this_.idKeyFormat = idKeyFormat
	return this_
}
func (this_ *IdClusterHandler) WithLockTimeout(lockTimeout time.Duration) *IdClusterHandler {
	this_.lockTimeout = lockTimeout
	return this_
}
func (this_ *IdClusterHandler) FormatIdKey(idType module_id.IdType) string {
	return fmt.Sprintf(this_.idKeyFormat, idType)
}
func (this_ *IdClusterHandler) Increment(ctx context.Context, idType module_id.IdType, size int64) (int64, error) {
	idKey := this_.FormatIdKey(idType)
	return this_.redisService.Incr(idKey, size)
}
func (this_ *IdClusterHandler) CacheId(ctx context.Context, idType module_id.IdType, id int64) error {
	idKey := this_.FormatIdKey(idType)
	return this_.redisService.Set(idKey, strconv.FormatInt(id, 10))
}

func (this_ *IdClusterHandler) RecoverLock(ctx context.Context, idType module_id.IdType) (unlocker module_id.Unlocker, err error) {
	idKey := this_.FormatIdKey(idType)
	lockKey := idKey + ":recoverLock"
	redisUnlocker := &RedisUnlocker{}
	redisUnlocker.lockKey = lockKey
	redisUnlocker.timeout = this_.lockTimeout
	redisUnlocker.token, err = this_.redisLocker.Lock(ctx, lockKey, redisUnlocker.timeout)
	if err != nil {
		return
	}
	unlocker = redisUnlocker
	return
}
func (this_ *IdClusterHandler) PersistLock(ctx context.Context, idType module_id.IdType) (unlocker module_id.Unlocker, err error) {
	idKey := this_.FormatIdKey(idType)
	lockKey := idKey + ":persistLock"
	redisUnlocker := &RedisUnlocker{}
	redisUnlocker.lockKey = lockKey
	redisUnlocker.timeout = this_.lockTimeout
	redisUnlocker.token, err = this_.redisLocker.Lock(ctx, lockKey, redisUnlocker.timeout)
	if err != nil {
		return
	}
	unlocker = redisUnlocker
	return
}

type RedisUnlocker struct {
	lockKey     string
	token       string
	timeout     time.Duration
	ctx         context.Context
	redisLocker *redis.Locker
}

func (this_ *RedisUnlocker) Unlock() {
	err := this_.redisLocker.Unlock(this_.ctx, this_.lockKey, this_.token)
	if err != nil {
		framework.Error("redisUnlocker redisLocker.Unlock lockKey:" + this_.lockKey + " token:" + this_.token + " err:" + err.Error())
		return
	}
}
