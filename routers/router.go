package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	task_group := router.Group("/task")
	Router(task_group)
}
