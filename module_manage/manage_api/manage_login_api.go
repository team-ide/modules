// 文件由 TeamIDE | coos 生成，请勿修改文件内容！通过 [TeamIDE:teamide@163.com] 的 [models:] 在 [2026-04-10 17:28] 生成

package manage_api

import (
	"github.com/team-ide/framework/web"
	"github.com/team-ide/modules/module_manage"
	"github.com/team-ide/modules/module_manage/manage_factory"
)

func NewWebApiManageLoginService() *WebApiManageLoginService {
	res := &WebApiManageLoginService{}
	return res
}

type WebApiManageLoginService struct {
}

func (this_ *WebApiManageLoginService) GetWebApi() (webApi *web.WebApi) {
	webApi = web.NewWebApi("/manage/login/")
	webApi.Add("page", this_.Page)
	webApi.Add("delete", this_.Delete)
	return
}

func (this_ *WebApiManageLoginService) Page(request *web.WebRequest) (res any, err error) {
	in := &module_manage.ManageLoginPageRequest{}
	if err = request.RequestJSON(in); err != nil {
		return
	}

	res, err = manage_factory.ManageLoginService.Page(in)
	return
}

func (this_ *WebApiManageLoginService) Delete(request *web.WebRequest) (res any, err error) {
	in := &module_manage.ManageLoginDeleteRequest{}
	if err = request.RequestJSON(in); err != nil {
		return
	}

	err = manage_factory.ManageLoginService.Delete(in)
	return
}
