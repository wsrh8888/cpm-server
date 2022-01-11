package cpm

import (
	"cpm/global"
	"cpm/model/cpm"
	"errors"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

func (cpmService *CpmService) AddImport(cpmProject cpm.CpmProject) (cpmInter cpm.CpmProject, err error) {
	var project cpm.CpmProject
	if !errors.Is(global.CPM_DB.Where("name = ?", cpmProject.Name).First(&project).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		println("项目已经存在")
		return cpmInter, errors.New("项目已经存在")
	}
	cpmProject.UUID = uuid.NewV4()
	err = global.CPM_DB.Create(&cpmProject).Error
	return cpmProject, err
}
