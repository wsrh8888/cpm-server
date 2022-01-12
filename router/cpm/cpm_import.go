package cpm

import (
	v1 "cpm/api/v1"

	"github.com/gin-gonic/gin"
)

type ImportRouter struct{}

func (*ImportRouter) InitCpmImport(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("cpmImport")
	baseApi := v1.ApiGroupApp.CpmApiGroup
	{
		baseRouter.POST("getCpmImport", baseApi.GetCpmImport)
		baseRouter.POST("getCpmImportUrlList", baseApi.GetCpmImportUrlList)
	}
	return baseRouter
}
