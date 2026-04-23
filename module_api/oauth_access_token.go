// 文件由 TeamIDE | coos 生成，请勿修改文件内容！通过 [TeamIDE:teamide@163.com] 的 [models:module_api/storage/oauth_access_token.coos] 在 [2026-04-23 16:33] 生成

package module_api

type IOauthAccessTokenService interface {
	GetAccessToken(name string) (res string, err error)
}

type OauthAccessToken struct {
	Name         string `json:"name" column:"name"`
	PlatformType string `json:"platformType" column:"platform_type"`
	OauthUrl     string `json:"oauthUrl" column:"oauth_url"`
	AppId        string `json:"appId" column:"app_id"`
	AppSecret    string `json:"appSecret" column:"app_secret"`
	Extends      string `json:"extends" column:"extends"`
	AccessToken  string `json:"accessToken" column:"access_token"`
	ExpiresIn    int64  `json:"expiresIn" column:"expires_in"`
	ExpiresAt    int64  `json:"expiresAt" column:"expires_at"`
	Status       int    `json:"status" column:"status"`
	CreatedAt    int64  `json:"createdAt" column:"created_at"`
	UpdatedAt    int64  `json:"updatedAt" column:"updated_at"`
	DeletedAt    int64  `json:"deletedAt" column:"deleted_at"`
	DisabledAt   int64  `json:"disabledAt" column:"disabled_at"`
}

func (this_ *OauthAccessToken) GetTableName() string {
	return "oauth_access_token"
}

func (this_ *OauthAccessToken) GetPrimaryKey() []string {
	return []string{"name"}
}

type IOauthAccessTokenStorage interface {
	GetByName(name string) (res *OauthAccessToken, err error)
	Insert(in *OauthAccessToken) (res int64, err error)
	Update(in *OauthAccessToken) (res int64, err error)
	UpdateAccessToken(name string, accessToken string, expiresIn int64) (res int64, err error)
	Query(in *OauthAccessToken) (res []*OauthAccessToken, err error)
	Page(in *OauthAccessToken, pageNo int64, pageSize int64) (res []*OauthAccessToken, err error)
}
