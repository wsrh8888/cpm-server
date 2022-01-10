package cpm

import (
	"cpm/global"
	"cpm/model/cpm"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type language struct{}

var Language = new(language)

/**
 * @description: 表名
 */
func (u *language) TableName() string {
	return "sys_users"
}

func (u *language) Initialize() error {
	entities := []cpm.CpmProjectLanguage{
		{LanguageId: "1000", LanguageName: "vue2"},
		{LanguageId: "2000", LanguageName: "vue3"},
		{LanguageId: "3000", LanguageName: "react"},
	}
	if err := global.CPM_DB.Create(&entities).Error; err != nil {
		return errors.Wrap(err, u.TableName()+"表数据初始化失败!")
		// return errors.Wrap(err, u.TableName()+"表数据初始化失败!")
	}
	return nil
}

func (u *language) CheckDataExist() bool {
	if errors.Is(global.CPM_DB.Where("language_id = ?", "1000").First(&cpm.CpmProjectLanguage{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
