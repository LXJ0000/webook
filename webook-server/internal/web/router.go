package web

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
	"webook-server/internal/repository"
	cache2 "webook-server/internal/repository/cache"
	dao2 "webook-server/internal/repository/dao"
	"webook-server/internal/service"
	"webook-server/internal/web/middleware"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	//跨域解决方案 https://github.com/gin-contrib/cors
	r.Use(cors.New(cors.Config{
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true, // 允许带 cookie
		AllowOriginFunc: func(origin string) bool {
			if strings.Contains(origin, "http://localhost") {
				return true
			}
			return strings.Contains(origin, "www.example.com")
		},
		ExposeHeaders: []string{"token"},
		MaxAge:        12 * time.Hour,
	}))

	initUserRouter(r)

	return r
}

func initUserRouter(r *gin.Engine) {
	dao := dao2.NewUserDao(db)
	cache := cache2.NewUserCache(redisClient)
	
	repo := repository.NewUserRepository(dao, cache)
	svc := service.NewUserService(repo)
	user := NewUserHandler(svc)

	userGroup := r.Group("/user")
	userGroup.POST("/login", user.Login)
	userGroup.POST("/register", user.Register)

	authUserGroup := userGroup.Use(middleware.JwtAuthMiddleware())
	authUserGroup.GET("/", user.Profile)
}
