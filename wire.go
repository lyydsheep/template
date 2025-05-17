//go:build wireinject

package main

import (
	"your-module-name/api/controller"
	"your-module-name/api/router"
	"your-module-name/common/middleware"
	"your-module-name/dal/cache"
	"your-module-name/dal/dao"
	"your-module-name/logic/appService"
	"your-module-name/logic/domainService"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitializeApp() *gin.Engine {
	wire.Build(router.RegisterRoutersAndMiddleware,
		middleware.GetHandlerFunc, controller.NewBuildController,
		wire.Bind(new(appService.DemoAppService), new(*appService.DemoAppServiceV1)), appService.NewDemoAppServiceV1,
		wire.Bind(new(domainService.DemoDomainService), new(*domainService.DemoDomainServiceV1)), domainService.NewDemoDomainServiceV1,
		wire.Bind(new(dao.DemoDAO), new(*dao.DemoDAOV1)), wire.Bind(new(cache.DemoCache), new(*cache.DemoCacheV1)),
		dao.NewDemoDAO, cache.NewCacheV1,
	)
	return nil
}
