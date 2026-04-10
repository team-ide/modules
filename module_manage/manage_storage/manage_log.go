// 文件由 TeamIDE | coos 生成，请勿修改文件内容！通过 [TeamIDE:teamide@163.com] 的 [models:] 在 [2026-04-10 16:15] 生成

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
	r, err := this_.dbService.Insert(context.Background(), in)
	if err != nil {
		return
	}
	res, err = r.RowsAffected()
	return
}

func (this_ *ManageLogStorage) DeleteByIds(ids []int64) (res int64, err error) {
	m := this_.dbService.SqlDelete("manage_log")
	m.Where().In("log_id", ids)
	r, err := m.Exec(context.Background())
	if err != nil {
		return
	}
	res, err = r.RowsAffected()
	return
}

func (this_ *ManageLogStorage) Query(in *module_manage.ManageLog) (res []*module_manage.ManageLog, err error) {
	m := this_.dbService.ModelSelect(in)
	res, err = db.DoQueryListWithModel[*module_manage.ManageLog](context.Background(), m)
	return
}

func (this_ *ManageLogStorage) Page(in *module_manage.ManageLog, pageNo int64, pageSize int64) (res []*module_manage.ManageLog, err error) {
	m := this_.dbService.ModelSelect(in)
	res, err = db.DoQueryPageWithModel[*module_manage.ManageLog](context.Background(), m, pageNo, pageSize)
	return
}

func (this_ *ManageLogStorage) Count(in *module_manage.ManageLog) (res int64, err error) {
	m := this_.dbService.ModelCount(in)
	m.Where().AndWhereSql("status != 9")
	res, err = m.Count(context.Background())
	return
}
