// 文件由 TeamIDE | coos 生成，请勿修改文件内容！通过 [TeamIDE:teamide@163.com] 的 [models:] 在 [2026-04-10 10:27] 生成

package manage_install

import (
	"github.com/team-ide/framework/db"
)

func TableManageRoleCreate(dbService db.IService) (err error) {
	var version = "v1.0.0"
	var tableName = "manage_role"
	var moduleName = "manage"
	// 定义表结构
	table := &db.Table{
		Name:    tableName,
		Comment: "管理后台 角色",
		ColumnList: []*db.Column{
			{Name: "role_id", DataType: "bigint", Length: 20, NotNull: true, Key: true, Comment: "角色 ID"},
			{Name: "name", DataType: "varchar", Length: 100, Comment: "名称"},
			{Name: "is_super", DataType: "tinyint", Length: 3, Comment: "角色 是否 是超管 如果是 则拥有所有权限 1是 2否"},
			{Name: "create_at", DataType: "bigint", Length: 20, Comment: "创建 时间戳 毫秒"},
			{Name: "update_at", DataType: "bigint", Length: 20, Comment: "修改 时间戳 毫秒"},
		},
		IndexList: []*db.Index{},
	}
	err = db.TableCreate(dbService, moduleName, version, tableName, table)
	return
}

func TableManageRoleUserCreate(dbService db.IService) (err error) {
	var version = "v1.0.0"
	var tableName = "manage_role_user"
	var moduleName = "manage"
	// 定义表结构
	table := &db.Table{
		Name:    tableName,
		Comment: "管理后台 角色 用户",
		ColumnList: []*db.Column{
			{Name: "role_user_id", DataType: "bigint", Length: 20, NotNull: true, Key: true, Comment: "角色 用户 ID"},
			{Name: "role_id", DataType: "bigint", Length: 20, NotNull: true, Comment: "角色 ID"},
			{Name: "user_id", DataType: "bigint", Length: 20, NotNull: true, Comment: "用户 ID"},
			{Name: "create_at", DataType: "bigint", Length: 20, Comment: "创建 时间戳 毫秒"},
			{Name: "update_at", DataType: "bigint", Length: 20, Comment: "修改 时间戳 毫秒"},
		},
		IndexList: []*db.Index{
			{Name: "", ColumnNames: []string{"role_id"}, Comment: "用于 根据 角色 查询"},
			{Name: "", ColumnNames: []string{"user_id"}, Comment: "用于 根据 用户 查询"},
		},
	}
	err = db.TableCreate(dbService, moduleName, version, tableName, table)
	return
}
