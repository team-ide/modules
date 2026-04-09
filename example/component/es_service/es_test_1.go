package es_service

import "github.com/team-ide/framework/es"

var (
	esTest1 es.IService
)

func InitEsTest1(name string, cfg *es.Config) (err error) {
	// 如果 已存在 则 直接 返回
	if esTest1 != nil {
		return
	}
	// 创建 服务
	esTest1, err = NewEsService(name, cfg)
	if err != nil {
		return
	}

	return
}

func GetEsTest1() (res es.IService) {
	return esTest1
}
