package cpm

import (
	"cpm/global"
	"cpm/model/cpm"
	"cpm/model/cpm/request"
	"cpm/model/cpm/response"
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
func (cpmService *CpmService) AddProject(cpmProject cpm.CpmProject) (cpmInter cpm.CpmProject, err error) {
	var project cpm.CpmProject
	if !errors.Is(global.CPM_DB.Where("name = ?", cpmProject.Name).First(&project).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		println("项目已经存在")
		return cpmInter, errors.New("项目已经存在")
	}
	cpmProject.UUID = uuid.NewV4()
	err = global.CPM_DB.Create(&cpmProject).Error
	return cpmProject, err
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

/**
 * @description: 分页获取项目详情列表
 * @param {request.SysDictionaryDetailSearch} info
 * @return {*}
 */
func (cpmService *CpmService) GetProject(info request.SysDictionaryDetailSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.CPM_DB.Model(&cpm.CpmProject{})
	var sysDictionaryDetails []cpm.CpmProject
	var project []response.ProjectList
	print(info.TypeId)
	if info.UUID != uuid.Nil {
		db = db.Where("uuid = ?", info.UUID)
	}
	if info.TypeId != "" {
		db = db.Where("type_id = ?", info.TypeId)
	}
	if info.LanguageId != "" {
		db = db.Where("status = ?", info.LanguageId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Preload("Author").Preload("Type").Preload("Language").Find(&sysDictionaryDetails).Error
	for _, v := range sysDictionaryDetails {
		project = append(project, response.ProjectList{
			UUID:         v.UUID,
			Name:         v.Name,
			AuthorName:   v.Author.NickName,
			LanguageName: v.Language.LanguageName,
			TypeName:     v.Type.TypeName,
		})
	}

	return project, total, err
}
