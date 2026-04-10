package start

import (
	"context"
	"example/common"
	"example/factory"
	"example/storage"
	"fmt"

	"github.com/team-ide/framework/util"
	"github.com/team-ide/modules/module_manage"
	"github.com/team-ide/modules/module_manage/manage_factory"
	"github.com/team-ide/modules/module_manage/manage_init"
)

func init() {
	common.Starter.AddInitDataFunc("init test data", 100, InitTestData)
}

var (
	InitManageUserSaveFile = "data/manage-user.json"
)

func InitTestData() (err error) {

	tbTest1 := &storage.TbTest1{}
	err = factory.TbTest1Storage.Insert(context.Background(), tbTest1)
	if err != nil {
		return
	}
	fmt.Println("insert table ["+tbTest1.GetTableName()+"] row:", util.GetStringValue(tbTest1))
	tbTest1List, err := factory.TbTest1Storage.Query(context.Background(), &storage.TbTest1{})
	if err != nil {
		return
	}
	for _, one := range tbTest1List {
		fmt.Println(util.GetStringValue(one))
	}

	err = manage_init.InitManageSuperUser("管理员", "admin", InitManageUserSaveFile)
	if err != nil {
		return
	}

	login, err := manage_factory.ManageService.Login(&module_manage.LoginRequest{
		Account:  "admin",
		Password: util.GetMd5("f52b880eaf"),
	})
	if err != nil {
		return
	}
	fmt.Println("login:" + util.GetStringValue(login))
	loginInfo := manage_factory.LoginCache.Get(login.Token)
	fmt.Println("loginInfo:" + util.GetStringValue(loginInfo))

	return
}
