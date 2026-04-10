// 文件由 TeamIDE | coos 生成，请勿修改文件内容！通过 [TeamIDE:teamide@163.com] 的 [models:] 在 [2026-04-10 10:50] 生成

package manage_api

import (
	"github.com/team-ide/framework/web"
	"github.com/team-ide/modules/module_manage"
	"github.com/team-ide/modules/module_manage/manage_factory"
)

type WebApiManagePermissionService struct {
}

func (this_ *WebApiManagePermissionService) GetWebApi(webApi *web.WebApi) {
	webApi = web.NewWebApi("/manage/permission/")
	webApi.Add("add", this_.Add)
	webApi.Add("query", this_.Query)
	webApi.Add("query/by/ids", this_.QueryByIds)
	webApi.Add("query/by/role/ids", this_.QueryByRoleIds)
	webApi.Add("query/by/user/ids", this_.QueryByUserIds)
	webApi.Add("delete", this_.Delete)
	return
}

func (this_ *WebApiManagePermissionService) Add(request *web.WebRequest) (res any, err error) {
	in := &module_manage.ManagePermissionAddRequest{}
	if err = request.RequestJSON(in); err != nil {
		return
	}

	res, err = manage_factory.ManagePermissionService.Add(in)
	return
}

func (this_ *WebApiManagePermissionService) Query(request *web.WebRequest) (res any, err error) {
	in := &module_manage.ManagePermissionQueryRequest{}
	if err = request.RequestJSON(in); err != nil {
		return
	}

	res, err = manage_factory.ManagePermissionService.Query(in)
	return
}

func (this_ *WebApiManagePermissionService) QueryByIds(request *web.WebRequest) (res any, err error) {
	ids := []int64{}
	if err = request.RequestJSON(ids); err != nil {
		return
	}

	res, err = manage_factory.ManagePermissionService.QueryByIds(ids)
	return
}

func (this_ *WebApiManagePermissionService) QueryByRoleIds(request *web.WebRequest) (res any, err error) {
	roleIds := []int64{}
	if err = request.RequestJSON(roleIds); err != nil {
		return
	}

	res, err = manage_factory.ManagePermissionService.QueryByRoleIds(roleIds)
	return
}

func (this_ *WebApiManagePermissionService) QueryByUserIds(request *web.WebRequest) (res any, err error) {
	userIds := []int64{}
	if err = request.RequestJSON(userIds); err != nil {
		return
	}

	res, err = manage_factory.ManagePermissionService.QueryByUserIds(userIds)
	return
}

func (this_ *WebApiManagePermissionService) Delete(request *web.WebRequest) (res any, err error) {
	in := &module_manage.ManagePermissionDeleteRequest{}
	if err = request.RequestJSON(in); err != nil {
		return
	}

	err = manage_factory.ManagePermissionService.Delete(in)
	return
}
