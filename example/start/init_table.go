package start

import (
	"example/component"
	"example/component/db_service"
	"example/storage/install"

	"github.com/team-ide/modules/module_id/id_storage"
	"github.com/team-ide/modules/module_manage/manage_install"
)

func init() {
	component.Starter.AddInitTableFunc("init module id table", 100, func() error {
		return id_storage.Install(db_service.GetDbTest1())
	})
	component.Starter.AddInitTableFunc("init module manage table", 100, func() error {
		return manage_install.InstallTable(db_service.GetDbTest1())
	})
	component.Starter.AddInitTableFunc("init table v100 tb_test1", 100, func() error {
		return install.V100TbTest1Create(db_service.GetDbTest1())
	})
}
