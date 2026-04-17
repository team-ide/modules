package start

import (
	"example/component"
	"example/component/web_service"
	"example/config"
	"github.com/team-ide/framework/web"
	"github.com/team-ide/modules/module_manage/manage_api"

	"github.com/team-ide/framework/web/web_fasthttp"
	"github.com/team-ide/framework/web/web_gin"
)

func init() {
	component.Starter.AddServerStartFunc("start web api", 100, startWebApi)

	component.Starter.AddServerStartFunc("start web manage", 100, startWebServerManage)
}

func startWebApi() (err error) {

	cfg := config.GetConfig()

	var webServiceOpt = web_fasthttp.BindWebService

	if err = web_service.InitWebApi("web_api", cfg.WebApi, webServiceOpt); err != nil {
		return
	}

	component.Starter.SetShouldWait(true)
	return
}

func startWebServerManage() (err error) {
	cfg := config.GetConfig()

	var webServiceOpt = web_gin.BindWebService

	var webHandler = func(s *web.WebServer) {
		s.AddWebFilters(manage_api.NewLogFilter())
		s.AddWebFilters(manage_api.NewLoginFilter())
		s.AddWebApis(manage_api.NewWebApiManageService().GetWebApi())
	}
	if err = web_service.InitWebManage("web_manage", cfg.WebManage, webServiceOpt, webHandler); err != nil {
		return
	}

	component.Starter.SetShouldWait(true)
	return
}
