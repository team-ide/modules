// 文件由 TeamIDE | coos 生成，请勿修改文件内容！通过 [TeamIDE:teamide@163.com] 的 [models:] 在 [2026-03-18 10:47] 生成

package manage_install

import (
	"github.com/team-ide/framework/db"
)

func TableManageUserCreate(dbService db.IService) (err error) {
	var version = "v1.0.0"
	var tableName = "manage_user"
	var moduleName = "manage"
	// 定义表结构
	table := &db.Table{
		Name:    tableName,
		Comment: "管理后台 用户",
		ColumnList: []*db.Column{
			{Name: "user_id", DataType: "bigint", Length: 20, NotNull: true, Key: true, Comment: "用户 ID"},
			{Name: "name", DataType: "varchar", Length: 100, Comment: "名称"},
			{Name: "account", DataType: "varchar", Length: 100, Comment: "登录账号"},
			{Name: "salt", DataType: "varchar", Length: 50, Comment: "密码 盐 设置密码时候 自动生成"},
			{Name: "password", DataType: "varchar", Length: 100, Comment: "密码  md5(salt + md5(password))"},
			{Name: "status", DataType: "int", Length: 10, Default: "1", Comment: "状态 1：正常 2：禁用 9：删除"},
			{Name: "create_at", DataType: "bigint", Length: 20, Comment: "创建 时间戳 毫秒"},
			{Name: "update_at", DataType: "bigint", Length: 20, Comment: "修改 时间戳 毫秒"},
			{Name: "delete_at", DataType: "bigint", Length: 20, Comment: "删除 时间戳 毫秒"},
			{Name: "disable_at", DataType: "bigint", Length: 20, Comment: "禁用 时间戳 毫秒"},
		},
		IndexList: []*db.Index{
			{Name: "", ColumnNames: []string{"name"}, Comment: "用于 根据 名称 查询"},
			{Name: "", ColumnNames: []string{"account"}, Comment: "用于 根据 账号 查询"},
			{Name: "", ColumnNames: []string{"create_at"}},
			{Name: "", ColumnNames: []string{"update_at"}},
			{Name: "", ColumnNames: []string{"delete_at"}},
			{Name: "", ColumnNames: []string{"disable_at"}},
		},
	}
	err = db.TableCreate(dbService, moduleName, version, tableName, table)
	return
}

func TableManageUserAddColumnXxx(dbService db.IService) (err error) {
	var version = "v1.0.0"
	var tableName = "manage_user"
	var moduleName = "manage"

	column := &db.Column{Name: "xxx", DataType: "varchar", Length: 100, Comment: "名称"}

	err = db.TableColumnAdd(dbService, moduleName, version, tableName, column)
	return
}

func TableManageUserAddIndexXxx(dbService db.IService) (err error) {
	var version = "v1.0.0"
	var tableName = "manage_user"
	var moduleName = "manage"

	index := &db.Index{Name: "xxx", ColumnNames: []string{"x", "xx"}, Comment: "名称"}

	err = db.TableIndexAdd(dbService, moduleName, version, tableName, index)
	return
}

func TableManageUserDropColumnXxx(dbService db.IService) (err error) {
	var version = "v1.0.0"
	var tableName = "manage_user"
	var moduleName = "manage"

	err = db.TableColumnDrop(dbService, moduleName, version, tableName, "xxx")
	return
}

func TableManageUserDropIndexXxx(dbService db.IService) (err error) {
	var version = "v1.0.0"
	var tableName = "manage_user"
	var moduleName = "manage"

	err = db.TableIndexDrop(dbService, moduleName, version, tableName, "xxx")
	return
}
