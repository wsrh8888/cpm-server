package cpm

import (
	"cpm/global"
	"cpm/model/cpm"
	"errors"

	"gorm.io/gorm"
)

type CpmVersionService struct{}

/**
 * @description: 创建项目
 * @param {*}
 * @return {*}
 */
func (cpmService *CpmVersionService) AddVersion(cpmVersion cpm.CpmVersion) (err error, cpmInter cpm.CpmVersion) {
	var version cpm.CpmVersion
	if !errors.Is(global.CPM_DB.Where("version_id = ?", cpmVersion.VersionId).First(&version).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("该版本已经发布"), cpmInter
	}
	err = global.CPM_DB.Create(&cpmVersion).Error
	return err, cpmVersion
}

func (cpmService *CpmService) GetVersion(id string) (err error, cpmInter cpm.CpmVersion) {

	err = global.CPM_DB.Where("version_id = ?", id).Preload("Publish").First(&cpmInter).Error
	return
}
