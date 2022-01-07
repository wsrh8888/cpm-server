package cpm

import (
	v1 "cpm/api/v1"

	"github.com/gin-gonic/gin"
)

type ProjectRouter struct{}

func (*ProjectRouter) InitCpmProject(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("cpmProject")
	projectApi := v1.ApiGroupApp.CpmApiGroup.ProjectApi
	{
		baseRouter.POST("addCpmProject", projectApi.AddCpmProject)
	}
	return baseRouter
}
