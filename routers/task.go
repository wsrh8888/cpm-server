package routers

import (
	"cpm/controller"

	"github.com/gin-gonic/gin"
)

func Router(router *gin.RouterGroup) {
	router.POST("/get_all", controller.SaveTask)
}
