package api_github

import (
	"github.com/team-ide/framework/httpx"
	"github.com/team-ide/modules/module_api"
	"net/url"
	"strings"
)

var BaiduPlatform *module_api.Platform

func init() {
	BaiduPlatform = &module_api.Platform{}
	BaiduPlatform.Type = "github"
	BaiduPlatform.SetGetAccessToken(GetAccessTokenFunc)
}

var GetAccessTokenUrl = "https://github.com/login/oauth/access_token"

type GetAccessTokenResponse struct {
	AccessToken           string `json:"access_token"`
	ExpiresIn             int64  `json:"expires_in"`
	RefreshToken          string `json:"refresh_token"`
	RefreshTokenExpiresIn int64  `json:"refresh_token_expires_in"`
	Scope                 string `json:"scope"`
	TokenType             string `json:"token_type"`
	Error                 string `json:"error"`
	ErrorDescription      string `json:"error_description"`
}

func GetAccessTokenFunc(httpService httpx.IService, appId, appSecret string, extends map[string]any) (res *module_api.OauthInfo, err error) {
	var code string
	if extends != nil {
		code = extends["code"].(string)
	}
	apiUrl := GetAccessTokenUrl + "?grant_type=client_credentials&client_id=" + appId + "&client_secret=" + appSecret
	data := url.Values{}
	data.Set("client_id", appId)
	data.Set("client_secret", appSecret)
	data.Set("code", code)
	resp, err := httpService.PostRequest(apiUrl, strings.NewReader(data.Encode()), func(in *httpx.Request) {
		in.SetHeader("Content-Type", "application/x-www-form-urlencoded")
		in.SetHeader("Accept", "application/json")
	})
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
