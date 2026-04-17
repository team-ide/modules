// 文件由 TeamIDE | coos 生成，请勿修改文件内容！通过 [TeamIDE:teamide@163.com] 的 [models:module_manage/storage/manage_login.coos] 在 [2026-04-17 09:44] 生成

package module_manage

type ManageLoginPageRequest struct {
	*ManageLogin
	PageNo   int64 `json:"pageNo"`
	PageSize int64 `json:"pageSize"`
}

type ManageLoginPageResponse struct {
	Total int64          `json:"total"`
	List  []*ManageLogin `json:"list"`
}

type ManageLoginDeleteRequest struct {
	LoginIds []int64 `json:"loginIds"`
}

type IManageLoginService interface {
	Add(in *ManageLogin) (err error)
	Page(in *ManageLoginPageRequest) (res *ManageLoginPageResponse, err error)
	Delete(in *ManageLoginDeleteRequest) (err error)
}

type ManageLogin struct {
	LoginId       int64  `json:"loginId" column:"login_id"`
	UserId        int64  `json:"userId" column:"user_id"`
	UserAccount   string `json:"userAccount" column:"user_account"`
	UserName      string `json:"userName" column:"user_name"`
	LoginIp       string `json:"loginIp" column:"login_ip"`
	SourceType    int    `json:"sourceType" column:"source_type"`
	SourceInfo    string `json:"sourceInfo" column:"source_info"`
	Token         string `json:"token" column:"token"`
	Status        int    `json:"status" column:"status"`
	LoginAt       int64  `json:"loginAt" column:"login_at"`
	LogoutAt      int64  `json:"logoutAt" column:"logout_at"`
	CreateAt      int64  `json:"createAt" column:"create_at"`
	UpdateAt      int64  `json:"updateAt" column:"update_at"`
	DeleteAt      int64  `json:"deleteAt" column:"delete_at"`
	UseAt         int64  `json:"useAt" column:"use_at"`
	LoginAtBefore int64  `json:"loginAtBefore"`
	LoginAtAfter  int64  `json:"loginAtAfter"`
}

func (this_ *ManageLogin) GetTableName() string {
	return "manage_login"
}

func (this_ *ManageLogin) GetPrimaryKey() []string {
	return []string{"login_id"}
}

type IManageLoginStorage interface {
	GetById(id int64) (res *ManageLogin, err error)
	GetByToken(token string) (res *ManageLogin, err error)
	Insert(in *ManageLogin) (res int64, err error)
	Update(in *ManageLogin) (res int64, err error)
	Logout(id int64) (res int64, err error)
	Use(id int64) (res int64, err error)
	DeleteByIds(ids []int64) (res int64, err error)
	Query(in *ManageLogin) (res []*ManageLogin, err error)
	Page(in *ManageLogin, pageNo int64, pageSize int64) (res []*ManageLogin, err error)
	Count(in *ManageLogin) (res int64, err error)
	RemoveByIds(ids []int64) (res int64, err error)
}
