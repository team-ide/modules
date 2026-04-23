// 文件由 TeamIDE | coos 生成，请勿修改文件内容！通过 [TeamIDE:teamide@163.com] 的 [models:] 在 [2026-04-23 15:33] 生成

package api_install

import (
	"github.com/team-ide/framework/db"
)

func InstallDbApiTable(dbService db.IService) (err error) {
	if err = TableOauthAccessTokenCreate(dbService); err != nil {
		return
	}
	return
}
