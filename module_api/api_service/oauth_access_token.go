package api_service

import (
	"errors"
	"github.com/team-ide/framework"
	"github.com/team-ide/framework/httpx"
	"github.com/team-ide/modules/module_api"
	"github.com/team-ide/modules/module_api/api_factory"
	"time"
)

func NewOauthAccessTokenService(httpService httpx.IService) *OauthAccessTokenService {
	res := &OauthAccessTokenService{
		httpService: httpService,
	}
	return res
}

type OauthAccessTokenService struct {
	httpService httpx.IService
}

func (this_ *OauthAccessTokenService) GetAccessToken(name string) (res string, err error) {
	info := api_factory.OauthAccessTokenCache.Get(name)
	if info == nil {
		err = errors.New("GetAccessToken name [" + name + "] info not found")
		framework.Error(err.Error())
		return
	}
	var nowTime = time.Now()
	var expiresAtTime = time.UnixMilli(info.ExpiresAt).Format("2006-01-02 15:04:05")
	if info.ExpiresAt > nowTime.UnixMilli() {
		framework.Debug("GetAccessToken name [" + name + "] accessToken 过期时间 [" + expiresAtTime + "] 未过期，直接使用")
		res = info.AccessToken
		return
	}
	framework.Debug("GetAccessToken name [" + name + "] accessToken 过期时间 [" + expiresAtTime + "] 已过期，重新获取")
	r, err := module_api.GetAccessToken(info.PlatformType, this_.httpService, info)
	if err != nil {
		return
	}
	info.AccessToken = r.AccessToken
	info.ExpiresAt = nowTime.UnixMilli() + r.ExpiresIn*1000
	_, err = api_factory.OauthAccessTokenStorage.UpdateAccessToken(info.Name, r.AccessToken, r.ExpiresIn, info.ExpiresAt)
	if err != nil {
		err = errors.New("GetAccessToken name [" + name + "] UpdateAccessToken error:" + err.Error())
		framework.Error(err.Error())
		return
	}
	res = r.AccessToken
	return
}
