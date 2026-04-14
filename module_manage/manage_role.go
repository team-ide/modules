// 文件由 TeamIDE | coos 生成，请勿修改文件内容！通过 [TeamIDE:teamide@163.com] 的 [models:module_manage/storage/manage_role.coos] 在 [2026-04-14 14:18] 生成

package module_manage

type ManageRoleAddRequest struct {
	*ManageRole
}

type ManageRoleAddResponse struct {
}

type ManageRoleListRequest struct {
	*ManageRole
}

type ManageRoleListResponse struct {
	List []*ManageRole `json:"list"`
}

type ManageRoleDeleteRequest struct {
	RoleIds []int64 `json:"roleIds"`
}

type ManageUserRoleRequest struct {
	UserId int64 `json:"userId"`
}

type ManageAddRoleUsersResponse struct {
	RoleId  int64   `json:"roleId"`
	UserIds []int64 `json:"userIds"`
}

type ManageAddRoleUserResponse struct {
	RoleId int64 `json:"roleId"`
	UserId int64 `json:"userId"`
}

type IManageRoleService interface {
	Add(in *ManageRoleAddRequest) (res *ManageRoleAddResponse, err error)
	List(in *ManageRoleListRequest) (res *ManageRoleListResponse, err error)
	Delete(roleIds []int64) (err error)
	GetUserRoles(userId int64) (res []*ManageRole, err error)
	AddRoleUsers(roleId int64, userIds []int64) (err error)
	AddRoleUser(roleId int64, userId int64) (err error)
}

type ManageRole struct {
	RoleId   int64  `json:"roleId" column:"role_id"`
	Name     string `json:"name" column:"name"`
	IsSuper  int8   `json:"isSuper" column:"is_super"`
	CreateAt int64  `json:"createAt" column:"create_at"`
	UpdateAt int64  `json:"updateAt" column:"update_at"`
}

func (this_ *ManageRole) GetTableName() string {
	return "manage_role"
}

func (this_ *ManageRole) GetPrimaryKey() []string {
	return []string{"role_id"}
}

type IManageRoleStorage interface {
	GetById(id int64) (res *ManageRole, err error)
	GetByIds(ids []int64) (res []*ManageRole, err error)
	Insert(in *ManageRole) (res int64, err error)
	Update(in *ManageRole) (res int64, err error)
	DeleteByIds(ids []int64) (res int64, err error)
	Query(in *ManageRole) (res []*ManageRole, err error)
	Page(in *ManageRole, pageNo int64, pageSize int64) (res []*ManageRole, err error)
	Count(in *ManageRole) (res int64, err error)
}

type ManageRoleUser struct {
	RoleUserId int64 `json:"roleUserId" column:"role_user_id"`
	RoleId     int64 `json:"roleId" column:"role_id"`
	UserId     int64 `json:"userId" column:"user_id"`
	CreateAt   int64 `json:"createAt" column:"create_at"`
	UpdateAt   int64 `json:"updateAt" column:"update_at"`
}

func (this_ *ManageRoleUser) GetTableName() string {
	return "manage_role_user"
}

func (this_ *ManageRoleUser) GetPrimaryKey() []string {
	return []string{"role_user_id"}
}

type IManageRoleUserStorage interface {
	Insert(in *ManageRoleUser) (res int64, err error)
	DeleteByIds(ids []int64) (res int64, err error)
	QueryByRoleId(roleId int64) (res []*ManageRoleUser, err error)
	QueryRoleUserIds(roleId int64) (res []int64, err error)
	QueryByUserId(userId int64) (res []*ManageRoleUser, err error)
	QueryUserRoleIds(userId int64) (res []int64, err error)
	Query(in *ManageRoleUser) (res []*ManageRoleUser, err error)
	QueryByRoleIdAndUserId(roleId int64, userId int64) (res []*ManageRoleUser, err error)
}
