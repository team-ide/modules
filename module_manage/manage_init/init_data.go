package manage_init

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"

	"github.com/team-ide/framework"
	"github.com/team-ide/framework/util"
	"github.com/team-ide/modules/module_manage"
	"github.com/team-ide/modules/module_manage/manage_factory"
)

func InitManageUser(name string, account string, password string, saveFile string) (userId int64, err error) {
	// 检测是否有用户  如果没有 生成用户
	finds, err := manage_factory.ManageUserStorage.QueryByAccount(account)
	if err != nil {
		return
	}
	if len(finds) > 0 {
		userId = finds[0].UserId
		framework.Info("已有 [" + account + "] 管理员用户，不用初始化")

		return
	}
	user := &module_manage.ManageUser{}
	var saveData = map[string]any{}
	if saveFile != "" {
		if e, _ := util.PathExists(saveFile); e {
			bs, _ := os.ReadFile(saveFile)
			if len(bs) > 0 {
				_ = json.Unmarshal(bs, &saveData)
			}
		}
		user.Name = util.GetStringValue(saveData["name"])
		user.Account = util.GetStringValue(saveData["account"])
		user.Salt = util.GetStringValue(saveData["salt"])
		user.Password = util.GetStringValue(saveData["password"])
	}
	if user.Name != name {
		user.Name = name
	}
	if user.Account != account {
		user.Account = account
	}
	user.Status = 1

	if user.Salt == "" {
		user.Salt = util.GetUuid()[0:10]
	}
	if password == "" {
		password = user.Password
		if password == "" {
			password = util.GetUuid()[0:10]
		}
	}

	passwordMD5 := util.GetMd5(password)

	passwordEncode := manage_factory.ManageService.EncodePassword(user.Salt, passwordMD5)
	user.Password = passwordEncode
	user.UserId = manage_factory.ManageId.GenUserId()
	_, err = manage_factory.ManageUserStorage.Insert(user)
	if err != nil {
		return
	}
	user.Password = password

	if saveFile != "" {
		saveData["userId"] = user.UserId
		saveData["name"] = user.Name
		saveData["account"] = user.Account
		saveData["salt"] = user.Salt
		saveData["password"] = password
		saveData["createAt"] = time.Now()

		bs, _ := json.MarshalIndent(saveData, "", "\t")

		saveDir := filepath.Dir(saveFile)
		if err = os.MkdirAll(saveDir, os.ModePerm); err != nil {
			return
		}
		var f *os.File
		f, err = os.Create(saveFile)
		if err != nil {
			return
		}
		_, err = f.WriteString(string(bs))
		if err != nil {
			return
		}
		framework.Info("账号 [" + account + "] 生成成功，用户信息保存在 [" + saveFile + "] 中")
	}
	return
}

func InitManageRole(name string, isSupper int8) (roleId int64, err error) {
	finds, err := manage_factory.ManageRoleStorage.Query(&module_manage.ManageRole{IsSuper: isSupper, Name: name})
	if err != nil {
		return
	}
	if len(finds) > 0 {
		roleId = finds[0].RoleId
		framework.Info("已有 [" + name + "] 角色 [" + util.GetStringValue(roleId) + "]，直接使用，不用初始化")
		return
	}

	role := &module_manage.ManageRole{
		Name:    name,
		IsSuper: isSupper,
	}
	roleId = manage_factory.ManageId.GenRoleId()
	role.RoleId = roleId
	_, err = manage_factory.ManageRoleStorage.Insert(role)
	if err != nil {
		return
	}
	framework.Info("角色 [" + name + "] [" + util.GetStringValue(roleId) + "] 生成成功")
	return
}

func InitManageRoleUser(roleId int64, userId int64) (err error) {
	finds, err := manage_factory.ManageRoleUserStorage.Query(&module_manage.ManageRoleUser{RoleId: userId, UserId: userId})
	if err != nil {
		return
	}
	if len(finds) > 0 {
		return
	}

	add := &module_manage.ManageRoleUser{RoleId: userId, UserId: userId}
	add.RoleUserId = manage_factory.ManageId.GenRoleId()
	_, err = manage_factory.ManageRoleUserStorage.Insert(add)
	if err != nil {
		return
	}
	framework.Info("角色 [" + util.GetStringValue(roleId) + "] 用户 [" + util.GetStringValue(userId) + "] 绑定成功")
	return
}
