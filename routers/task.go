package routers

import (
	"cpm/controller"

	"github.com/gin-gonic/gin"
)

func Router(router *gin.RouterGroup) {
	router.GET("/get_all", controller.SavePackage)
}
