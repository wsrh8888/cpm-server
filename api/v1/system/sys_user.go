package system

import (
	"cpm/model/common/response"
	"cpm/model/system"
	systemReq "cpm/model/system/request"
	systemRes "cpm/model/system/response"

	"github.com/gin-gonic/gin"
)

func (b *BaseApi) Login(c *gin.Context) {

}

func (b *BaseApi) Register(c *gin.Context) {
	var r systemReq.Register
	_ = c.ShouldBindJSON(&r)
	user := &system.SysUser{Email: r.Email, NickName: r.NickName, Password: r.Password, HeaderImg: r.HeaderImg}
	println(user.Email)
	println(r.Email)
	err, userReturn := userService.Register(*user)
	if err != nil {
		response.FailWithDetailed(systemRes.SysUserResponse{User: userReturn}, err.Error(), c)
	} else {
		response.OkWithDetailed(systemRes.SysUserResponse{User: userReturn}, "注册成功", c)
	}
}
