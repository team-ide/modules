// 文件由 TeamIDE | coos 生成，请勿修改文件内容！通过 [TeamIDE:teamide@163.com] 的 [models:] 在 [2026-04-14 14:18] 生成

package manage_service

import (
	"github.com/team-ide/framework"
	"github.com/team-ide/modules/module_manage"
	"github.com/team-ide/modules/module_manage/manage_factory"
)

func NewManageLoginService() *ManageLoginService {
	res := &ManageLoginService{}
	return res
}

type ManageLoginService struct {
}

func (this_ *ManageLoginService) Add(in *module_manage.ManageLogin) (err error) {
	if in.LoginId == 0 {
		in.LoginId = manage_factory.ManageId.GenLoginId()
	}
	_, err = manage_factory.ManageLoginStorage.Insert(in)
	if err != nil {
		framework.Error("call ManageLoginStorage func Insert error:" + err.Error())
		return
	}
	return
}

func (this_ *ManageLoginService) Page(in *module_manage.ManageLoginPageRequest) (res *module_manage.ManageLoginPageResponse, err error) {
	res = &module_manage.ManageLoginPageResponse{}
	res.List, err = manage_factory.ManageLoginStorage.Page(in.ManageLogin, in.PageNo, in.PageSize)
	if err != nil {
		framework.Error("call ManageLoginStorage func Page error:" + err.Error())
		return
	}
	res.Total, err = manage_factory.ManageLoginStorage.Count(in.ManageLogin)
	if err != nil {
		framework.Error("call ManageLoginStorage func Count error:" + err.Error())
		return
	}
	return
}

func (this_ *ManageLoginService) Delete(in *module_manage.ManageLoginDeleteRequest) (err error) {
	_, err = manage_factory.ManageLoginStorage.DeleteByIds(in.LoginIds)
	if err != nil {
		framework.Error("call ManageLoginStorage func DeleteByIds error:" + err.Error())
		return
	}
	return
}
