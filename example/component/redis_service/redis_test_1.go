package redis_service

import "github.com/team-ide/framework/redis"

var (
	redisTest1 redis.IService
)

func InitRedisTest1(name string, cfg *redis.Config) (err error) {
	// 如果 已存在 则 直接 返回
	if redisTest1 != nil {
		return
	}
	// 创建 服务
	redisTest1, err = NewRedisService(name, cfg)
	if err != nil {
		return
	}

	return
}

func GetRedisTest1() (res redis.IService) {
	return redisTest1
}
