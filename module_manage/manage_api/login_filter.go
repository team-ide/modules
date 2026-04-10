package manage_api

import (
	"time"

	"github.com/team-ide/framework/web"
	"github.com/team-ide/modules/module_manage"
	"github.com/team-ide/modules/module_manage/manage_factory"
)

func NewLoginFilter() *LoginFilter {
	res := &LoginFilter{}
	return res
}

type LoginFilter struct {
}

func (this_ *LoginFilter) GetName() (res string) {
	res = "manage login filter"
	return
}
func (this_ *LoginFilter) Order() (res int) {
	res = 100
	return
}
func (this_ *LoginFilter) GetMatch() (res web.WebRequestMatch) {
	res.Includes = append(res.Includes, "/*")
	//res.Methods = append(res.Methods, "POST")

	return
}
func (this_ *LoginFilter) DoFilter(req *web.WebRequest, chain web.WebFilterChain) (err error) {
	var shouldLogin bool

	if req.WebApiRouter != nil && !req.WebApiRouter.NotLogin {
		shouldLogin = true
	}
	if shouldLogin {
		login := GetLogin(req)
		if login == nil {
			err = module_manage.ERROR_SHOULD_LOGIN
			return
		}
		if err = login.Valid(); err != nil {
			return
		}

		// 鉴权
		if err = login.Auth(req.Path); err != nil {
			return
		}

		login.UseAt = time.Now().UnixMilli()
		_, _ = manage_factory.ManageLoginStorage.Use(login.LoginId)
	}
	err = chain(req)
	return
}

func GetLogin(req *web.WebRequest) (login *module_manage.LoginInfo) {
	token := req.GetHeader("token")
	if token == "" {
		token = req.GetParam("token")
	}
	if token != "" {
		login = manage_factory.LoginCache.Get(token)
	}
	return
}
