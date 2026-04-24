// 文件由 TeamIDE | coos 生成，请勿修改文件内容！通过 [TeamIDE:teamide@163.com] 的 [models:module_manage/storage/manage_user.coos] 在 [2026-04-24 17:44] 生成

package module_manage

type ManageUserAddRequest struct {
	*ManageUser
	Password string `json:"password"`
}

type ManageUserAddResponse struct {
}

type ManageUserListRequest struct {
	*ManageUser
}

type ManageUserListResponse struct {
	List []*ManageUser `json:"list"`
}

type ManageUserPageRequest struct {
	*ManageUser
	PageNo   int64 `json:"pageNo"`
	PageSize int64 `json:"pageSize"`
}

type ManageUserPageResponse struct {
	Total int64         `json:"total"`
	List  []*ManageUser `json:"list"`
}

type ManageUserDeleteRequest struct {
	UserIds []int64 `json:"userIds"`
}

type IManageUserService interface {
	Add(in *ManageUserAddRequest) (res *ManageUserAddResponse, err error)
	List(in *ManageUserListRequest) (res *ManageUserListResponse, err error)
	Page(in *ManageUserPageRequest) (res *ManageUserPageResponse, err error)
	Delete(userIds []int64) (err error)
	Remove(userIds []int64) (err error)
}

type ManageUser struct {
	UserId    int64  `json:"userId" column:"user_id"`
	Name      string `json:"name" column:"name"`
	Account   string `json:"account" column:"account"`
	Salt      string `json:"salt" column:"salt"`
	Password  string `json:"password" column:"password"`
	Status    int    `json:"status" column:"status"`
	CreateAt  int64  `json:"createAt" column:"create_at"`
	UpdateAt  int64  `json:"updateAt" column:"update_at"`
	DeleteAt  int64  `json:"deleteAt" column:"delete_at"`
	DisableAt int64  `json:"disableAt" column:"disable_at"`
}

func (this_ *ManageUser) GetTableName() string {
	return "manage_user"
}

func (this_ *ManageUser) GetPrimaryKey() []string {
	return []string{"user_id"}
}

type IManageUserStorage interface {
	GetById(id int64) (res *ManageUser, err error)
	GetByIds(ids []int64) (res []*ManageUser, err error)
	Insert(in *ManageUser) (res int64, err error)
	Update(in *ManageUser) (res int64, err error)
	Disable(id int64) (res int64, err error)
	Enable(id int64) (res int64, err error)
	UpdatePassword(id int64, salt string, password string) (res int64, err error)
	QueryByAccount(account string) (res []*ManageUser, err error)
	CountByAccount(account string) (res int64, err error)
	DeleteByIds(ids []int64) (res int64, err error)
	Query(in *ManageUser) (res []*ManageUser, err error)
	Page(in *ManageUser, pageNo int64, pageSize int64) (res []*ManageUser, err error)
	Count(in *ManageUser) (res int64, err error)
	RemoveByIds(ids []int64) (res int64, err error)
}
