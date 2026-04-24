// 文件由 TeamIDE | coos 生成，请勿修改文件内容！通过 [TeamIDE:teamide@163.com] 的 [models:module_api/api.coos] 在 [2026-04-24 15:51] 生成

package module_api

import (
	"github.com/team-ide/modules/module_id"
)

var IdTypeModuleApiId = module_id.IdType("module_api:id")

type IGenId interface {
	GenId() (res int64)
}

type IOauthAccessTokenCache interface {
	Add(data *OauthAccessToken)
	Get(name string) (res *OauthAccessToken)
	GetIfPresent(name string) (res *OauthAccessToken)
}
