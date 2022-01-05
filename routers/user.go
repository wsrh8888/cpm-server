package routers

import (
	"cpm/controller"

	"github.com/gin-gonic/gin"
)

func RouterUser(router *gin.RouterGroup) {
	router.POST("/get_all", controller.SaveTask)
}
