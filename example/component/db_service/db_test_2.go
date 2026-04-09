package db_service

import (
	"github.com/team-ide/framework/db"
)

// 引入 支持的数据库 驱动
import _ "github.com/team-ide/framework/driver/mysql"
import _ "github.com/team-ide/framework/driver/sqlite_noc"

// 引入 支持的数据库 方言
import _ "github.com/team-ide/framework/dialects/dialect_mysql"
import _ "github.com/team-ide/framework/dialects/dialect_sqlite"

var (
	dbTest2 db.IService
)

func InitDbTest2(name string, cfg *db.Config) (err error) {
	// 如果 已存在 则 直接 返回
	if dbTest2 != nil {
		return
	}
	// 创建 服务
	dbTest2, err = NewDbService(name, cfg)
	if err != nil {
		return
	}

	return
}

func GetDbTest2() (res db.IService) {
	return dbTest2
}
