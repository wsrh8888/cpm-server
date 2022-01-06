package system

import (
	systemReq "cpm/model/system/request"

	"github.com/gin-gonic/gin"
)

func (b *BaseApi) Login(c *gin.Context) {

}

func (b *BaseApi) Register(c *gin.Context) {
	var r systemReq.Register
	_ = c.ShouldBindJSON(&r)

}
