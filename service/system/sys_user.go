package system

import (
	"cpm/global"
	"cpm/model/system"
	"cpm/utils"
	"errors"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type UserService struct{}

func (userService *UserService) Login(u *system.SysUser) (err error, userInter *system.SysUser) {
	if nil == global.CPM_DB {
		return fmt.Errorf("db not init"), nil
	}
	var user system.SysUser
	u.Password = utils.MD5V([]byte(u.Password))
	println(111)
	println(u.Email, u.Password)
	err = global.CPM_DB.Where("email = ? AND password = ?", u.Email, u.Password).First(&user).Error
	return err, &user
}

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
