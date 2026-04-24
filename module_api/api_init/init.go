package api_init

import (
	"github.com/team-ide/framework/db"
	"github.com/team-ide/framework/httpx"
	"github.com/team-ide/modules/module_api/api_factory"
	"github.com/team-ide/modules/module_api/api_service"
	"github.com/team-ide/modules/module_api/api_storage"
)

func InitFactoryService(dbService db.IService, httpService httpx.IService) {

	api_factory.OauthAccessTokenCache = api_service.NewOauthAccessTokenCache()

	api_factory.OauthAccessTokenStorage = api_storage.NewOauthAccessTokenStorage(dbService)

	api_factory.OauthAccessTokenService = api_service.NewOauthAccessTokenService(httpService)

	return
}
