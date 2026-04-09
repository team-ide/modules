package db_service

import (
	"errors"
	"example/common"
	"math"

	"github.com/team-ide/framework"
	"github.com/team-ide/framework/db"
)

func NewDbService(name string, cfg *db.Config) (ser db.IService, err error) {
	if cfg == nil {
		err = errors.New("组件 [db] [" + name + "] 配置为空，请检查配置")
		framework.Error(err.Error())
		return
	}

	if cfg.Disabled {
		framework.Warn("组件 [db] [" + name + "] 已禁用，不创建")
		return
	}

	// 创建 服务
	s, err := db.New(cfg, nil)
	if err != nil {
		err = errors.New("组件 [db] [" + name + "] 创建 失败: " + err.Error())
		framework.Error(err.Error())
		return
	}
	s.OpenShowExecSql().OpenShowQuerySql()
	ser = s
	framework.Info("组件 [db] [" + name + "] 创建 成功")

	// 数据 方言、模式等配置
	modelOption := db.NewModelOption()
	s.SetModelOption(modelOption)

	common.Starter.OnEvent(framework.EventStop, func(args ...any) {
		framework.Warn("监听 停止事件 关闭 组件 [db] [" + name + "]")
		ser.Close()
	}, math.MaxInt)

	return
}
