// 文件由 TeamIDE | coos 生成，请勿修改文件内容！通过 [TeamIDE:teamide@163.com] 的 [models:] 在 [2026-04-17 16:28] 生成

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

func (this_ *ManageUserService) List(in *module_manage.ManageUserListRequest) (res *module_manage.ManageUserListResponse, err error) {
	res = &module_manage.ManageUserListResponse{}
	res.List, err = manage_factory.ManageUserStorage.Query(in.ManageUser)
	if err != nil {
		framework.Error("call ManageUserStorage func Query error:" + err.Error())
		return
	}
	return
}

func (this_ *ManageUserService) Page(in *module_manage.ManageUserPageRequest) (res *module_manage.ManageUserPageResponse, err error) {
	res = &module_manage.ManageUserPageResponse{}
	res.List, err = manage_factory.ManageUserStorage.Page(in.ManageUser, in.PageNo, in.PageSize)
	if err != nil {
		framework.Error("call ManageUserStorage func Page error:" + err.Error())
		return
	}
	res.Total, err = manage_factory.ManageUserStorage.Count(in.ManageUser)
	if err != nil {
		framework.Error("call ManageUserStorage func Count error:" + err.Error())
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
