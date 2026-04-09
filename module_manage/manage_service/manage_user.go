// 文件由 TeamIDE | coos 生成，请勿修改文件内容！通过 [TeamIDE:teamide@163.com] 的 [models:] 在 [2026-04-09 11:11] 生成

package manage_service

import (
	"errors"

	"github.com/team-ide/framework"
	"github.com/team-ide/modules/module_manage"
	"github.com/team-ide/modules/module_manage/manage_factory"
)

func NewManageUserService(manageUserStorage module_manage.IManageUserStorage) *ManageUserService {
	res := &ManageUserService{}
	res.manageUserStorage = manageUserStorage
	return res
}

type ManageUserService struct {
	manageUserStorage module_manage.IManageUserStorage
}

func (this_ *ManageUserService) Add(in *module_manage.ManageUserAddRequest) (res *module_manage.ManageUserAddResponse, err error) {
	if in.Name == "" {
		err = errors.New("name 不能为空")
		return
	}
	if in.Password == "" {
		err = errors.New("password 不能为空")
		return
	}
	if in.UserId == 0 {
		in.UserId = manage_factory.ManageId.GenUserId()
	}
	_, err = manage_factory.ManageUserStorage.Insert(in.ManageUser)
	if err != nil {
		framework.Error("call ManageUserStorage func Insert error:" + err.Error())
		return
	}
	return
}

func (this_ *ManageUserService) Delete(in *module_manage.ManageUserDeleteRequest) (err error) {
	return
}

func (this_ *ManageUserService) A(token string) (err error) {
	return
}

func (this_ *ManageUserService) Xxx(in *module_manage.ManageUserDeleteRequest) (err error) {
	return
}

// ---------- start ----------

func (this_ *ManageUserService) Aaa() (err error) {
	return
}

// ---------- end ----------

// ---------- start ----------

func (this_ *ManageUserService) B() (err error) {
	return
}

// ---------- end ----------

// ---------- start ----------

func (this_ *ManageUserService) C() (err error) {
	return
}

// ---------- end ----------
