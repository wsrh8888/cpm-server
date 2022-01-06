package request

// User register structure
type Register struct {
	Email     string `json:"email"`
	Password  string `json:"passWord"`
	NickName  string `json:"nickName" gorm:"default:'QMPlusUser'"`
	HeaderImg string `json:"headerImg" gorm:"default:'https://qmplusimg.henrongyi.top/gva_header.jpg'"`
}

// User login structure
type Login struct {
	Email     string `json:"email"`     // 用户名
	Password  string `json:"password"`  // 密码
	Captcha   string `json:"captcha"`   // 验证码
	CaptchaId string `json:"captchaId"` // 验证码ID
}
