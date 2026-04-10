// 文件由 TeamIDE | coos 生成，请勿修改文件内容！通过 [TeamIDE:teamide@163.com] 的 [models:] 在 [2026-04-10 11:02] 生成

package manage_service

import (
	"github.com/team-ide/framework"
	"github.com/team-ide/modules/module_manage"
	"github.com/team-ide/modules/module_manage/manage_factory"
)

func NewManagePermissionService() *ManagePermissionService {
	res := &ManagePermissionService{}
	return res
}

type ManagePermissionService struct {
}

func (this_ *ManagePermissionService) Add(in *module_manage.ManagePermissionAddRequest) (res *module_manage.ManagePermissionAddResponse, err error) {
	if in.PermissionId == 0 {
		in.PermissionId = manage_factory.ManageId.GenPermissionId()
	}
	_, err = manage_factory.ManagePermissionStorage.Insert(in.ManagePermission)
	if err != nil {
		framework.Error("call ManagePermissionStorage func Insert error:" + err.Error())
		return
	}
	return
}

func (this_ *ManagePermissionService) Query(in *module_manage.ManagePermissionQueryRequest) (res []*module_manage.ManagePermission, err error) {
	var list []*module_manage.ManagePermission
	if len(in.PermissionIds) > 0 {
		list, err = this_.QueryByIds(in.PermissionIds)
		if err != nil {
			framework.Error("call this func QueryByIds error:" + err.Error())
			return
		}
		res = append(res, list...)
	}
	if len(in.RoleIds) > 0 {
		list, err = this_.QueryByRoleIds(in.RoleIds)
		if err != nil {
			framework.Error("call this func QueryByRoleIds error:" + err.Error())
			return
		}
		res = append(res, list...)
	}
	if len(in.UserIds) > 0 {
		list, err = this_.QueryByUserIds(in.UserIds)
		if err != nil {
			framework.Error("call this func QueryByUserIds error:" + err.Error())
			return
		}
		res = append(res, list...)
	}
	return
}

func (this_ *ManagePermissionService) QueryByIds(ids []int64) (res []*module_manage.ManagePermission, err error) {
	res, err = manage_factory.ManagePermissionStorage.QueryByIds(ids)
	if err != nil {
		framework.Error("call ManagePermissionStorage func QueryByIds error:" + err.Error())
		return
	}
	return
}

func (this_ *ManagePermissionService) QueryByRoleIds(roleIds []int64) (res []*module_manage.ManagePermission, err error) {
	res, err = manage_factory.ManagePermissionStorage.QueryByRoleIds(roleIds)
	if err != nil {
		framework.Error("call ManagePermissionStorage func QueryByRoleIds error:" + err.Error())
		return
	}
	return
}

func (this_ *ManagePermissionService) QueryByUserIds(userIds []int64) (res []*module_manage.ManagePermission, err error) {
	res, err = manage_factory.ManagePermissionStorage.QueryByUserIds(userIds)
	if err != nil {
		framework.Error("call ManagePermissionStorage func QueryByUserIds error:" + err.Error())
		return
	}
	return
}

func (this_ *ManagePermissionService) Delete(in *module_manage.ManagePermissionDeleteRequest) (err error) {
	if len(in.PermissionIds) > 0 {
		_, err = manage_factory.ManagePermissionStorage.DeleteByIds(in.PermissionIds)
		if err != nil {
			framework.Error("call ManagePermissionStorage func DeleteByIds error:" + err.Error())
			return
		}
	}
	if len(in.RoleIds) > 0 {
		_, err = manage_factory.ManagePermissionStorage.DeleteByRoleIds(in.RoleIds)
		if err != nil {
			framework.Error("call ManagePermissionStorage func DeleteByRoleIds error:" + err.Error())
			return
		}
	}
	if len(in.UserIds) > 0 {
		_, err = manage_factory.ManagePermissionStorage.DeleteByUserIds(in.UserIds)
		if err != nil {
			framework.Error("call ManagePermissionStorage func DeleteByUserIds error:" + err.Error())
			return
		}
	}
	return
}
