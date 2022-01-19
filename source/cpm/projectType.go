package cpm

import (
	"cpm/global"
	"cpm/model/cpm"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type projectType struct{}

var Type = new(projectType)

/**
 * @description: 表名
 */
func (u *projectType) TableName() string {
	return "sys_users"
}

func (u *projectType) Initialize() error {
	entities := []cpm.CpmType{
		{ID: "10000", Name: "组件库"},
		{ID: "20000", Name: "静态资源"},
		{ID: "30000", Name: "AB Test"},
	}
	if err := global.CPM_DB.Create(&entities).Error; err != nil {
		return errors.Wrap(err, u.TableName()+"表数据初始化失败!")
		// return errors.Wrap(err, u.TableName()+"表数据初始化失败!")
	}
	return nil
}

func (u *projectType) CheckDataExist() bool {
	if errors.Is(global.CPM_DB.Where("id = ?", "1000").First(&cpm.CpmType{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
