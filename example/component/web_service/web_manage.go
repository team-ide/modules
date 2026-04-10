package web_service

import (
	"github.com/team-ide/framework/web"
)

var (
	webManage *web.WebServer
)

func InitWebManage(name string, cfg *web.Config, opts ...web.Option) (err error) {
	// 如果 已存在 则 直接 返回
	if webManage != nil {
		return
	}
	// 创建 服务
	webManage, err = NewWebService(name, cfg, opts...)
	if err != nil {
		return
	}

	return
}

func GetWebManage() (res *web.WebServer) {
	return webManage
}
