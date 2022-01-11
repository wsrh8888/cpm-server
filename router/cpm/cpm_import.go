package cpm

import (
	v1 "cpm/api/v1"

	"github.com/gin-gonic/gin"
)

type VersionImport struct{}

func (*VersionImport) InitCpmImport(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("cpmImport")
	baseApi := v1.ApiGroupApp.CpmApiGroup
	{
		baseRouter.POST("addCpmVersion", baseApi.AddCpmVersion)
		baseRouter.POST("getCpmVersion", baseApi.GetCpmVersion)
	}
	return baseRouter
}
