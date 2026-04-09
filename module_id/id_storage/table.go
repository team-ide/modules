package id_storage

import (
	"github.com/team-ide/framework/db"
)

var (
	ModuleName    = "module_id"
	TableNameTbId = "tb_id"
)

var (
	TableTbId = &db.Table{
		Name:    TableNameTbId,
		Comment: "id 存储",
		ColumnList: []*db.Column{
			{Name: "id_type", DataType: "varchar", Length: 100, NotNull: true, Key: true, Comment: "id 类型"},
			{Name: "id_value", DataType: "bigint", Length: 20, NotNull: true, Comment: "id 值"},
			{Name: "id_comment", DataType: "varchar", Length: 200, Comment: "id 注释"},
			{Name: "create_at", DataType: "bigint", Length: 20, Comment: "创建 时间戳 毫秒"},
			{Name: "update_at", DataType: "bigint", Length: 20, Comment: "修改 时间戳 毫秒"},
		},
		IndexList: []*db.Index{},
	}
)

func Install(dbService db.IService) (err error) {
	err = V100TbIdCreate(dbService)
	return
}

func V100TbIdCreate(dbService db.IService) (err error) {
	version := "v1.0.0"
	err = db.TableCreate(dbService, ModuleName, version, TableNameTbId, TableTbId)
	return
}
