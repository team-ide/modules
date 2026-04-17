package start

import (
	"example/component"
	"example/component/db_service"
	"github.com/team-ide/framework"

	"github.com/team-ide/modules/module_id"
	"github.com/team-ide/modules/module_id/id_single"
	"github.com/team-ide/modules/module_manage/manage_init"
)

func init() {
	component.Starter.AddInitFactoryFunc("init module mange factory", 200, InitModuleManageFactory)
}

func InitModuleManageFactory() (err error) {
	framework.Info("init module mange factory start")

	idHandler := id_single.NewIdSingleHandler(db_service.GetDbTest1())

	module_id.IdWorker = module_id.NewIdGenerator(idHandler)

	manage_init.InitFactoryService(db_service.GetDbTest1())

	framework.Info("init module mange factory success")
	return
}
