// 文件由 TeamIDE | coos 生成，请勿修改文件内容！通过 [TeamIDE:teamide@163.com] 的 [models:] 在 [2026-04-10 10:50] 生成

package manage_storage

import (
	"context"
	"time"

	"github.com/team-ide/framework/db"
	"github.com/team-ide/modules/module_manage"
)

func NewManagePermissionStorage(dbService db.IService) *ManagePermissionStorage {
	res := &ManagePermissionStorage{}
	res.dbService = dbService
	return res
}

type ManagePermissionStorage struct {
	dbService db.IService
}

func (this_ *ManagePermissionStorage) Insert(in *module_manage.ManagePermission) (res int64, err error) {
	in.CreateAt = time.Now().UnixMilli()
	r, err := this_.dbService.Insert(context.Background(), in)
	if err != nil {
		return
	}
	res, err = r.RowsAffected()
	return
}

func (this_ *ManagePermissionStorage) QueryByIds(ids []int64) (res []*module_manage.ManagePermission, err error) {
	m := this_.dbService.SqlSelect("manage_permission")
	m.Where().In("permission_id", ids)
	res, err = db.DoQueryListWithModel[*module_manage.ManagePermission](context.Background(), m)
	return
}

func (this_ *ManagePermissionStorage) QueryByRoleId(roleId int64) (res []*module_manage.ManagePermission, err error) {
	m := this_.dbService.SqlSelect("manage_permission")
	m.Where().Eq("role_id", roleId)
	res, err = db.DoQueryListWithModel[*module_manage.ManagePermission](context.Background(), m)
	return
}

func (this_ *ManagePermissionStorage) QueryByRoleIds(roleIds []int64) (res []*module_manage.ManagePermission, err error) {
	m := this_.dbService.SqlSelect("manage_permission")
	m.Where().In("role_id", roleIds)
	res, err = db.DoQueryListWithModel[*module_manage.ManagePermission](context.Background(), m)
	return
}

func (this_ *ManagePermissionStorage) QueryByUserId(userId int64) (res []*module_manage.ManagePermission, err error) {
	m := this_.dbService.SqlSelect("manage_permission")
	m.Where().Eq("user_id", userId)
	res, err = db.DoQueryListWithModel[*module_manage.ManagePermission](context.Background(), m)
	return
}

func (this_ *ManagePermissionStorage) QueryByUserIds(userIds []int64) (res []*module_manage.ManagePermission, err error) {
	m := this_.dbService.SqlSelect("manage_permission")
	m.Where().In("user_id", userIds)
	res, err = db.DoQueryListWithModel[*module_manage.ManagePermission](context.Background(), m)
	return
}

func (this_ *ManagePermissionStorage) Query(in *module_manage.ManagePermission) (res []*module_manage.ManagePermission, err error) {
	m := this_.dbService.ModelSelect(in)
	res, err = db.DoQueryListWithModel[*module_manage.ManagePermission](context.Background(), m)
	return
}

func (this_ *ManagePermissionStorage) DeleteByIds(ids []int64) (res int64, err error) {
	m := this_.dbService.SqlDelete("manage_permission")
	m.Where().In("permission_id", ids)
	r, err := m.Exec(context.Background())
	if err != nil {
		return
	}
	res, err = r.RowsAffected()
	return
}

func (this_ *ManagePermissionStorage) DeleteByRoleIds(roleIds []int64) (res int64, err error) {
	m := this_.dbService.SqlDelete("manage_permission")
	m.Where().In("role_id", roleIds)
	r, err := m.Exec(context.Background())
	if err != nil {
		return
	}
	res, err = r.RowsAffected()
	return
}

func (this_ *ManagePermissionStorage) DeleteByUserIds(userIds []int64) (res int64, err error) {
	m := this_.dbService.SqlDelete("manage_permission")
	m.Where().In("user_id", userIds)
	r, err := m.Exec(context.Background())
	if err != nil {
		return
	}
	res, err = r.RowsAffected()
	return
}
