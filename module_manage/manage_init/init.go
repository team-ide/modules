package manage_init

import (
	"github.com/team-ide/framework/db"
	"github.com/team-ide/modules/module_manage/manage_factory"
	"github.com/team-ide/modules/module_manage/manage_service"
	"github.com/team-ide/modules/module_manage/manage_storage"
)

func InitFactoryService(dbService db.IService) {
	manage_factory.LoginCache = manage_service.NewLoginCache()

	manage_factory.ManageId = manage_service.NewManageId()

	manage_factory.ManageService = manage_service.NewManageService()

	manage_factory.ManageUserStorage = manage_storage.NewManageUserStorage(dbService)

	manage_factory.ManageRoleStorage = manage_storage.NewManageRoleStorage(dbService)
	manage_factory.ManageRoleUserStorage = manage_storage.NewManageRoleUserStorage(dbService)

	manage_factory.ManagePermissionStorage = manage_storage.NewManagePermissionStorage(dbService)

	manage_factory.ManageLoginStorage = manage_storage.NewManageLoginStorage(dbService)

	manage_factory.ManageLogStorage = manage_storage.NewManageLogStorage(dbService)

	manage_factory.ManageUserService = manage_service.NewManageUserService()
	manage_factory.ManageRoleService = manage_service.NewManageRoleService()
	manage_factory.ManagePermissionService = manage_service.NewManagePermissionService()
	manage_factory.ManageLoginService = manage_service.NewManageLoginService()
	manage_factory.ManageLogService = manage_service.NewManageLogService()
	return
}
