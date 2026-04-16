// 文件由 TeamIDE | coos 生成，请勿修改文件内容！通过 [TeamIDE:teamide@163.com] 的 [models:module_manage/storage/manage_log.coos] 在 [2026-04-16 14:18] 生成

package module_manage

type ManageLogPageRequest struct {
	*ManageLog
	PageNo   int64 `json:"pageNo"`
	PageSize int64 `json:"pageSize"`
}

type ManageLogPageResponse struct {
	Total int64        `json:"total"`
	List  []*ManageLog `json:"list"`
}

type ManageLogDeleteRequest struct {
	LogIds []int64 `json:"logIds"`
}

type IManageLogService interface {
	Add(in *ManageLog) (err error)
	Page(in *ManageLogPageRequest) (res *ManageLogPageResponse, err error)
	Delete(in *ManageLogDeleteRequest) (err error)
}

type ManageLog struct {
	StartTimeBefore int64  `json:"startTimeBefore"`
	StartTimeAfter  int64  `json:"startTimeAfter"`
	LogId           int64  `json:"logId" column:"log_id"`
	LoginId         int64  `json:"loginId" column:"login_id"`
	UserId          int64  `json:"userId" column:"user_id"`
	UserName        string `json:"userName" column:"user_name"`
	UserAccount     string `json:"userAccount" column:"user_account"`
	Ip              string `json:"ip" column:"ip"`
	Path            string `json:"path" column:"path"`
	Comment         string `json:"comment" column:"comment"`
	Method          string `json:"method" column:"method"`
	Param           string `json:"param" column:"param"`
	Data            string `json:"data" column:"data"`
	UserAgent       string `json:"userAgent" column:"user_agent"`
	Error           string `json:"error" column:"error"`
	UseTime         int    `json:"useTime" column:"use_time"`
	StartAt         int64  `json:"startAt" column:"start_at"`
	EndAt           int64  `json:"endAt" column:"end_at"`
	CreateAt        int64  `json:"createAt" column:"create_at"`
}

func (this_ *ManageLog) GetTableName() string {
	return "manage_log"
}

func (this_ *ManageLog) GetPrimaryKey() []string {
	return []string{"log_id"}
}

type IManageLogStorage interface {
	Insert(in *ManageLog) (res int64, err error)
	DeleteByIds(ids []int64) (res int64, err error)
	Query(in *ManageLog) (res []*ManageLog, err error)
	Page(in *ManageLog, pageNo int64, pageSize int64) (res []*ManageLog, err error)
	Count(in *ManageLog) (res int64, err error)
}
