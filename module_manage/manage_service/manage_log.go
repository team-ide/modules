// 文件由 TeamIDE | coos 生成，请勿修改文件内容！通过 [TeamIDE:teamide@163.com] 的 [models:] 在 [2026-04-10 16:15] 生成

package manage_service

import (
	"github.com/team-ide/framework"
	"github.com/team-ide/modules/module_manage"
	"github.com/team-ide/modules/module_manage/manage_factory"
)

func NewManageLogService() *ManageLogService {
	res := &ManageLogService{}
	return res
}

type ManageLogService struct {
}

func (this_ *ManageLogService) Add(in *module_manage.ManageLog) (err error) {
	if in.LogId == 0 {
		in.LogId = manage_factory.ManageId.GenLogId()
	}
	_, err = manage_factory.ManageLogStorage.Insert(in)
	if err != nil {
		framework.Error("call ManageLogStorage func Insert error:" + err.Error())
		return
	}
	return
}

func (this_ *ManageLogService) Page(in *module_manage.ManageLogPageRequest) (res *module_manage.ManageLogPageResponse, err error) {
	res = &module_manage.ManageLogPageResponse{}
	res.List, err = manage_factory.ManageLogStorage.Page(in.ManageLog, in.PageNo, in.PageSize)
	if err != nil {
		framework.Error("call ManageLogStorage func Page error:" + err.Error())
		return
	}
	res.Total, err = manage_factory.ManageLogStorage.Count(in.ManageLog)
	if err != nil {
		framework.Error("call ManageLogStorage func Count error:" + err.Error())
		return
	}
	return
}

func (this_ *ManageLogService) Delete(in *module_manage.ManageLogDeleteRequest) (err error) {
	_, err = manage_factory.ManageLogStorage.DeleteByIds(in.LogIds)
	if err != nil {
		framework.Error("call ManageLogStorage func DeleteByIds error:" + err.Error())
		return
	}
	return
}
