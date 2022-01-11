package cpm

import (
	v1 "cpm/api/v1"

	"github.com/gin-gonic/gin"
)

type VersionRouter struct{}

func (*VersionRouter) InitCpmVersion(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("cpmVersion")
	baseApi := v1.ApiGroupApp.CpmApiGroup
	{
		baseRouter.POST("addCpmVersion", baseApi.AddCpmVersion)
		baseRouter.POST("getCpmVersion", baseApi.GetCpmVersion)
	}
	return baseRouter
}
