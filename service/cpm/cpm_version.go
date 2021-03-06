package cpm

import (
	"cpm/global"
	"cpm/model/cpm"
	"cpm/model/cpm/request"
	"errors"

	"gorm.io/gorm"
)

type CpmVersionService struct{}

/**
 * @description: 创建项目
 * @param {*}
 * @return {*}
 */
func (cpmService *CpmVersionService) AddVersion(cpmVersion request.CpmVersionAdd) (cpmInter interface{}, err error) {
	var version cpm.CpmVersion
	var importList []cpm.CpmImport
	// print(cpmVersion.ProjectId)
	// println(cpmVersion.ProjectId, cpmVersion.Version)
	if !errors.Is(global.CPM_DB.Where("project_id = ? and version = ?", cpmVersion.ProjectId, cpmVersion.Version).First(&version).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return cpmInter, errors.New("该版本已经发布")
	}
	err = global.CPM_DB.Table("cpm_versions").Create(&cpmVersion).Error

	for _, v := range cpmVersion.CpmImportList {
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

func (cpmService *CpmService) GetVersion(version cpm.CpmVersion) (cpmInter interface{}, err error) {
	// var versionList []cpm.CpmVersion
	// err = global.CPM_DB.Where("project_id = ?", version.ProjectId).Preload("Publish").Find(&versionList).Error
	// var importList []response.VersionList
	// for _, v := range versionList {
	// 	importList = append(importList, response.VersionList{
	// 		Version:   v.Version,
	// 		Publisher: v.Publish.NickName,
	// 		CreatedAt: v.CreatedAt,
	// 	})
	// }
	return cpmInter, err
}
