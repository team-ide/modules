package install

import (
	"errors"
	"github.com/team-ide/framework"
	"github.com/team-ide/framework/db"
)

func V100TbTest1Create(dbService db.IService) (err error) {
	version := "v1.0.0"
	tableName := "tb_test1"
	framework.Info("install version [" + version + "] [" + tableName + "] start")

	var tableCheckExists = dbService.TableCheckExists(dbService.GetSqlConn(), "", "", tableName)
	if !tableCheckExists {
		framework.Info("table [" + tableName + "] create start")
		table := &db.Table{
			Name:    tableName,
			Comment: "管理后台 用户",
			ColumnList: []*db.Column{
				{Name: "user_id", DataType: "bigint", Length: 20, NotNull: true, Key: true, Comment: "用户 ID",
					AutoIncrement: true, AutoIncrementCreate: true, AutoIncrementStart: 1000, AutoIncrementName: "seq_" + tableName + "_id"},
				{Name: "user_name", DataType: "varchar", Length: 100, Comment: "名称"},
				{Name: "account", DataType: "varchar", Length: 100, Comment: "登录账号"},
				{Name: "status", DataType: "int", Length: 10, Default: "0", Comment: "状态 1：正常 2 禁用 9：删除"},
				{Name: "extend", DataType: "varchar", Length: 2000, Comment: "扩展信息，json 格式存储"},
				{Name: "create_at", DataType: "bigint", Length: 20, Comment: "创建 时间戳 毫秒"},
				{Name: "update_at", DataType: "bigint", Length: 20, Comment: "修改 时间戳 毫秒"},
				{Name: "delete_at", DataType: "bigint", Length: 20, Comment: "删除 时间戳 毫秒"},
			},
			IndexList: []*db.Index{
				{Name: "idx_" + tableName + "_account", ColumnNames: []string{"account"}},
			},
		}
		err = dbService.TableCreate(dbService.GetSqlConn(), dbService.GetDDLHandler(), "", "", table)
		if err != nil {
			err = errors.New("install version [" + version + "] [" + tableName + "] create table error:" + err.Error())
			framework.Error(err.Error())
			return
		}
		framework.Info("table [" + tableName + "] create end")

	}

	framework.Info("install version [" + version + "] [tableName] end")

	return
}
