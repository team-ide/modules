// 文件由 TeamIDE | coos 生成，请勿修改文件内容！通过 [TeamIDE:teamide@163.com] 的 [models:module_manage/service/manage_init.coos] 在 [2026-04-08 14:27] 生成

package module_manage

type IManageInitialize interface {
	InitManageUser(name string, account string) (userId int64, err error)
	InitSuperRole(name string) (roleId int64, err error)
}
