// 文件由 TeamIDE | coos 生成，请勿修改文件内容！通过 [TeamIDE:teamide@163.com] 的 [models:] 在 [2026-04-10 14:03] 生成

package manage_service

import (
	"github.com/team-ide/framework"
	"github.com/team-ide/modules/module_manage"
	"github.com/team-ide/modules/module_manage/manage_factory"
)

func NewManageRoleService() *ManageRoleService {
	res := &ManageRoleService{}
	return res
}

type ManageRoleService struct {
}

func (this_ *ManageRoleService) Add(in *module_manage.ManageRoleAddRequest) (res *module_manage.ManageRoleAddResponse, err error) {
	if in.RoleId == 0 {
		in.RoleId = manage_factory.ManageId.GenRoleId()
	}
	_, err = manage_factory.ManageRoleStorage.Insert(in.ManageRole)
	if err != nil {
		framework.Error("call ManageRoleStorage func Insert error:" + err.Error())
		return
	}
	return
}

func (this_ *ManageRoleService) Delete(roleIds []int64) (err error) {
	_, err = manage_factory.ManageRoleStorage.DeleteByIds(roleIds)
	if err != nil {
		framework.Error("call ManageRoleStorage func DeleteByIds error:" + err.Error())
		return
	}
	return
}

func (this_ *ManageRoleService) GetUserRoles(userId int64) (res []*module_manage.ManageRole, err error) {
	userRoleIds, err := manage_factory.ManageRoleUserStorage.QueryUserRoleIds(userId)
	if err != nil {
		framework.Error("call ManageRoleUserStorage func QueryUserRoleIds error:" + err.Error())
		return
	}
	res, err = manage_factory.ManageRoleStorage.GetByIds(userRoleIds)
	if err != nil {
		framework.Error("call ManageRoleStorage func GetByIds error:" + err.Error())
		return
	}
	return
}

func (this_ *ManageRoleService) AddRoleUsers(roleId int64, userIds []int64) (err error) {
	for _, userId := range userIds {
		err = this_.AddRoleUser(roleId, userId)
		if err != nil {
			framework.Error("call this func AddRoleUser error:" + err.Error())
			return
		}
	}
	return
}

func (this_ *ManageRoleService) AddRoleUser(roleId int64, userId int64) (err error) {
	roleUsers, err := manage_factory.ManageRoleUserStorage.QueryByRoleIdAndUserId(roleId, userId)
	if err != nil {
		framework.Error("call ManageRoleUserStorage func QueryByRoleIdAndUserId error:" + err.Error())
		return
	}
	if len(roleUsers) > 0 {
		return
	}
	a := &module_manage.ManageRoleUser{}
	a.RoleId = roleId
	a.UserId = userId
	_, err = manage_factory.ManageRoleUserStorage.Insert(a)
	if err != nil {
		framework.Error("call ManageRoleUserStorage func Insert error:" + err.Error())
		return
	}
	return
}
