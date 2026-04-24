package api_baidu

import (
	"github.com/team-ide/framework/httpx"
	"github.com/team-ide/modules/module_api"
)

var platform *module_api.Platform

func GetPlatform() *module_api.Platform {
	if platform == nil {
		platform = &module_api.Platform{}
		platform.Type = "baidu"
		platform.SetGetAccessToken(GetAccessTokenFunc)
	}
	return platform
}

var GetAccessTokenUrl = "https://aip.baidubce.com/oauth/2.0/token"

type GetAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

func GetAccessTokenFunc(httpService httpx.IService, appId, appSecret string, extends map[string]any) (res *module_api.OauthInfo, err error) {
	var grantType = "client_credentials"
	if extends != nil {
		if v, ok := extends["grant_type"]; ok {
			grantType = v.(string)
		}
	}
	apiUrl := GetAccessTokenUrl + "?grant_type=" + grantType + "&client_id=" + appId + "&client_secret=" + appSecret
	resp, err := httpService.GetRequest(apiUrl)
	if err != nil {
		return
	}
	r, _, err := httpx.ResponseToObj[*GetAccessTokenResponse](resp)
	if err != nil {
		return
	}
	res = &module_api.OauthInfo{}
	res.AccessToken = r.AccessToken
	res.ExpiresIn = r.ExpiresIn
	return
}
