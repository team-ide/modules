package module_api

import (
	"errors"
	"github.com/team-ide/framework"
	"github.com/team-ide/framework/httpx"
	"github.com/team-ide/framework/util"
)

var (
	platformCache = map[string]*Platform{}
)

func RegisterPlatform(platform *Platform) {
	platformCache[platform.Type] = platform
}

func GetPlatform(platformType string) *Platform {
	return platformCache[platformType]
}

type OauthInfo struct {
	AccessToken string `json:"accessToken"`
	// ExpiresIn 有效期（秒），部分平台可能返回时间戳，业务层转换后存储
	ExpiresIn int64 `json:"expiresIn"`
}

type GetAccessTokenFunc = func(httpService httpx.IService, appId, appSecret string, extends map[string]any) (res *OauthInfo, err error)

type Platform struct {
	Type           string `json:"type,omitempty"`
	getAccessToken GetAccessTokenFunc
}

func (this_ *Platform) SetGetAccessToken(getAccessToken GetAccessTokenFunc) *Platform {
	this_.getAccessToken = getAccessToken
	return this_
}

func GetAccessToken(platformType string, httpService httpx.IService, info *OauthAccessToken) (res *OauthInfo, err error) {
	platform := GetPlatform(platformType)
	if platform == nil {
		err = errors.New("platform [" + platformType + "] is null")
		framework.Error(err.Error())
		return
	}
	if platform.getAccessToken == nil {
		err = errors.New("platform [" + platformType + "] name [" + info.Name + "] getAccessToken is null")
		framework.Error(err.Error())
		return
	}
	appId := info.AppId
	appSecret := info.AppSecret
	var extends = map[string]any{}
	if info.Extends != "" {
		err = util.JsonBytesToObj([]byte(info.Extends), &extends)
		if err != nil {
			err = errors.New("platform [" + platformType + "] name [" + info.Name + "] extends to map error:" + err.Error())
			framework.Error(err.Error())
			return
		}
	}
	res, err = platform.getAccessToken(httpService, appId, appSecret, extends)
	if err != nil {
		err = errors.New("platform [" + platformType + "] name [" + info.Name + "] getAccessToken error:" + err.Error())
		framework.Error(err.Error())
		return
	}
	return
}
