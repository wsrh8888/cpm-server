package initialize

import (
	"cpm/middleware"
	"cpm/router"

	"github.com/gin-gonic/gin"
)

// import "github.com/gin-gonic/gin"

func Routers() *gin.Engine {
	Router := gin.Default()
	systemRouter := router.RouterGroupApp.System
	cpmRouter := router.RouterGroupApp.Cpm
	// 解决跨域
	Router.Use(middleware.Cors()) // 直接放行全部跨域请求

	PublicGroup := Router.Group("")
	{
		systemRouter.InitBaseRouter(PublicGroup)
		systemRouter.InitInitRouter(PublicGroup)
	}

	PrivateGroup := Router.Group("")
	{
		cpmRouter.InitCpmProject(PrivateGroup)
	}
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}
	return Router
}
