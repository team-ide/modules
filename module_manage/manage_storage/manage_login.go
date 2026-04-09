// 文件由 TeamIDE | coos 生成，请勿修改文件内容！通过 [TeamIDE:teamide@163.com] 的 [models:] 在 [2026-04-09 11:11] 生成

package manage_storage

import (
	"context"
	"time"

	"github.com/team-ide/framework/db"
	"github.com/team-ide/modules/module_manage"
)

func NewManageLoginStorage(dbService db.IService) *ManageLoginStorage {
	res := &ManageLoginStorage{}
	res.dbService = dbService
	return res
}

type ManageLoginStorage struct {
	dbService db.IService
}

func (this_ *ManageLoginStorage) GetById(id int64) (res *module_manage.ManageLogin, err error) {
	m := this_.dbService.SqlSelect("manage_login")
	m.Where().Eq("login_id", id)
	res, err = db.DoQueryOneWithModel[*module_manage.ManageLogin](context.Background(), m)
	return
}

func (this_ *ManageLoginStorage) GetByToken(token string) (res *module_manage.ManageLogin, err error) {
	m := this_.dbService.SqlSelect("manage_login")
	m.Where().Eq("token", token)
	res, err = db.DoQueryOneWithModel[*module_manage.ManageLogin](context.Background(), m)
	return
}

func (this_ *ManageLoginStorage) Insert(in *module_manage.ManageLogin) (res int64, err error) {
	in.CreateAt = time.Now().UnixMilli()
	if in.Status == 0 {
		in.Status = 1
	}
	r, err := this_.dbService.Insert(context.Background(), in)
	if err != nil {
		return
	}
	res, err = r.RowsAffected()
	return
}

func (this_ *ManageLoginStorage) Update(in *module_manage.ManageLogin) (res int64, err error) {
	in.UpdateAt = time.Now().UnixMilli()
	r, err := this_.dbService.Update(context.Background(), in)
	if err != nil {
		return
	}
	res, err = r.RowsAffected()
	return
}

func (this_ *ManageLoginStorage) Logout(id int64) (res int64, err error) {
	m := this_.dbService.SqlUpdate("manage_login")
	m.Value("update_at", time.Now().UnixMilli())
	m.Value("status", 2)
	m.Value("logout_at", time.Now().UnixMilli())
	m.Where().Eq("login_id", id)
	r, err := m.Exec(context.Background())
	if err != nil {
		return
	}
	res, err = r.RowsAffected()
	return
}

func (this_ *ManageLoginStorage) Use(id int64) (res int64, err error) {
	m := this_.dbService.SqlUpdate("manage_login")
	m.Value("update_at", time.Now().UnixMilli())
	m.Value("use_at", time.Now().UnixMilli())
	m.Where().Eq("login_id", id)
	r, err := m.Exec(context.Background())
	if err != nil {
		return
	}
	res, err = r.RowsAffected()
	return
}

func (this_ *ManageLoginStorage) DeleteByIds(ids []int64) (res int64, err error) {
	m := this_.dbService.SqlUpdate("manage_login")
	m.Value("delete_at", time.Now().UnixMilli())
	m.Value("status", 9)
	m.Where().In("login_id", ids)
	r, err := m.Exec(context.Background())
	if err != nil {
		return
	}
	res, err = r.RowsAffected()
	return
}

func (this_ *ManageLoginStorage) Query(in *module_manage.ManageLogin) (res []*module_manage.ManageLogin, err error) {
	m := this_.dbService.ModelSelect(in)
	m.Where().AndWhereSql("status != 9")
	res, err = db.DoQueryListWithModel[*module_manage.ManageLogin](context.Background(), m)
	return
}

func (this_ *ManageLoginStorage) Page(in *module_manage.ManageLogin, pageNo int64, pageSize int64) (res []*module_manage.ManageLogin, err error) {
	m := this_.dbService.ModelSelect(in)
	m.Where().AndWhereSql("status != 9")
	res, err = db.DoQueryPageWithModel[*module_manage.ManageLogin](context.Background(), m, pageNo, pageSize)
	return
}

func (this_ *ManageLoginStorage) Count(in *module_manage.ManageLogin) (res int64, err error) {
	m := this_.dbService.ModelCount(in)
	m.Where().AndWhereSql("status != 9")
	res, err = m.Count(context.Background())
	return
}

func (this_ *ManageLoginStorage) RemoveByIds(ids []int64) (res int64, err error) {
	m := this_.dbService.SqlDelete("manage_login")
	m.Where().In("login_id", ids)
	m.Where().AndWhereSql("status = 9")
	r, err := m.Exec(context.Background())
	if err != nil {
		return
	}
	res, err = r.RowsAffected()
	return
}
