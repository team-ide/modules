// 文件由 TeamIDE | coos 生成，请勿修改文件内容！通过 [TeamIDE:teamide@163.com] 的 [models:] 在 [2026-04-24 15:45] 生成

package api_service

import (
	"context"

	"github.com/team-ide/framework"
	"github.com/team-ide/modules/module_api"
	"github.com/team-ide/modules/module_id"
)

func NewGenId() *GenId {
	res := &GenId{}
	return res
}

type GenId struct {
}

func (this_ *GenId) GenId() (res int64) {
	res, err := module_id.IdWorker.GetId(context.Background(), module_api.IdTypeModuleApiId)
	if err != nil {
		framework.Error("id worker get [IdTypeModuleApiId] id error:" + err.Error())
		return
	}
	return
}
