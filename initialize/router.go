package initialize

import "github.com/gin-gonic/gin"

// import "github.com/gin-gonic/gin"

func Routers() *gin.Engine {
	Router := gin.Default()

	PublicGroup := Router.Group("")
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}
	return Router
}
