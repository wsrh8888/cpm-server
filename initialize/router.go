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
	PublicGroup := Router.Group("")

	// 解决跨域
	Router.Use(middleware.Cors()) // 直接放行全部跨域请求

	{
		systemRouter.InitBaseRouter(PublicGroup)
	}
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}
	return Router
}
