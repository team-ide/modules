package module_manage

import (
	"time"
)

// loginSessionTimeout 登录 过期 时间 毫秒
var loginSessionTimeout int64 = 2 * 60 * 60 * 1000

func UploadLoginSessionTimeout(in int64) {
	loginSessionTimeout = in
}

func (this_ *LoginInfo) Valid() (err error) {
	useTime := time.Now().UnixMilli() - this_.UseAt
	if loginSessionTimeout > 0 {
		if useTime > loginSessionTimeout {
			err = ERROR_LOGIN_INVALID
			return
		}
	}
	if this_.User.Status != 1 {
		return ERROR_MANAGE_USER_INVALID
	}
	return
}

func (this_ *LoginInfo) Auth(path string) (err error) {
	return
}
