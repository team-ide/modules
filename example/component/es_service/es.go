package es_service

import (
	"errors"
	"example/component"
	"github.com/team-ide/framework"
	"github.com/team-ide/framework/es"
	"math"
)

func NewEsService(name string, cfg *es.Config) (ser es.IService, err error) {
	if cfg == nil {
		err = errors.New("组件 [es] [" + name + "] 配置为空，请检查配置")
		framework.Error(err.Error())
		return
	}

	if cfg.Disabled {
		framework.Warn("组件 [es] [" + name + "] 已禁用，不创建")
		return
	}

	// 创建 服务
	ser, err = es.New(cfg)
	if err != nil {
		err = errors.New("组件 [es] [" + name + "] 创建 失败: " + err.Error())
		framework.Error(err.Error())
		return
	}
	framework.Info("组件 [es] [" + name + "] 创建 成功")

	component.Starter.OnEvent(framework.EventStop, func(args ...any) {
		framework.Warn("监听 停止事件 关闭 组件 [es] [" + name + "]")
		ser.Close()
	}, math.MaxInt)

	return
}
