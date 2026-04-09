package web_service

import (
	"github.com/team-ide/framework/web"
)

var (
	webApi *web.WebServer
)

func InitWebApi(name string, cfg *web.Config, opts ...web.Option) (err error) {
	// 如果 已存在 则 直接 返回
	if webApi != nil {
		return
	}
	// 创建 服务
	webApi, err = NewWebService(name, cfg, opts...)
	if err != nil {
		return
	}

	return
}

func GetWebApi() (res *web.WebServer) {
	return webApi
}
