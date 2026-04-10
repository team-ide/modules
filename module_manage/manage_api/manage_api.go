// 文件由 TeamIDE | coos 生成，请勿修改文件内容！通过 [TeamIDE:teamide@163.com] 的 [models:] 在 [2026-04-10 17:29] 生成

package manage_api

import (
	"github.com/team-ide/framework/web"
	"github.com/team-ide/modules/module_manage"
	"github.com/team-ide/modules/module_manage/manage_factory"
)

func NewWebApiManageService() *WebApiManageService {
	res := &WebApiManageService{}
	return res
}

type WebApiManageService struct {
}

func (this_ *WebApiManageService) GetWebApi() (webApi *web.WebApi) {
	webApi = web.NewWebApi("/")
	webApi.Add("getSecurityCode", this_.GetSecurityCode).SetNotLogin().SetNotLog().SetComment("获取图形验证码")
	webApi.Add("login", this_.Login).SetNotLogin().SetComment("登录")
	webApi.Add("session", this_.Session).SetComment("获取登录用户信息")
	webApi.Add("logout", this_.Logout).SetNotLogin().SetComment("登出")
	webApi.Add("load/login", this_.LoadLogin).SetNotLogin().SetComment("加载登录信息")
	return
}

func (this_ *WebApiManageService) GetSecurityCode(request *web.WebRequest) (res any, err error) {
	in := &module_manage.GetSecurityCodeRequest{}
	if err = request.RequestJSON(in); err != nil {
		return
	}

	res, err = manage_factory.ManageService.GetSecurityCode(in)
	return
}

func (this_ *WebApiManageService) Login(request *web.WebRequest) (res any, err error) {
	in := &module_manage.LoginRequest{}
	if err = request.RequestJSON(in); err != nil {
		return
	}

	res, err = manage_factory.ManageService.Login(in)
	return
}

func (this_ *WebApiManageService) Session(request *web.WebRequest) (res any, err error) {
	var token string
	token = request.GetHeader("token")

	res, err = manage_factory.ManageService.Session(token)
	return
}

func (this_ *WebApiManageService) Logout(request *web.WebRequest) (res any, err error) {
	var token string
	token = request.GetHeader("token")

	err = manage_factory.ManageService.Logout(token)
	return
}

func (this_ *WebApiManageService) LoadLogin(request *web.WebRequest) (res any, err error) {
	var token string

	res, err = manage_factory.ManageService.LoadLogin(token)
	return
}
