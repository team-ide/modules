package manage_service

import (
	"github.com/team-ide/framework"
	"github.com/team-ide/modules/module_manage"
	"github.com/team-ide/modules/module_manage/manage_factory"
)

var (
	// LoginCacheLimit 缓存 数量
	LoginCacheLimit int = 0
	// LoginCacheTtl 缓存 过期时间 2 小时
	LoginCacheTtl int64 = 1000 * 60 * 60 * 2
)

func NewLoginCache() *LoginCache {
	res := new(LoginCache)
	opts := &framework.CacheOptions{
		Limit: LoginCacheLimit,
		Ttl:   LoginCacheTtl,
	}
	res.Cache = framework.NewCacheByOptions(opts, &module_manage.LoginInfo{})
	return res
}

type LoginCache struct {
	*framework.Cache[*module_manage.LoginInfo]
}

func (this_ *LoginCache) Add(data *module_manage.LoginInfo) {
	this_.Set(data.Login.Token, data)
	return
}

func (this_ *LoginCache) Get(token string) (res *module_manage.LoginInfo) {
	res, err := this_.GetOrLoad(token, func() (*module_manage.LoginInfo, error) {
		find, err := manage_factory.ManageService.LoadLogin(token)
		if err != nil {
			return nil, err
		}
		if find == nil {
			return nil, err
		}
		return find.LoginInfo, err
	})
	if err != nil {
		framework.Error("login cache load error:" + err.Error())
		return
	}
	return
}

func (this_ *LoginCache) GetIfPresent(token string) (res *module_manage.LoginInfo) {
	res, _ = this_.Cache.Get(token)
	return
}
