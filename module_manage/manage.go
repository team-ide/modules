// 文件由 TeamIDE | coos 生成，请勿修改文件内容！通过 [TeamIDE:teamide@163.com] 的 [models:module_manage/service/manage.coos] 在 [2026-04-09 11:11] 生成

package module_manage

import (
	"github.com/team-ide/framework"
	"github.com/team-ide/modules/module_id"
)

var ERROR_SHOULD_LOGIN = framework.NewError("10000", "请先登录")

var ERROR_LOGIN_INVALID = framework.NewError("10001", "登录凭证已失效")

var ERROR_LOGIN_EXPIRE = framework.NewError("10002", "登录凭证已过期")

var ERROR_ACCOUNT_OR_PASSWORD = framework.NewError("10005", "登录账号或密码错误")

var ERROR_MANAGE_USER_INVALID = framework.NewError("10004", "管理员账号已失效")

var ERROR_OLD_PASSWORD = framework.NewError("10006", "旧密码错误")

type LoginRequest struct {
	Account         string `json:"account"`
	Password        string `json:"password"`
	DeviceId        string `json:"deviceId"`
	SecurityCodeKey string `json:"securityCodeKey"`
	SecurityCode    string `json:"securityCode"`
	SourceType      int    `json:"sourceType"`
	SourceInfo      string `json:"sourceInfo"`
}

type LoginInfo struct {
	User  *ManageUser   `json:"user"`
	Roles []*ManageRole `json:"roles"`
	Login *ManageLogin  `json:"login"`
}

type LoginResponse struct {
	*LoginInfo
}

type GetSecurityCodeRequest struct {
	DeviceId string `json:"deviceId"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
}

type GetSecurityCodeResponse struct {
	SecurityCodeBase64 string `json:"securityCodeBase64"`
	ExpirationTime     int    `json:"expirationTime"`
	SecurityCodeKey    string `json:"securityCodeKey"`
}

type IManageService interface {
	GetSecurityCode(in *GetSecurityCodeRequest) (out *GetSecurityCodeResponse, err error)
	EncodePassword(salt string, password string) (res string)
	Login(in *LoginRequest) (res *LoginResponse, err error)
	Session(token string) (res *LoginResponse, err error)
	Logout(token string) (err error)
	LoadLogin(token string) (res *LoginResponse, err error)
}

var IdTypeModuleManageManageUser = module_id.IdType("module_manage:manage_user")

var IdTypeModuleManageManageRole = module_id.IdType("module_manage:manage_role")

var IdTypeModuleManageManageLogin = module_id.IdType("module_manage:manage_login")

var IdTypeModuleManageManageLog = module_id.IdType("module_manage:manage_log")

type IManageId interface {
	GenUserId() (res int64)
	GenRoleId() (res int64)
	GenLoginId() (res int64)
	GenLogId() (res int64)
}

type ILoginCache interface {
	Add(data *LoginInfo)
	Get(token string) (res *LoginInfo)
	GetIfPresent(token string) (res *LoginInfo)
}
