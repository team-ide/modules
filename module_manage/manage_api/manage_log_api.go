// 文件由 TeamIDE | coos 生成，请勿修改文件内容！通过 [TeamIDE:teamide@163.com] 的 [models:] 在 [2026-04-10 17:28] 生成

package manage_api

import (
	"github.com/team-ide/framework/web"
	"github.com/team-ide/modules/module_manage"
	"github.com/team-ide/modules/module_manage/manage_factory"
)

func NewWebApiManageLogService() *WebApiManageLogService {
	res := &WebApiManageLogService{}
	return res
}

type WebApiManageLogService struct {
}

func (this_ *WebApiManageLogService) GetWebApi() (webApi *web.WebApi) {
	webApi = web.NewWebApi("/manage/log/")
	webApi.Add("page", this_.Page)
	webApi.Add("delete", this_.Delete)
	return
}

func (this_ *WebApiManageLogService) Page(request *web.WebRequest) (res any, err error) {
	in := &module_manage.ManageLogPageRequest{}
	if err = request.RequestJSON(in); err != nil {
		return
	}

	res, err = manage_factory.ManageLogService.Page(in)
	return
}

func (this_ *WebApiManageLogService) Delete(request *web.WebRequest) (res any, err error) {
	in := &module_manage.ManageLogDeleteRequest{}
	if err = request.RequestJSON(in); err != nil {
		return
	}

	err = manage_factory.ManageLogService.Delete(in)
	return
}
