package api_init

import (
	"github.com/team-ide/framework"
	"github.com/team-ide/framework/util"
	"github.com/team-ide/modules/module_api"
	"github.com/team-ide/modules/module_api/api_factory"
)

func InitOauth(name string, platformType, appId, appSecret string, extends map[string]any) (err error) {
	d := &module_api.OauthAccessToken{
		Name:         name,
		PlatformType: platformType,
		AppId:        appId,
		AppSecret:    appSecret,
	}
	if extends == nil {
		extends = make(map[string]any)
	}
	d.Extends, err = util.ObjToJson(d)
	if err != nil {
		return
	}
	finds, err := api_factory.OauthAccessTokenStorage.Query(&module_api.OauthAccessToken{Name: name})
	if err != nil {
		return
	}
	if len(finds) > 0 {
		if finds[0].AppId != appId || finds[0].AppSecret != appSecret || finds[0].Extends != d.Extends {
			_, err = api_factory.OauthAccessTokenStorage.DeleteByName(name)
			if err != nil {
				return
			}
			finds = nil
		}
	}
	if len(finds) > 0 {
		_, err = api_factory.OauthAccessTokenStorage.Update(d)
	} else {
		_, err = api_factory.OauthAccessTokenStorage.Insert(d)
	}
	if err != nil {
		return
	}
	framework.Info("oauth [" + name + "] 生成成功")
	return
}
