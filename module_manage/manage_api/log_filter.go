package manage_api

import (
	"time"

	"github.com/team-ide/framework/web"
	"github.com/team-ide/modules/module_manage"
	"github.com/team-ide/modules/module_manage/manage_factory"
)

func NewLogFilter() *LogFilter {
	res := &LogFilter{}
	return res
}

type LogFilter struct {
}

func (this_ *LogFilter) GetName() (res string) {
	res = "manage log filter"
	return
}
func (this_ *LogFilter) Order() (res int) {
	res = 200
	return
}
func (this_ *LogFilter) GetMatch() (res web.WebRequestMatch) {
	res.Includes = append(res.Includes, "/*")
	//res.Methods = append(res.Methods, "POST")

	return
}
func (this_ *LogFilter) DoFilter(req *web.WebRequest, chain web.WebFilterChain) (err error) {
	log := &module_manage.ManageLog{}
	log.StartAt = time.Now().UnixMilli()
	err = chain(req)
	if req.WebApiRouter != nil && !req.WebApiRouter.NotLog {
		log.EndAt = time.Now().UnixMilli()
		log.UseTime = int(log.EndAt - log.StartAt)
		login := GetLogin(req)
		var loginUser *module_manage.ManageUser
		if login != nil {
			log.LoginId = login.LoginId
			loginUser = login.User
		} else {
			if req.Response != nil {
				if loginRes, isLoginRes := req.Response.(module_manage.LoginResponse); isLoginRes {
					loginUser = loginRes.User
				}
			}
		}
		if loginUser != nil {
			log.UserId = loginUser.UserId
			log.UserName = loginUser.Name
			log.UserAccount = loginUser.Account
		}
		log.Ip = req.ClientIP()
		log.UserAgent = req.UserAgent()
		log.Method = req.Method
		log.Path = req.Path
		log.Comment = req.WebApiRouter.Comment
		if err != nil {
			log.Error = err.Error()
		}
		// 获取GET请求的所有查询参数
		log.Param = req.GetPathParam()
		log.Data = string(req.GetData())
		_ = manage_factory.ManageLogService.Add(log)
	}
	return
}
