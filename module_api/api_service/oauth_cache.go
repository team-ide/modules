package api_service

import (
	"github.com/team-ide/framework"
	"github.com/team-ide/modules/module_api"
	"github.com/team-ide/modules/module_api/api_factory"
)

var (
	// OauthAccessTokenCacheLimit 缓存 数量
	OauthAccessTokenCacheLimit int = 0
	// OauthAccessTokenCacheTtl 缓存 过期时间 2 小时
	OauthAccessTokenCacheTtl int64 = 1000 * 60 * 60 * 2
)

func NewOauthAccessTokenCache() *OauthAccessTokenCache {
	res := new(OauthAccessTokenCache)
	opts := &framework.CacheOptions{
		Limit: OauthAccessTokenCacheLimit,
		Ttl:   OauthAccessTokenCacheTtl,
	}
	res.Cache = framework.NewCacheByOptions(opts, &module_api.OauthAccessToken{})
	return res
}

type OauthAccessTokenCache struct {
	*framework.Cache[*module_api.OauthAccessToken]
}

func (this_ *OauthAccessTokenCache) Add(data *module_api.OauthAccessToken) {
	this_.Set(data.Name, data)
	return
}

func (this_ *OauthAccessTokenCache) Get(name string) (res *module_api.OauthAccessToken) {
	res, err := this_.GetOrLoad(name, func() (*module_api.OauthAccessToken, error) {
		find, err := api_factory.OauthAccessTokenStorage.GetByName(name)
		if err != nil {
			return nil, err
		}
		if find == nil {
			return nil, err
		}

		return find, err
	})
	if err != nil {
		framework.Error("login cache load error:" + err.Error())
		return
	}
	return
}

func (this_ *OauthAccessTokenCache) GetIfPresent(token string) (res *module_api.OauthAccessToken) {
	res, _ = this_.Cache.Get(token)
	return
}
