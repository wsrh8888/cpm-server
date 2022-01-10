package system

import (
	"cpm/global"
	"cpm/model/system"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type user struct{}

var User = new(user)

/**
 * @description: 表名
 */
func (u *user) TableName() string {
	return "sys_users"
}

func (u *user) Initialize() error {
	entities := []system.SysUser{
		{UUID: uuid.NewV4(), Email: "751135385@qq.com", Password: "e10adc3949ba59abbe56e057f20f883e", NickName: "超级管理员", HeaderImg: "https://qmplusimg.henrongyi.top/gva_header.jpg"},
		{UUID: uuid.NewV4(), Email: "754308302@qq.com", Password: "e10adc3949ba59abbe56e057f20f883e", NickName: "QMPlusUser", HeaderImg: "https:///qmplusimg.henrongyi.top/1572075907logo.png"},
	}
	if err := global.CPM_DB.Create(&entities).Error; err != nil {
		return errors.Wrap(err, u.TableName()+"表数据初始化失败!")
		// return errors.Wrap(err, u.TableName()+"表数据初始化失败!")
	}
	return nil
}

func (u *user) CheckDataExist() bool {
	if errors.Is(global.CPM_DB.Where("email = ?", "754308302@qq.com").First(&system.SysUser{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
