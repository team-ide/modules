// 文件由 TeamIDE | coos 生成，请勿修改文件内容！通过 [TeamIDE:teamide@163.com] 的 [models:] 在 [2026-04-23 15:33] 生成

package manage_storage

import (
	"context"
	"time"

	"github.com/team-ide/framework/db"
	"github.com/team-ide/modules/module_manage"
)

func NewManageUserStorage(dbManage db.IService) *ManageUserStorage {
	res := &ManageUserStorage{}
	res.dbManage = dbManage
	return res
}

type ManageUserStorage struct {
	dbManage db.IService
}

func (this_ *ManageUserStorage) GetById(id int64) (res *module_manage.ManageUser, err error) {
	m := this_.dbManage.SqlSelect("manage_user")
	m.CanSelectAll()
	m.SelectExclude("password", "salt")
	m.Where().Eq("user_id", id)
	res, err = db.DoQueryOneWithModel[*module_manage.ManageUser](context.Background(), m)
	return
}

func (this_ *ManageUserStorage) GetByIds(ids []int64) (res []*module_manage.ManageUser, err error) {
	m := this_.dbManage.SqlSelect("manage_user")
	m.CanSelectAll()
	m.SelectExclude("password", "salt")
	m.Where().In("user_id", ids)
	res, err = db.DoQueryListWithModel[*module_manage.ManageUser](context.Background(), m)
	return
}

func (this_ *ManageUserStorage) Insert(in *module_manage.ManageUser) (res int64, err error) {
	in.CreateAt = time.Now().UnixMilli()
	if in.Status == 0 {
		in.Status = 1
	}
	r, err := this_.dbManage.Insert(context.Background(), in, func(in *db.ModelInsert) {
		in.ValueExclude("update_at", "delete_at", "disable_at")
	})
	if err != nil {
		return
	}
	res, err = r.RowsAffected()
	return
}

func (this_ *ManageUserStorage) Update(in *module_manage.ManageUser) (res int64, err error) {
	in.UpdateAt = time.Now().UnixMilli()
	r, err := this_.dbManage.Update(context.Background(), in, func(in *db.ModelUpdate) {
	})
	if err != nil {
		return
	}
	res, err = r.RowsAffected()
	return
}

func (this_ *ManageUserStorage) Disable(id int64) (res int64, err error) {
	m := this_.dbManage.SqlUpdate("manage_user", func(in *db.SqlUpdate) {
	})
	m.Value("update_at", time.Now().UnixMilli())
	m.Value("status", 2)
	m.Value("disable_at", time.Now().UnixMilli())
	m.Where().Eq("user_id", id)
	r, err := m.Exec(context.Background())
	if err != nil {
		return
	}
	res, err = r.RowsAffected()
	return
}

func (this_ *ManageUserStorage) Enable(id int64) (res int64, err error) {
	m := this_.dbManage.SqlUpdate("manage_user", func(in *db.SqlUpdate) {
	})
	m.Value("update_at", time.Now().UnixMilli())
	m.Value("status", 1)
	m.Where().Eq("user_id", id)
	r, err := m.Exec(context.Background())
	if err != nil {
		return
	}
	res, err = r.RowsAffected()
	return
}

func (this_ *ManageUserStorage) UpdatePassword(id int64, salt string, password string) (res int64, err error) {
	m := this_.dbManage.SqlUpdate("manage_user", func(in *db.SqlUpdate) {
	})
	m.Value("salt", salt)
	m.Value("password", password)
	m.Value("update_at", time.Now().UnixMilli())
	m.Where().Eq("user_id", id)
	r, err := m.Exec(context.Background())
	if err != nil {
		return
	}
	res, err = r.RowsAffected()
	return
}

func (this_ *ManageUserStorage) QueryByAccount(account string) (res []*module_manage.ManageUser, err error) {
	m := this_.dbManage.SqlSelect("manage_user")
	m.CanSelectAll()
	m.Where().Eq("account", account)
	m.Where().AndWhereSql("status != 9")
	res, err = db.DoQueryListWithModel[*module_manage.ManageUser](context.Background(), m)
	return
}

func (this_ *ManageUserStorage) CountByAccount(account string) (res int64, err error) {
	m := this_.dbManage.SqlCount("manage_user", func(in *db.SqlCount) {
	})
	m.Where().Eq("account", account)
	m.Where().AndWhereSql("status != 9")
	res, err = m.Count(context.Background())
	return
}

func (this_ *ManageUserStorage) DeleteByIds(ids []int64) (res int64, err error) {
	m := this_.dbManage.SqlUpdate("manage_user", func(in *db.SqlUpdate) {
	})
	m.Value("delete_at", time.Now().UnixMilli())
	m.Value("status", 9)
	m.Where().In("user_id", ids)
	r, err := m.Exec(context.Background())
	if err != nil {
		return
	}
	res, err = r.RowsAffected()
	return
}

func (this_ *ManageUserStorage) Query(in *module_manage.ManageUser) (res []*module_manage.ManageUser, err error) {
	m := this_.dbManage.ModelSelect(in, func(in *db.ModelSelect) {
		in.CanSelectAll()
		in.SelectExclude("password", "salt")
	})
	m.WhereOperator("name", "%like%")
	m.Where().AndWhereSql("status != 9")
	res, err = db.DoQueryListWithModel[*module_manage.ManageUser](context.Background(), m)
	return
}

func (this_ *ManageUserStorage) Page(in *module_manage.ManageUser, pageNo int64, pageSize int64) (res []*module_manage.ManageUser, err error) {
	m := this_.dbManage.ModelSelect(in, func(in *db.ModelSelect) {
		in.CanSelectAll()
		in.SelectExclude("password", "salt")
	})
	m.WhereOperator("name", "%like%")
	m.Where().AndWhereSql("status != 9")
	res, err = db.DoQueryPageWithModel[*module_manage.ManageUser](context.Background(), m, pageNo, pageSize)
	return
}

func (this_ *ManageUserStorage) Count(in *module_manage.ManageUser) (res int64, err error) {
	m := this_.dbManage.ModelCount(in, func(in *db.ModelCount) {
	})
	m.WhereOperator("name", "%like%")
	m.Where().AndWhereSql("status != 9")
	res, err = m.Count(context.Background())
	return
}

func (this_ *ManageUserStorage) RemoveByIds(ids []int64) (res int64, err error) {
	m := this_.dbManage.SqlDelete("manage_user", func(in *db.SqlDelete) {
	})
	m.Where().In("user_id", ids)
	m.Where().AndWhereSql("status = 9")
	r, err := m.Exec(context.Background())
	if err != nil {
		return
	}
	res, err = r.RowsAffected()
	return
}
