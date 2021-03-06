package cpm

import (
	"cpm/model/common/response"
	"cpm/model/cpm"
	"cpm/model/cpm/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

type VersionApi struct{}

// @Tags Menu
// @Summary 新增项目
// @Param data body system.SysBaseMenu true "路由path, 父菜单ID, 路由name, 对应前端文件路径, 排序标记"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"添加成功"}"
// @Router /menu/addBaseMenu [post]
func (v VersionApi) AddCpmVersion(c *gin.Context) {
	var cpmVersion request.CpmVersionAdd
	if errBind := c.ShouldBindJSON(&cpmVersion); errBind != nil {
		response.FailWithMessage(errBind.Error(), c)
		return
	}
	if _, err := cpmVersionService.AddVersion(cpmVersion); err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithDetailed(nil, "版本发布成功", c)
	}
}

func (*ProjectApi) GetCpmVersion(c *gin.Context) {
	var cpmVersion cpm.CpmVersion
	_ = c.ShouldBindJSON(&cpmVersion)
	if info, err := cpmService.GetVersion(cpmVersion); err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":   0,
			"result": info,
		})
	}
}
