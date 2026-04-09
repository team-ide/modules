package mongodb_service

import "github.com/team-ide/framework/mongodb"

var (
	mongodbTest1 mongodb.IService
)

func InitMongodbTest1(name string, cfg *mongodb.Config) (err error) {
	// 如果 已存在 则 直接 返回
	if mongodbTest1 != nil {
		return
	}
	// 创建 服务
	mongodbTest1, err = NewMongodbService(name, cfg)
	if err != nil {
		return
	}

	return
}

func GetMongodbTest1() (res mongodb.IService) {
	return mongodbTest1
}
