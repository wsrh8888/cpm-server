package system

import (
	"cpm/global"
	"cpm/model/system"
	"cpm/utils"
	"errors"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type UserService struct{}

func (userService *UserService) Register(u system.SysUser) (err error, userInter system.SysUser) {
	var user system.SysUser
	if !errors.Is(global.CPM_DB.Where("email = ?", u.Email).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("邮箱已注册"), userInter
	}
	u.Password = utils.MD5V([]byte(u.Password))
	u.UUID = uuid.NewV4()
	err = global.CPM_DB.Create(&u).Error
	return err, u
}
