package system

import (
	"cpm/global"
	"cpm/model/common/response"

	"github.com/gin-gonic/gin"
)

type DBApi struct{}

// InitDB
// @Tags InitDB
// @Summary 初始化用户数据库
// @Produce  application/json
// @Param data body request.InitDB true "初始化数据库参数"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"自动创建数据库成功"}"
// @Router /init/initdb [post]
func (i *DBApi) InitDB(c *gin.Context) {
	if global.CPM_DB != nil {
		response.FailWithMessage("已存在数据库配置", c)
		return
	}
	if err := initDBService.InitDB(); err != nil {
		response.FailWithMessage("自动创建数据库失败，请查看后台日志，检查后在进行初始化", c)
		return
	}
	response.OkWithData("自动创建数据库成功", c)
}

// CheckDB
// @Tags CheckDB
// @Summary 初始化用户数据库
// @Produce  application/json
// @Success 200 {string} string "{"code":0,"data":{},"msg":"探测完成"}"
// @Router /init/checkdb [post]
func (i *DBApi) CheckDB(c *gin.Context) {
	if global.CPM_DB != nil {
		response.OkWithDetailed(gin.H{"needInit": false}, "数据库无需初始化", c)
		return
	} else {
		response.OkWithDetailed(gin.H{"needInit": true}, "前往初始化数据库", c)
		return
	}
}
