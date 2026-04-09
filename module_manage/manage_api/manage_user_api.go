// 文件由 TeamIDE | coos 生成，请勿修改文件内容！通过 [TeamIDE:teamide@163.com] 的 [models:] 在 [2026-04-09 11:11] 生成

package manage_api

import (
	"github.com/team-ide/framework/web"
	"github.com/team-ide/modules/module_manage"
	"github.com/team-ide/modules/module_manage/manage_factory"
)

type WebApiManageUserService struct {
}

func (this_ *WebApiManageUserService) GetWebApi(webApi *web.WebApi) {
	webApi = web.NewWebApi("/manage/user/")
	webApi.Add("add", this_.Add)
	webApi.Add("delete", this_.Delete)
	webApi.Add("a", this_.A)
	return
}

func (this_ *WebApiManageUserService) Add(request *web.WebRequest) (res any, err error) {
	in := &module_manage.ManageUserAddRequest{}
	if err = request.RequestJSON(in); err != nil {
		return
	}

	res, err = manage_factory.ManageUserService.Add(in)
	return
}

func (this_ *WebApiManageUserService) Delete(request *web.WebRequest) (res any, err error) {
	in := &module_manage.ManageUserDeleteRequest{}
	if err = request.RequestJSON(in); err != nil {
		return
	}

	err = manage_factory.ManageUserService.Delete(in)
	return
}

func (this_ *WebApiManageUserService) A(request *web.WebRequest) (res any, err error) {
	var token string
	token = request.GetHeader(token)

	err = manage_factory.ManageUserService.A(token)
	return
}
