// 文件由 TeamIDE | coos 生成，请勿修改文件内容！通过 [TeamIDE:teamide@163.com] 的 [models:] 在 [2026-03-17 15:38] 生成

package manage_install

import (
	"github.com/team-ide/framework/db"
)

func InstallTable(dbService db.IService) (err error) {
	if err = TableManageLoginCreate(dbService); err != nil {
		return
	}
	if err = TableManageRoleCreate(dbService); err != nil {
		return
	}
	if err = TableManageRolePermissionCreate(dbService); err != nil {
		return
	}
	if err = TableManageRoleUserCreate(dbService); err != nil {
		return
	}
	if err = TableManageUserCreate(dbService); err != nil {
		return
	}
	if err = TableManageUserAddColumnXxx(dbService); err != nil {
		return
	}
	if err = TableManageUserAddIndexXxx(dbService); err != nil {
		return
	}
	if err = TableManageUserDropColumnXxx(dbService); err != nil {
		return
	}
	if err = TableManageUserDropIndexXxx(dbService); err != nil {
		return
	}
	return
}
