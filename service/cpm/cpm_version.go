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
func (cpmService *CpmVersionService) AddVersion(cpmVersion cpm.CpmVersion) (cpmInter cpm.CpmVersion, err error) {
	var version cpm.CpmVersion
	var importList []cpm.CpmImport
	if !errors.Is(global.CPM_DB.Where("project_id = ? and version = ?", cpmVersion.ProjectId, cpmVersion.Version).First(&version).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return cpmInter, errors.New("该版本已经发布")
	}
	err = global.CPM_DB.Create(&cpmVersion).Error
	for _, v := range cpmVersion.ImportList {
		importList = append(importList, cpm.CpmImport{
			ProjectId: cpmVersion.ProjectId,
			Version:   cpmVersion.Version,
			Name:      v.Name,
			List:      v.List,
			Md:        v.Md,
		})
	}
	var err2 error = global.CPM_DB.Create(&importList).Error
	println(err, importList[0].Md)
	return cpmVersion, err2

}

func (cpmService *CpmService) GetVersion(version cpm.CpmVersion) (cpmInter cpm.CpmVersion, err error) {
	err = global.CPM_DB.Where("project_id = ? and version= ?", version.ProjectId, version.Version).Preload("Publish").First(&cpmInter).Error
	return
}
