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

	// 系统相关不需鉴权的接口
	PublicGroup := Router.Group("")
	{
		systemRouter.InitBaseRouter(PublicGroup)
		systemRouter.InitInitRouter(PublicGroup)
	}

	// 项目相关的接口
	PrivateGroup := Router.Group("")
	{
		cpmRouter.InitCpmProject(PrivateGroup)
		cpmRouter.InitCpmVersion(PrivateGroup)
		cpmRouter.InitCpmImport(PrivateGroup)
	}
	// 健康监测
	{
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}
	return Router
}
