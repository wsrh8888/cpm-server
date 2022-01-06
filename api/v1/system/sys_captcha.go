package system

import (
	"cpm/global"
	"cpm/model/common/response"

	systemRes "cpm/model/system/response"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

var store = base64Captcha.DefaultMemStore

type BaseApi struct{}

func (b *BaseApi) Captcha(c *gin.Context) {
	// 也可以自定义参数
	driver := base64Captcha.NewDriverDigit(global.CPM_CONFIG.Captcha.ImgHeight, global.CPM_CONFIG.Captcha.ImgWidth, global.CPM_CONFIG.Captcha.KeyLong, 0.7, 80)
	// 生成base64图片
	cp := base64Captcha.NewCaptcha(driver, store)
	// 获取
	if id, b64s, err := cp.Generate(); err != nil {
		response.FailWithMessage("验证码获取失败", c)
	} else {

		response.OkWithDetailed(systemRes.SysCaptchaResponse{
			CaptchaId:     id,
			PicPath:       b64s,
			CaptchaLength: global.CPM_CONFIG.Captcha.KeyLong,
		}, "验证码获取成功", c)
	}
}
