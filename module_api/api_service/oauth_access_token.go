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
	info, err := api_factory.OauthAccessTokenStorage.GetByName(name)
	if err != nil {
		return
	}
	if info == nil {
		err = errors.New("GetAccessToken name [" + name + "] info not found")
		framework.Error(err.Error())
		return
	}
	if info.ExpiresAt > time.Now().UnixMilli() {
		res = info.AccessToken
		return
	}
	r, err := module_api.GetAccessToken(info.PlatformType, this_.httpService, info)
	if err != nil {
		return
	}
	_, err = api_factory.OauthAccessTokenStorage.UpdateAccessToken(info.Name, r.AccessToken, r.ExpiresIn)
	if err != nil {
		err = errors.New("GetAccessToken name [" + name + "] UpdateAccessToken error:" + err.Error())
		framework.Error(err.Error())
		return
	}
	res = r.AccessToken
	return
}
