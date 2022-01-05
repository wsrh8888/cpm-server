package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	project_group := router.Group("/project")
	user_group := router.Group("/user")

	RouterProject(project_group)
	RouterUser(user_group)
}
