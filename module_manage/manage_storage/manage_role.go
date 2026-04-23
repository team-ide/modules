// 文件由 TeamIDE | coos 生成，请勿修改文件内容！通过 [TeamIDE:teamide@163.com] 的 [models:] 在 [2026-04-23 15:33] 生成

package manage_storage

import (
	"context"
	"time"

	"github.com/team-ide/framework/db"
	"github.com/team-ide/modules/module_manage"
)

func NewManageRoleStorage(dbManage db.IService) *ManageRoleStorage {
	res := &ManageRoleStorage{}
	res.dbManage = dbManage
	return res
}

type ManageRoleStorage struct {
	dbManage db.IService
}

func (this_ *ManageRoleStorage) GetById(id int64) (res *module_manage.ManageRole, err error) {
	m := this_.dbManage.SqlSelect("manage_role")
	m.CanSelectAll()
	m.Where().Eq("role_id", id)
	res, err = db.DoQueryOneWithModel[*module_manage.ManageRole](context.Background(), m)
	return
}

func (this_ *ManageRoleStorage) GetByIds(ids []int64) (res []*module_manage.ManageRole, err error) {
	m := this_.dbManage.SqlSelect("manage_role")
	m.CanSelectAll()
	m.Where().In("role_id", ids)
	res, err = db.DoQueryListWithModel[*module_manage.ManageRole](context.Background(), m)
	return
}

func (this_ *ManageRoleStorage) Insert(in *module_manage.ManageRole) (res int64, err error) {
	in.CreateAt = time.Now().UnixMilli()
	r, err := this_.dbManage.Insert(context.Background(), in, func(in *db.ModelInsert) {
	})
	if err != nil {
		return
	}
	res, err = r.RowsAffected()
	return
}

func (this_ *ManageRoleStorage) Update(in *module_manage.ManageRole) (res int64, err error) {
	in.UpdateAt = time.Now().UnixMilli()
	r, err := this_.dbManage.Update(context.Background(), in, func(in *db.ModelUpdate) {
	})
	if err != nil {
		return
	}
	res, err = r.RowsAffected()
	return
}

func (this_ *ManageRoleStorage) DeleteByIds(ids []int64) (res int64, err error) {
	m := this_.dbManage.SqlDelete("manage_role", func(in *db.SqlDelete) {
	})
	m.Where().In("role_id", ids)
	r, err := m.Exec(context.Background())
	if err != nil {
		return
	}
	res, err = r.RowsAffected()
	return
}

func (this_ *ManageRoleStorage) Query(in *module_manage.ManageRole) (res []*module_manage.ManageRole, err error) {
	m := this_.dbManage.ModelSelect(in, func(in *db.ModelSelect) {
		in.CanSelectAll()
	})
	m.WhereOperator("name", "%like%")
	res, err = db.DoQueryListWithModel[*module_manage.ManageRole](context.Background(), m)
	return
}

func (this_ *ManageRoleStorage) Page(in *module_manage.ManageRole, pageNo int64, pageSize int64) (res []*module_manage.ManageRole, err error) {
	m := this_.dbManage.ModelSelect(in, func(in *db.ModelSelect) {
		in.CanSelectAll()
	})
	m.WhereOperator("name", "%like%")
	res, err = db.DoQueryPageWithModel[*module_manage.ManageRole](context.Background(), m, pageNo, pageSize)
	return
}

func (this_ *ManageRoleStorage) Count(in *module_manage.ManageRole) (res int64, err error) {
	m := this_.dbManage.ModelCount(in, func(in *db.ModelCount) {
	})
	m.WhereOperator("name", "%like%")
	res, err = m.Count(context.Background())
	return
}

func NewManageRoleUserStorage(dbManage db.IService) *ManageRoleUserStorage {
	res := &ManageRoleUserStorage{}
	res.dbManage = dbManage
	return res
}

type ManageRoleUserStorage struct {
	dbManage db.IService
}

func (this_ *ManageRoleUserStorage) Insert(in *module_manage.ManageRoleUser) (res int64, err error) {
	in.CreateAt = time.Now().UnixMilli()
	r, err := this_.dbManage.Insert(context.Background(), in, func(in *db.ModelInsert) {
	})
	if err != nil {
		return
	}
	res, err = r.RowsAffected()
	return
}

func (this_ *ManageRoleUserStorage) DeleteByIds(ids []int64) (res int64, err error) {
	m := this_.dbManage.SqlDelete("manage_role_user", func(in *db.SqlDelete) {
	})
	m.Where().In("role_user_id", ids)
	r, err := m.Exec(context.Background())
	if err != nil {
		return
	}
	res, err = r.RowsAffected()
	return
}

func (this_ *ManageRoleUserStorage) QueryByRoleId(roleId int64) (res []*module_manage.ManageRoleUser, err error) {
	m := this_.dbManage.SqlSelect("manage_role_user")
	m.CanSelectAll()
	m.Where().Eq("role_id", roleId)
	res, err = db.DoQueryListWithModel[*module_manage.ManageRoleUser](context.Background(), m)
	return
}

func (this_ *ManageRoleUserStorage) QueryRoleUserIds(roleId int64) (res []int64, err error) {
	m := this_.dbManage.SqlSelect("manage_role_user")
	m.CanSelectAll()
	m.Where().Eq("role_id", roleId)
	r, err := db.DoQueryListWithModel[*module_manage.ManageRoleUser](context.Background(), m)
	if err != nil {
		return
	}
	for _, one := range r {
		res = append(res, one.UserId)
	}
	return
}

func (this_ *ManageRoleUserStorage) QueryByUserId(userId int64) (res []*module_manage.ManageRoleUser, err error) {
	m := this_.dbManage.SqlSelect("manage_role_user")
	m.CanSelectAll()
	m.Where().Eq("user_id", userId)
	res, err = db.DoQueryListWithModel[*module_manage.ManageRoleUser](context.Background(), m)
	return
}

func (this_ *ManageRoleUserStorage) QueryUserRoleIds(userId int64) (res []int64, err error) {
	m := this_.dbManage.SqlSelect("manage_role_user")
	m.CanSelectAll()
	m.Where().Eq("user_id", userId)
	r, err := db.DoQueryListWithModel[*module_manage.ManageRoleUser](context.Background(), m)
	if err != nil {
		return
	}
	for _, one := range r {
		res = append(res, one.RoleId)
	}
	return
}

func (this_ *ManageRoleUserStorage) Query(in *module_manage.ManageRoleUser) (res []*module_manage.ManageRoleUser, err error) {
	m := this_.dbManage.ModelSelect(in, func(in *db.ModelSelect) {
		in.CanSelectAll()
	})
	res, err = db.DoQueryListWithModel[*module_manage.ManageRoleUser](context.Background(), m)
	return
}

func (this_ *ManageRoleUserStorage) QueryByRoleIdAndUserId(roleId int64, userId int64) (res []*module_manage.ManageRoleUser, err error) {
	m := this_.dbManage.SqlSelect("manage_role_user")
	m.CanSelectAll()
	m.Where().Eq("role_id", roleId)
	m.Where().Eq("user_id", userId)
	res, err = db.DoQueryListWithModel[*module_manage.ManageRoleUser](context.Background(), m)
	return
}
