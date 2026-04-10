// 文件由 TeamIDE | coos 生成，请勿修改文件内容！通过 [TeamIDE:teamide@163.com] 的 [models:] 在 [2026-04-10 10:27] 生成

package manage_install

import (
	"github.com/team-ide/framework/db"
)

func TableManagePermissionCreate(dbService db.IService) (err error) {
	var version = "v1.0.0"
	var tableName = "manage_permission"
	var moduleName = "manage"
	// 定义表结构
	table := &db.Table{
		Name:    tableName,
		Comment: "管理后台 权限",
		ColumnList: []*db.Column{
			{Name: "permission_id", DataType: "bigint", Length: 20, NotNull: true, Key: true, Comment: "角色 权限 ID"},
			{Name: "role_id", DataType: "bigint", Length: 20, NotNull: true, Comment: "角色 ID 给角色授权"},
			{Name: "user_id", DataType: "bigint", Length: 20, NotNull: true, Comment: "用户 ID 给用户授权"},
			{Name: "permission_type", DataType: "varchar", Length: 100, Comment: "权限类型 path: 路径 button: 按钮 tag: 标签 page: 页面"},
			{Name: "permission", DataType: "varchar", Length: 1000, Comment: "权限 根据类型设置 如：/admin/user/add"},
			{Name: "authorizable", DataType: "int", Length: 10, Comment: "可授权 2:不可授权 1:可以授权"},
			{Name: "start_at", DataType: "bigint", Length: 20, Comment: "权限开始 时间戳 毫秒"},
			{Name: "expired_at", DataType: "bigint", Length: 20, Comment: "过期 时间戳 毫秒"},
			{Name: "expired_duration", DataType: "bigint", Length: 20, Comment: "过期时间 分钟"},
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
