// 文件由 TeamIDE | coos 生成，请勿修改文件内容！通过 [TeamIDE:teamide@163.com] 的 [models:] 在 [2026-04-10 17:28] 生成

package manage_api

import (
	"github.com/team-ide/framework/web"
	"github.com/team-ide/modules/module_manage"
	"github.com/team-ide/modules/module_manage/manage_factory"
)

func NewWebApiManageRoleService() *WebApiManageRoleService {
	res := &WebApiManageRoleService{}
	return res
}

type WebApiManageRoleService struct {
}

func (this_ *WebApiManageRoleService) GetWebApi() (webApi *web.WebApi) {
	webApi = web.NewWebApi("/manage/role/")
	webApi.Add("add", this_.Add)
	webApi.Add("delete", this_.Delete)
	webApi.Add("get/user/roles", this_.GetUserRoles)
	webApi.Add("add/role/users", this_.AddRoleUsers)
	webApi.Add("add/role/user", this_.AddRoleUser)
	return
}

func (this_ *WebApiManageRoleService) Add(request *web.WebRequest) (res any, err error) {
	in := &module_manage.ManageRoleAddRequest{}
	if err = request.RequestJSON(in); err != nil {
		return
	}

	res, err = manage_factory.ManageRoleService.Add(in)
	return
}

func (this_ *WebApiManageRoleService) Delete(request *web.WebRequest) (res any, err error) {
	in := &module_manage.ManageRoleDeleteRequest{}
	if err = request.RequestJSON(in); err != nil {
		return
	}
	roleIds := in.RoleIds
	err = manage_factory.ManageRoleService.Delete(roleIds)
	return
}

func (this_ *WebApiManageRoleService) GetUserRoles(request *web.WebRequest) (res any, err error) {
	in := &module_manage.ManageUserRoleRequest{}
	if err = request.RequestJSON(in); err != nil {
		return
	}
	userId := in.UserId
	res, err = manage_factory.ManageRoleService.GetUserRoles(userId)
	return
}

func (this_ *WebApiManageRoleService) AddRoleUsers(request *web.WebRequest) (res any, err error) {
	in := &module_manage.ManageAddRoleUsersResponse{}
	if err = request.RequestJSON(in); err != nil {
		return
	}
	roleId := in.RoleId
	userIds := in.UserIds
	err = manage_factory.ManageRoleService.AddRoleUsers(roleId, userIds)
	return
}

func (this_ *WebApiManageRoleService) AddRoleUser(request *web.WebRequest) (res any, err error) {
	in := &module_manage.ManageAddRoleUserResponse{}
	if err = request.RequestJSON(in); err != nil {
		return
	}
	roleId := in.RoleId
	userId := in.UserId
	err = manage_factory.ManageRoleService.AddRoleUser(roleId, userId)
	return
}
