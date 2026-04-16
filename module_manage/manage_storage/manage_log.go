// 文件由 TeamIDE | coos 生成，请勿修改文件内容！通过 [TeamIDE:teamide@163.com] 的 [models:] 在 [2026-04-16 16:25] 生成

package manage_storage

import (
	"context"
	"time"

	"github.com/team-ide/framework/db"
	"github.com/team-ide/modules/module_manage"
)

func NewManageLogStorage(dbService db.IService) *ManageLogStorage {
	res := &ManageLogStorage{}
	res.dbService = dbService
	return res
}

type ManageLogStorage struct {
	dbService db.IService
}

func (this_ *ManageLogStorage) Insert(in *module_manage.ManageLog) (res int64, err error) {
	in.CreateAt = time.Now().UnixMilli()
	r, err := this_.dbService.Insert(context.Background(), in, func(in *db.ModelInsert) {
	})
	if err != nil {
		return
	}
	res, err = r.RowsAffected()
	return
}

func (this_ *ManageLogStorage) DeleteByIds(ids []int64) (res int64, err error) {
	m := this_.dbService.SqlDelete("manage_log", func(in *db.SqlDelete) {
	})
	m.Where().In("log_id", ids)
	r, err := m.Exec(context.Background())
	if err != nil {
		return
	}
	res, err = r.RowsAffected()
	return
}

func (this_ *ManageLogStorage) Query(in *module_manage.ManageLog) (res []*module_manage.ManageLog, err error) {
	m := this_.dbService.ModelSelect(in, func(in *db.ModelSelect) {
		in.CanSelectAll()
		in.WhereInclude("log_id", "login_id", "user_id", "user_name", "user_account", "ip", "path")
	})
	if in.StartTimeBefore > 0 {
		m.Where().AndWhereSql("start_at >= ?", in.StartTimeBefore)
	}
	if in.StartTimeAfter > 0 {
		m.Where().AndWhereSql("start_at <= ?", in.StartTimeAfter)
	}
	res, err = db.DoQueryListWithModel[*module_manage.ManageLog](context.Background(), m)
	return
}

func (this_ *ManageLogStorage) Page(in *module_manage.ManageLog, pageNo int64, pageSize int64) (res []*module_manage.ManageLog, err error) {
	m := this_.dbService.ModelSelect(in, func(in *db.ModelSelect) {
		in.CanSelectAll()
		in.WhereInclude("log_id", "login_id", "user_id", "user_name", "user_account", "ip", "path")
	})
	if in.StartTimeBefore > 0 {
		m.Where().AndWhereSql("start_at >= ?", in.StartTimeBefore)
	}
	if in.StartTimeAfter > 0 {
		m.Where().AndWhereSql("start_at <= ?", in.StartTimeAfter)
	}
	m.AppendSql("order by start_at desc")
	res, err = db.DoQueryPageWithModel[*module_manage.ManageLog](context.Background(), m, pageNo, pageSize)
	return
}

func (this_ *ManageLogStorage) Count(in *module_manage.ManageLog) (res int64, err error) {
	m := this_.dbService.ModelCount(in, func(in *db.ModelCount) {
		in.WhereInclude("log_id", "login_id", "user_id", "user_name", "user_account", "ip", "path")
	})
	if in.StartTimeBefore > 0 {
		m.Where().AndWhereSql("start_at >= ?", in.StartTimeBefore)
	}
	if in.StartTimeAfter > 0 {
		m.Where().AndWhereSql("start_at <= ?", in.StartTimeAfter)
	}
	res, err = m.Count(context.Background())
	return
}
