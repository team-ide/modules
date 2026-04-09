package web_service

import (
	"errors"
	"example/common"
	"github.com/team-ide/framework"
	"math"

	"github.com/team-ide/framework/web"
)

func NewWebService(name string, cfg *web.Config, opts ...web.Option) (ser *web.WebServer, err error) {
	if cfg == nil {
		err = errors.New("组件 [web] [" + name + "] 配置为空，请检查配置")
		framework.Error(err.Error())
		return
	}
	if cfg.Disabled {
		framework.Warn("组件 [web] [" + name + "] 已禁用，不创建")
		return
	}

	// 创建 服务
	ser = web.New(name, cfg, opts...)

	err = ser.Start()

	if err != nil {
		err = errors.New("组件 [web] [" + name + "] 创建 失败: " + err.Error())
		framework.Error(err.Error())
		return
	}

	common.Starter.OnEvent(framework.EventStop, func(args ...any) {
		framework.Warn("监听 停止事件 关闭 组件 [web] [" + name + "]")
		ser.Close()
	}, math.MaxInt)

	return
}
