// 文件由 TeamIDE | coos 生成，请勿修改文件内容！通过 [TeamIDE:teamide@163.com] 的 [models:] 在 [2026-04-10 16:15] 生成

package manage_install

import (
	"github.com/team-ide/framework/db"
)

func TableManageLogCreate(dbService db.IService) (err error) {
	var version = "v1.0.0"
	var tableName = "manage_log"
	var moduleName = "manage"
	// 定义表结构
	table := &db.Table{
		Name:    tableName,
		Comment: "管理后台 日志 记录",
		ColumnList: []*db.Column{
			{Name: "log_id", DataType: "bigint", Length: 20, NotNull: true, Key: true, Comment: "日志id"},
			{Name: "login_id", DataType: "bigint", Length: 20, Default: "null", Comment: "登录id"},
			{Name: "user_id", DataType: "bigint", Length: 20, Default: "null", Comment: "用户id"},
			{Name: "user_name", DataType: "varchar", Length: 50, Default: "null", Comment: "用户名称"},
			{Name: "user_account", DataType: "varchar", Length: 50, Default: "null", Comment: "用户账号"},
			{Name: "ip", DataType: "varchar", Length: 50, Default: "null", Comment: "ip"},
			{Name: "path", DataType: "varchar", Length: 100, Default: "null", Comment: "请求"},
			{Name: "comment", DataType: "varchar", Length: 100, Comment: "备注"},
			{Name: "method", DataType: "varchar", Length: 20, Default: "null", Comment: "方法"},
			{Name: "param", DataType: "text", Default: "null", Comment: "参数"},
			{Name: "data", DataType: "text", Default: "null", Comment: "数据"},
			{Name: "user_agent", DataType: "text", Default: "null", Comment: "user-agent"},
			{Name: "error", DataType: "varchar", Length: 200, Default: "null", Comment: "异常"},
			{Name: "use_time", DataType: "int", Length: 10, Default: "0", Comment: "使用时长"},
			{Name: "start_at", DataType: "bigint", Length: 20, Default: "null", Comment: "开始时间"},
			{Name: "end_at", DataType: "bigint", Length: 20, Default: "null", Comment: "结束时间"},
			{Name: "create_at", DataType: "bigint", Length: 20, NotNull: true, Comment: "创建时间"},
		},
		IndexList: []*db.Index{
			{Name: "manage_log_login_id", ColumnNames: []string{"login_id"}},
			{Name: "manage_log_user_id", ColumnNames: []string{"user_id"}},
			{Name: "manage_log_user_name", ColumnNames: []string{"user_name"}},
			{Name: "manage_log_user_account", ColumnNames: []string{"user_account"}},
			{Name: "manage_log_ip", ColumnNames: []string{"ip"}},
			{Name: "manage_log_path", ColumnNames: []string{"path"}},
			{Name: "manage_log_comment", ColumnNames: []string{"comment"}},
			{Name: "manage_log_use_time", ColumnNames: []string{"use_time"}},
			{Name: "manage_log_start_at", ColumnNames: []string{"start_at"}},
			{Name: "manage_log_end_at", ColumnNames: []string{"end_at"}},
			{Name: "manage_log_create_at", ColumnNames: []string{"create_at"}},
		},
	}
	err = db.TableCreate(dbService, moduleName, version, tableName, table)
	return
}
