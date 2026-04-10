// 文件由 TeamIDE | coos 生成，请勿修改文件内容！通过 [TeamIDE:teamide@163.com] 的 [models:module_manage/storage/manage_permission.coos] 在 [2026-04-10 10:50] 生成

package module_manage

type ManagePermissionAddRequest struct {
	*ManagePermission
}

type ManagePermissionAddResponse struct {
}

type ManagePermissionDeleteRequest struct {
	PermissionIds []int64 `json:"permissionIds"`
	RoleIds       []int64 `json:"roleIds"`
	UserIds       []int64 `json:"userIds"`
}

type ManagePermissionQueryRequest struct {
	PermissionIds []int64 `json:"permissionIds"`
	RoleIds       []int64 `json:"roleIds"`
	UserIds       []int64 `json:"userIds"`
}

type IManagePermissionService interface {
	Add(in *ManagePermissionAddRequest) (res *ManagePermissionAddResponse, err error)
	Query(in *ManagePermissionQueryRequest) (res []*ManagePermission, err error)
	QueryByIds(ids []int64) (res []*ManagePermission, err error)
	QueryByRoleIds(roleIds []int64) (res []*ManagePermission, err error)
	QueryByUserIds(userIds []int64) (res []*ManagePermission, err error)
	Delete(in *ManagePermissionDeleteRequest) (err error)
}

type ManagePermission struct {
	PermissionId    int64  `json:"permissionId" column:"permission_id"`
	RoleId          int64  `json:"roleId" column:"role_id"`
	UserId          int64  `json:"userId" column:"user_id"`
	PermissionType  string `json:"permissionType" column:"permission_type"`
	Permission      string `json:"permission" column:"permission"`
	Authorizable    int    `json:"authorizable" column:"authorizable"`
	StartAt         int64  `json:"startAt" column:"start_at"`
	ExpiredAt       int64  `json:"expiredAt" column:"expired_at"`
	ExpiredDuration int64  `json:"expiredDuration" column:"expired_duration"`
	CreateAt        int64  `json:"createAt" column:"create_at"`
	UpdateAt        int64  `json:"updateAt" column:"update_at"`
}

func (this_ *ManagePermission) GetTableName() string {
	return "manage_permission"
}

func (this_ *ManagePermission) GetPrimaryKey() []string {
	return []string{"permission_id"}
}

type IManagePermissionStorage interface {
	Insert(in *ManagePermission) (res int64, err error)
	QueryByIds(ids []int64) (res []*ManagePermission, err error)
	QueryByRoleId(roleId int64) (res []*ManagePermission, err error)
	QueryByRoleIds(roleIds []int64) (res []*ManagePermission, err error)
	QueryByUserId(userId int64) (res []*ManagePermission, err error)
	QueryByUserIds(userIds []int64) (res []*ManagePermission, err error)
	Query(in *ManagePermission) (res []*ManagePermission, err error)
	DeleteByIds(ids []int64) (res int64, err error)
	DeleteByRoleIds(roleIds []int64) (res int64, err error)
	DeleteByUserIds(userIds []int64) (res int64, err error)
}
