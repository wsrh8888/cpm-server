package cpm

import (
	"cpm/model/common/request"
	"cpm/model/common/response"
	cpmRequestModel "cpm/model/cpm/request"

	"cpm/model/cpm"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProjectApi struct{}

func (*ProjectApi) GetCpmProject(c *gin.Context) {
	var pageInfo cpmRequestModel.SysDictionaryDetailSearch
	_ = c.ShouldBindJSON(&pageInfo)

	if list, total, err := cpmService.GetProject(pageInfo); err != nil {
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
	// var cpmProject cpm.CpmProject
	// _ = c.ShouldBindJSON(&cpmProject)
	// if info, err := cpmService.GetProject(cpmProject.ID); err != nil {
	// 	response.FailWithMessage(err.Error(), c)
	// } else {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"code":   0,
	// 		"result": info,
	// 	})
	// }
}

// @Tags Menu
// @Summary 新增项目
// @Param data body system.SysBaseMenu true "路由path, 父菜单ID, 路由name, 对应前端文件路径, 排序标记"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"添加成功"}"
// @Router /menu/addBaseMenu [post]
func (*ProjectApi) AddCpmProject(c *gin.Context) {
	var cpmProject cpm.CpmProject
	_ = c.ShouldBindJSON(&cpmProject)
	if info, err := cpmService.AddProject(cpmProject); err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":   0,
			"result": info,
		})
	}
}

// @Tags Menu
// @Summary 删除项目
// @Param data body system.SysBaseMenu true "路由path, 父菜单ID, 路由name, 对应前端文件路径, 排序标记"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"添加成功"}"
// @Router /menu/addBaseMenu [post]
func (*ProjectApi) DeleteCpmProject(c *gin.Context) {
	var uuidInfo request.GetByUuid
	_ = c.ShouldBindJSON(&uuidInfo)
	if err := cpmService.DeleteProject(uuidInfo.UUID); err != nil {
		response.FailWithMessage("删除失败", c)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":   0,
			"result": "删除成功",
		})
	}
}
