package start

import (
	"example/common"
	"example/component/web_service"
	"example/config"

	"github.com/team-ide/framework/web/web_fasthttp"
	"github.com/team-ide/framework/web/web_gin"
)

func init() {
	common.Starter.AddServerStartFunc("start web api", 100, startWebApi)

	common.Starter.AddServerStartFunc("start web manage", 100, startWebServerManage)
}

func startWebApi() (err error) {

	cfg := config.GetConfig()

	var webServiceOpt = web_fasthttp.BindWebService

	if err = web_service.InitWebApi("web_api", cfg.WebApi, webServiceOpt); err != nil {
		return
	}

	common.Starter.SetShouldWait(true)
	return
}

func startWebServerManage() (err error) {
	cfg := config.GetConfig()

	var webServiceOpt = web_gin.BindWebService

	if err = web_service.InitWebApi("web_manage", cfg.WebManage, webServiceOpt); err != nil {
		return
	}

	common.Starter.SetShouldWait(true)
	return
}
