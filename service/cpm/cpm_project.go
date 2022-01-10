package cpm

import (
	"cpm/global"
	"cpm/model/cpm"
	"errors"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type CpmService struct{}

/**
 * @description: 创建项目
 * @param {*}
 * @return {*}
 */
func (cpmService *CpmService) AddProject(cpmProject cpm.CpmProject) (err error, cpmInter cpm.CpmProject) {
	var project cpm.CpmProject
	if !errors.Is(global.CPM_DB.Where("name = ?", cpmProject.Name).First(&project).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		println("项目已经存在")
		return errors.New("项目已经存在"), cpmInter
	}
	cpmProject.UUID = uuid.NewV4()
	err = global.CPM_DB.Create(&cpmProject).Error
	return err, cpmProject
}

/**
 * @description: 删除项目
 * @param {*}
 * @return {*}
 */
func (cpmService *CpmService) DeleteProject(id uuid.UUID) (err error) {
	err = global.CPM_DB.Where("uuid = ?", id).Unscoped().Delete(&cpm.CpmProject{}).Error
	return err
}
