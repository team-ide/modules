// 文件由 TeamIDE | coos 生成，请勿修改文件内容！通过 [TeamIDE:teamide@163.com] 的 [models:] 在 [2026-04-24 16:08] 生成

package api_install

import (
	"github.com/team-ide/framework/db"
)

func TableOauthAccessTokenCreate(dbService db.IService) (err error) {
	var version = "v1.0.0"
	var tableName = "oauth_access_token"
	var moduleName = "api"
	// 定义表结构
	table := &db.Table{
		Name:    tableName,
		Comment: "第三方平台access_token统一存储表",
		ColumnList: []*db.Column{
			{Name: "name", DataType: "varchar", Length: 50, NotNull: true, Key: true, Comment: "名称 唯一"},
			{Name: "platform_type", DataType: "varchar", Length: 50, Comment: "平台类型：wechat_mp, wechat_mini, baidu, github, gitlab, gitee, alipay, dingtalk 等"},
			{Name: "oauth_url", DataType: "varchar", Length: 200, Comment: "oauth url 通常不用填写，如果是私服则填写"},
			{Name: "app_id", DataType: "varchar", Length: 200, Comment: "应用唯一标识（公众号AppId/支付宝AppId/GitHub ClientId等）"},
			{Name: "app_secret", DataType: "varchar", Length: 200, Comment: "应用 secret"},
			{Name: "extends", DataType: "varchar", Length: 2000, Comment: "扩展属性"},
			{Name: "access_token", DataType: "varchar", Length: 500, Comment: "access_token"},
			{Name: "expires_in", DataType: "bigint", Length: 20, Comment: "有效期（秒），部分平台可能返回时间戳，业务层转换后存储"},
			{Name: "expires_at", DataType: "bigint", Length: 20, Comment: "实际过期时间 = 当前时间 + expires_in，便于查询过期"},
			{Name: "status", DataType: "int", Length: 10, Default: "1", Comment: "状态 1：正常"},
			{Name: "created_at", DataType: "bigint", Length: 20, Comment: "创建 时间戳 毫秒"},
			{Name: "updated_at", DataType: "bigint", Length: 20, Comment: "修改 时间戳 毫秒"},
		},
		IndexList: []*db.Index{
			{Name: "", ColumnNames: []string{"platform_type"}},
			{Name: "", ColumnNames: []string{"app_id"}},
			{Name: "", ColumnNames: []string{"created_at"}},
			{Name: "", ColumnNames: []string{"updated_at"}},
		},
	}
	err = db.TableCreate(dbService, moduleName, version, tableName, table)
	return
}
