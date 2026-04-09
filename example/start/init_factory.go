package start

import (
	"example/common"
	"example/component/db_service"
	"example/factory"
	"example/storage"

	"github.com/team-ide/framework"
)

func init() {
	common.Starter.AddInitFactoryFunc("init factory", 100, InitFactory)
	common.Starter.AddInitFactoryFunc("init factory storage", 100, InitFactoryStorage)
}

func InitFactory() (err error) {
	framework.Info("init factory start")

	framework.Info("init factory success")
	return
}

func InitFactoryStorage() (err error) {
	framework.Info("init factory storage start")

	factory.TbTest1Storage = storage.NewTbTest1Storage(db_service.GetDbTest1())

	framework.Info("init factory storage success")
	return
}
