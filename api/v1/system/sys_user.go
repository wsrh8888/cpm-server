package system

import (
	"cpm/model/common/response"
	"cpm/model/system"
	systemReq "cpm/model/system/request"
	systemRes "cpm/model/system/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (b *BaseApi) Login(c *gin.Context) {
	var l systemReq.Login
	if errBind := c.ShouldBindJSON(&l); errBind != nil {
		response.FailWithMessage(errBind.Error(), c)
		return
	}
	if store.Verify(l.CaptchaId, l.Captcha, true) {
		u := &system.SysUser{Email: l.Email, Password: l.Password}
		if err, user := userService.Login(u); err != nil {
			response.FailWithMessage("用户名不存在或者密码错误", c)

		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":   0,
				"result": user,
			})
			// b.tokenNext(c, *user)
		}
	} else {
		response.FailWithMessage("验证码错误", c)

	}
}

func (b *BaseApi) Register(c *gin.Context) {
	var r systemReq.Register
	if errBind := c.ShouldBindJSON(&r); errBind != nil {
		response.FailWithMessage(errBind.Error(), c)
		return
	}
	user := &system.SysUser{Email: r.Email, NickName: r.NickName, Password: r.Password, HeaderImg: r.HeaderImg}
	err, userReturn := userService.Register(*user)
	if err != nil {
		response.FailWithDetailed(nil, err.Error(), c)
	} else {
		response.OkWithDetailed(systemRes.SysUserResponse{User: userReturn}, "注册成功", c)
	}
}
