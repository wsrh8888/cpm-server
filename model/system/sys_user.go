package system

import (
	"cpm/global"

	uuid "github.com/satori/go.uuid"
)

type SysUser struct {
	global.CPM_MODEL
	UUID      uuid.UUID `json:"uuid" gorm:"comment:用户UUID"`                                                           // 用户UUID
	Email     string    `json:"email" gorm:"comment:邮箱"`                                                              // 用户登录名
	Password  string    `json:"-"  gorm:"comment:用户登录密码"`                                                             // 用户登录密码
	NickName  string    `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                                            // 用户昵称
	HeaderImg string    `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"` // 用户头像
}
