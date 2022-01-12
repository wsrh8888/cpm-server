package request

// User register structure
type Register struct {
	Email     string `json:"email" binding:"required"`
	Password  string `json:"passWord" binding:"required"`
	NickName  string `json:"nickName" gorm:"default:'QMPlusUser'"  binding:"required"`
	HeaderImg string `json:"headerImg" gorm:"default:'https://qmplusimg.henrongyi.top/gva_header.jpg'"`
}

// User login structure
type Login struct {
	Email     string `json:"email" binding:"required"`    // 邮箱
	Password  string `json:"password" binding:"required"` // 密码
	Captcha   string `json:"captcha"`                     // 验证码
	CaptchaId string `json:"captchaId"`                   // 验证码ID
}
