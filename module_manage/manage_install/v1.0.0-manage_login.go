// 文件由 TeamIDE | coos 生成，请勿修改文件内容！通过 [TeamIDE:teamide@163.com] 的 [models:] 在 [2026-04-10 16:15] 生成

package manage_install

import (
	"github.com/team-ide/framework/db"
)

func TableManageLoginCreate(dbService db.IService) (err error) {
	var version = "v1.0.0"
	var tableName = "manage_login"
	var moduleName = "manage"
	// 定义表结构
	table := &db.Table{
		Name:    tableName,
		Comment: "管理后台 用户 登录记录",
		ColumnList: []*db.Column{
			{Name: "login_id", DataType: "bigint", Length: 20, NotNull: true, Key: true, Comment: "登录id"},
			{Name: "user_id", DataType: "bigint", Length: 20, Comment: "用户ID"},
			{Name: "user_account", DataType: "varchar", Length: 50, Comment: "用户 账号"},
			{Name: "user_name", DataType: "varchar", Length: 50, Comment: "用户 名称"},
			{Name: "login_ip", DataType: "varchar", Length: 50, Comment: "登录 ip"},
			{Name: "source_type", DataType: "int", Length: 10, Comment: "来源 类型"},
			{Name: "source_info", DataType: "varchar", Length: 100, Comment: "来源 信息"},
			{Name: "token", DataType: "varchar", Length: 100, Comment: "验证票据"},
			{Name: "status", DataType: "int", Length: 10, Default: "1", Comment: "状态 1：登录 2：登出 9：删除"},
			{Name: "login_at", DataType: "bigint", Length: 20, Comment: "登录 时间"},
			{Name: "logout_at", DataType: "bigint", Length: 20, Comment: "登出 时间"},
			{Name: "create_at", DataType: "bigint", Length: 20, Comment: "创建 时间"},
			{Name: "update_at", DataType: "bigint", Length: 20, Comment: "修改 时间"},
			{Name: "delete_at", DataType: "bigint", Length: 20, Comment: "删除 时间"},
			{Name: "use_at", DataType: "bigint", Length: 20, Comment: "使用 时间"},
		},
		IndexList: []*db.Index{
			{Name: "", ColumnNames: []string{"manage_login_user_id"}},
			{Name: "", ColumnNames: []string{"manage_login_user_account"}},
			{Name: "", ColumnNames: []string{"manage_login_user_name"}},
			{Name: "", ColumnNames: []string{"manage_login_source_type"}},
			{Name: "", ColumnNames: []string{"manage_login_token"}},
			{Name: "", ColumnNames: []string{"manage_login_create_at"}},
			{Name: "", ColumnNames: []string{"manage_login_logout_at"}},
			{Name: "", ColumnNames: []string{"manage_login_use_at"}},
		},
	}
	err = db.TableCreate(dbService, moduleName, version, tableName, table)
	return
}
