package routers

import (
	"cpm/controller"

	"github.com/gin-gonic/gin"
)

func RouterProject(router *gin.RouterGroup) {
	router.POST("/get_all", controller.SaveTask)
}
