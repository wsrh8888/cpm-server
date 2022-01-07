package cpm

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProjectApi struct{}

// @Tags Menu
// @Summary 新增项目
// @Param data body system.SysBaseMenu true "路由path, 父菜单ID, 路由name, 对应前端文件路径, 排序标记"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"添加成功"}"
// @Router /menu/addBaseMenu [post]
func (*ProjectApi) AddCpmProject(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ok": 111,
	})
}
