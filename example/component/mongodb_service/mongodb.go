package mongodb_service

import (
	"errors"
	"example/common"
	"github.com/team-ide/framework"
	"github.com/team-ide/framework/mongodb"
	"math"
)

func NewMongodbService(name string, cfg *mongodb.Config) (ser mongodb.IService, err error) {
	if cfg == nil {
		err = errors.New("组件 [mongodb] [" + name + "] 配置为空，请检查配置")
		framework.Error(err.Error())
		return
	}

	if cfg.Disabled {
		framework.Warn("组件 [mongodb] [" + name + "] 已禁用，不创建")
		return
	}

	// 创建 服务
	ser, err = mongodb.New(cfg)
	if err != nil {
		err = errors.New("组件 [mongodb] [" + name + "] 创建 失败: " + err.Error())
		framework.Error(err.Error())
		return
	}
	framework.Info("组件 [mongodb] [" + name + "] 创建 成功")

	common.Starter.OnEvent(framework.EventStop, func(args ...any) {
		framework.Warn("监听 停止事件 关闭 组件 [mongodb] [" + name + "]")
		ser.Close()
	}, math.MaxInt)

	return
}
