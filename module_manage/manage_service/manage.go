// 文件由 TeamIDE | coos 生成，请勿修改文件内容！通过 [TeamIDE:teamide@163.com] 的 [models:] 在 [2026-04-09 11:11] 生成

package manage_service

import (
	"context"
	"errors"
	"time"

	"github.com/team-ide/framework"
	"github.com/team-ide/framework/util"
	"github.com/team-ide/modules/module_id"
	"github.com/team-ide/modules/module_manage"
	"github.com/team-ide/modules/module_manage/manage_factory"
)

func NewManageService() *ManageService {
	res := &ManageService{}
	return res
}

type ManageService struct {
}

func (this_ *ManageService) GetSecurityCode(in *module_manage.GetSecurityCodeRequest) (out *module_manage.GetSecurityCodeResponse, err error) {
	return
}

func (this_ *ManageService) EncodePassword(salt string, password string) (res string) {
	res = util.GetMd5(salt + password)
	return
}

func (this_ *ManageService) Login(in *module_manage.LoginRequest) (res *module_manage.LoginResponse, err error) {
	if in.Account == "" {
		err = errors.New("登录账号不能为空")
		return
	}
	if in.Password == "" {
		err = errors.New("登录密码不能为空")
		return
	}
	findUsers, err := manage_factory.ManageUserStorage.QueryByAccount(in.Account)
	if err != nil {
		framework.Error("call ManageUserStorage func QueryByAccount error:" + err.Error())
		return
	}
	if len(findUsers) == 0 {
		err = module_manage.ERROR_ACCOUNT_OR_PASSWORD
		return
	}
	findUser := findUsers[0]
	inPassword := this_.EncodePassword(findUser.Salt, in.Password)
	if findUser.Password != inPassword {
		err = module_manage.ERROR_ACCOUNT_OR_PASSWORD
		return
	}
	if findUser.Status != 1 {
		err = module_manage.ERROR_MANAGE_USER_INVALID
		return
	}
	login := &module_manage.ManageLogin{}
	login.UserId = findUser.UserId
	login.UserAccount = findUser.Account
	login.UserName = findUser.Name
	login.CreateAt = time.Now().UnixMilli()
	login.LoginAt = time.Now().UnixMilli()
	login.UseAt = time.Now().UnixMilli()
	login.SourceType = in.SourceType
	login.SourceInfo = in.SourceInfo
	login.LoginId = manage_factory.ManageId.GenLoginId()
	login.Token = util.GetUuid()
	_, err = manage_factory.ManageLoginStorage.Insert(login)
	if err != nil {
		framework.Error("call ManageLoginStorage func Insert error:" + err.Error())
		return
	}
	res, err = this_.Session(login.Token)
	if err != nil {
		framework.Error("call this func Session error:" + err.Error())
		return
	}
	return
}

func (this_ *ManageService) Session(token string) (res *module_manage.LoginResponse, err error) {
	res, err = this_.LoadLogin(token)
	if err != nil {
		framework.Error("call this func LoadLogin error:" + err.Error())
		return
	}
	manage_factory.LoginCache.Add(res.LoginInfo)
	return
}

func (this_ *ManageService) Logout(token string) (err error) {
	return
}

func (this_ *ManageService) LoadLogin(token string) (res *module_manage.LoginResponse, err error) {
	login, err := manage_factory.ManageLoginStorage.GetByToken(token)
	if err != nil {
		framework.Error("call ManageLoginStorage func GetByToken error:" + err.Error())
		return
	}
	if login == nil {
		return
	}
	user, err := manage_factory.ManageUserStorage.GetById(login.UserId)
	if err != nil {
		framework.Error("call ManageUserStorage func GetById error:" + err.Error())
		return
	}
	if user == nil {
		return
	}
	userRoles, err := manage_factory.ManageRoleService.GetUserRoles(login.UserId)
	if err != nil {
		framework.Error("call ManageRoleService func GetUserRoles error:" + err.Error())
		return
	}
	res = &module_manage.LoginResponse{}
	res.LoginInfo = &module_manage.LoginInfo{}
	res.Login = login
	res.User = user
	res.Roles = userRoles
	return
}

func NewManageId() *ManageId {
	res := &ManageId{}
	return res
}

type ManageId struct {
}

func (this_ *ManageId) GenUserId() (res int64) {
	res, err := module_id.IdWorker.GetId(context.Background(), module_manage.IdTypeModuleManageManageUser)
	if err != nil {
		framework.Error("id worker get [IdTypeModuleManageManageUser] id error:" + err.Error())
		return
	}
	return
}

func (this_ *ManageId) GenRoleId() (res int64) {
	res, err := module_id.IdWorker.GetId(context.Background(), module_manage.IdTypeModuleManageManageRole)
	if err != nil {
		framework.Error("id worker get [IdTypeModuleManageManageRole] id error:" + err.Error())
		return
	}
	return
}

func (this_ *ManageId) GenLoginId() (res int64) {
	res, err := module_id.IdWorker.GetId(context.Background(), module_manage.IdTypeModuleManageManageLogin)
	if err != nil {
		framework.Error("id worker get [IdTypeModuleManageManageLogin] id error:" + err.Error())
		return
	}
	return
}

func (this_ *ManageId) GenLogId() (res int64) {
	res, err := module_id.IdWorker.GetId(context.Background(), module_manage.IdTypeModuleManageManageLog)
	if err != nil {
		framework.Error("id worker get [IdTypeModuleManageManageLog] id error:" + err.Error())
		return
	}
	return
}
