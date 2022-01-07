package cpm

import (
	v1 "cpm/api/v1"

	"github.com/gin-gonic/gin"
)

type CpmVersion struct{}

func (*CpmVersion) InitCpmVersion(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("cpmVersion")
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		baseRouter.POST("register", baseApi.Register)
	}
	return baseRouter
}
