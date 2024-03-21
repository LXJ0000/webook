//go:build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"webook-server/internal/handler"
	"webook-server/internal/repository"
	"webook-server/internal/repository/cache"
	"webook-server/internal/repository/dao"
	"webook-server/ioc"
)

func InitWebServer() *gin.Engine {
	wire.Build(
		//第三方依赖
		ioc.InitDB, ioc.InitRedis,
		ioc.InitSMSService,
		//dao
		dao.NewUserDao, dao.NewMenuDao,
		//cache
		cache.NewUserCache, cache.NewCodeCache,
		//repository
		repository.NewUserRepository, repository.NewCodeRepository, repository.NewMenuRepository,
		//service
		handler.NewUserHandler, handler.NewMenuHandler,

		ioc.InitGinMiddlewares,
		ioc.InitWebServer,
	)
	return gin.Default()
}
