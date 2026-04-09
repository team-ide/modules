package zk_service

import (
	"github.com/team-ide/framework/zookeeper"
)

var (
	zkTest1 zookeeper.IService
)

func InitZkTest1(name string, cfg *zookeeper.Config) (err error) {
	// 如果 已存在 则 直接 返回
	if zkTest1 != nil {
		return
	}
	// 创建 服务
	zkTest1, err = NewZkService(name, cfg)
	if err != nil {
		return
	}

	return
}

func GetZkTest1() (res zookeeper.IService) {
	return zkTest1
}
