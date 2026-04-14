// 文件由 TeamIDE | coos 生成，请勿修改文件内容！通过 [TeamIDE:teamide@163.com] 的 [models:] 在 [2026-04-14 14:18] 生成

package manage_service

import (
	"errors"

	"github.com/team-ide/framework"
	"github.com/team-ide/modules/module_manage"
	"github.com/team-ide/modules/module_manage/manage_factory"
)

func NewManageUserService() *ManageUserService {
	res := &ManageUserService{}
	return res
}

type ManageUserService struct {
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

func (this_ *ManageUserService) Delete(userIds []int64) (err error) {
	_, err = manage_factory.ManageUserStorage.DeleteByIds(userIds)
	if err != nil {
		framework.Error("call ManageUserStorage func DeleteByIds error:" + err.Error())
		return
	}
	return
}

func (this_ *ManageUserService) Remove(userIds []int64) (err error) {
	_, err = manage_factory.ManageUserStorage.RemoveByIds(userIds)
	if err != nil {
		framework.Error("call ManageUserStorage func RemoveByIds error:" + err.Error())
		return
	}
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
