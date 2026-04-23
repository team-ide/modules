package api_baidu

import (
	"github.com/team-ide/framework/httpx"
	"github.com/team-ide/modules/module_api"
)

var BaiduPlatform *module_api.Platform

func init() {
	BaiduPlatform = &module_api.Platform{}
	BaiduPlatform.Type = "baidu"
	BaiduPlatform.SetGetAccessToken(GetAccessTokenFunc)
}

var GetAccessTokenUrl = "https://aip.baidubce.com/oauth/2.0/token"

type GetAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

func GetAccessTokenFunc(httpService httpx.IService, appId, appSecret string, extends map[string]any) (res *module_api.OauthInfo, err error) {
	apiUrl := GetAccessTokenUrl + "?grant_type=client_credentials&client_id=" + appId + "&client_secret=" + appSecret
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
