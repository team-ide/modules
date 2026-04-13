package module_manage

import (
	"time"
)

// LoginSessionTimeout 登录 过期 时间 毫秒
var LoginSessionTimeout = 2 * 60 * 60 * 1000

func (this_ *LoginInfo) Valid() (err error) {
	useTime := time.Now().UnixMilli() - this_.UseAt
	if useTime > int64(LoginSessionTimeout) {
		err = ERROR_LOGIN_INVALID
		return
	}
	if this_.User.Status != 1 {
		return ERROR_MANAGE_USER_INVALID
	}
	return
}

func (this_ *LoginInfo) Auth(path string) (err error) {
	return
}
