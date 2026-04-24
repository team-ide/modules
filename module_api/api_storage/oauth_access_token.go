// 文件由 TeamIDE | coos 生成，请勿修改文件内容！通过 [TeamIDE:teamide@163.com] 的 [models:] 在 [2026-04-24 16:02] 生成

package api_storage

import (
	"context"
	"time"

	"github.com/team-ide/framework/db"
	"github.com/team-ide/modules/module_api"
)

func NewOauthAccessTokenStorage(dbApi db.IService) *OauthAccessTokenStorage {
	res := &OauthAccessTokenStorage{}
	res.dbApi = dbApi
	return res
}

type OauthAccessTokenStorage struct {
	dbApi db.IService
}

func (this_ *OauthAccessTokenStorage) GetByName(name string) (res *module_api.OauthAccessToken, err error) {
	m := this_.dbApi.SqlSelect("oauth_access_token")
	m.CanSelectAll()
	m.Where().Eq("name", name)
	res, err = db.DoQueryOneWithModel[*module_api.OauthAccessToken](context.Background(), m)
	return
}

func (this_ *OauthAccessTokenStorage) Insert(in *module_api.OauthAccessToken) (res int64, err error) {
	in.CreatedAt = time.Now().UnixMilli()
	if in.Status == 0 {
		in.Status = 1
	}
	r, err := this_.dbApi.Insert(context.Background(), in, func(in *db.ModelInsert) {
		in.ValueExclude("updated_at", "deleted_at", "disabled_at")
	})
	if err != nil {
		return
	}
	res, err = r.RowsAffected()
	return
}

func (this_ *OauthAccessTokenStorage) Update(in *module_api.OauthAccessToken) (res int64, err error) {
	in.UpdatedAt = time.Now().UnixMilli()
	r, err := this_.dbApi.Update(context.Background(), in, func(in *db.ModelUpdate) {
	})
	if err != nil {
		return
	}
	res, err = r.RowsAffected()
	return
}

func (this_ *OauthAccessTokenStorage) UpdateAccessToken(name string, accessToken string, expiresIn int64, expiresAt int64) (res int64, err error) {
	m := this_.dbApi.SqlUpdate("oauth_access_token", func(in *db.SqlUpdate) {
	})
	m.Value("access_token", accessToken)
	m.Value("expires_in", expiresIn)
	m.Value("expires_at", expiresAt)
	m.Value("updated_at", time.Now().UnixMilli())
	m.Value("updated_at", time.Now().UnixMilli())
	m.Value("access_token", accessToken)
	m.Value("expires_in", expiresIn)
	m.Value("expires_at", expiresAt)
	m.Where().Eq("name", name)
	r, err := m.Exec(context.Background())
	if err != nil {
		return
	}
	res, err = r.RowsAffected()
	return
}

func (this_ *OauthAccessTokenStorage) Query(in *module_api.OauthAccessToken) (res []*module_api.OauthAccessToken, err error) {
	m := this_.dbApi.ModelSelect(in, func(in *db.ModelSelect) {
		in.CanSelectAll()
	})
	m.Where().AndWhereSql("status != 9")
	res, err = db.DoQueryListWithModel[*module_api.OauthAccessToken](context.Background(), m)
	return
}

func (this_ *OauthAccessTokenStorage) Page(in *module_api.OauthAccessToken, pageNo int64, pageSize int64) (res []*module_api.OauthAccessToken, err error) {
	m := this_.dbApi.ModelSelect(in, func(in *db.ModelSelect) {
		in.CanSelectAll()
	})
	m.Where().AndWhereSql("status != 9")
	res, err = db.DoQueryPageWithModel[*module_api.OauthAccessToken](context.Background(), m, pageNo, pageSize)
	return
}
