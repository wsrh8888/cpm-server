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

func (cpmService *CpmService) GetImportInfo(info request.CpmProjectAllInfo) (cpmAll response.CpmProjectAllInfo, err error) {
	// 查询project相关的信息
	db := global.CPM_DB.Table("cpm_projects")
	db = db.Select("cpm_projects.uuid as id, cpm_projects.name, sys_users.nick_name as author_name, cpm_languages.name as language_name, cpm_types.name as type_name")
	db = db.Where("cpm_projects.name = ?", info.Name)
	db = db.Joins("left join sys_users on cpm_projects.author_id = sys_users.id")
	db = db.Joins("left join cpm_languages on cpm_projects.language_id = cpm_languages.id")
	db = db.Joins("left join cpm_types on cpm_projects.type_id = cpm_types.id")
	db.Scan(&cpmAll.CpmProject)

	// 查询version相关的信息
	dbVersion := global.CPM_DB.Table("cpm_versions")
	dbVersion = dbVersion.Select("sys_users.nick_name as publisher_name, cpm_versions.description,  cpm_versions.version,  cpm_versions.created_at, cpm_versions.keywords")
	dbVersion = dbVersion.Where("project_id = ?", cpmAll.CpmProject.Id)
	dbVersion = dbVersion.Order("created_at DESC")
	dbVersion = dbVersion.Joins("left join sys_users on cpm_versions.publish_id = sys_users.id")
	dbVersion.Scan(&cpmAll.CpmVersionNew)

	// 查询所有的version版本信息
	dbVersions := global.CPM_DB.Table("cpm_versions")
	dbVersions = dbVersions.Select("sys_users.nick_name as publisher_name, cpm_versions.description,  cpm_versions.version,  cpm_versions.created_at, cpm_versions.keywords")
	dbVersions = dbVersions.Where("project_id = ?", cpmAll.CpmProject.Id)
	dbVersions = dbVersions.Order("version DESC")
	dbVersions = dbVersions.Joins("left join sys_users on cpm_versions.publish_id = sys_users.id")
	dbVersions.Find(&cpmAll.CpmVersions)

	// 查询import
	global.CPM_DB.Where("project_id = ? and version = ?", cpmAll.CpmProject.Id, cpmAll.CpmVersionNew.Version).Find(&cpmAll.CpmImport)

	return cpmAll, err
}

/**
 * @description: 分页获取项目详情列表
 * @param {request.SysDictionaryDetailSearch} info
 * @return {*}
 */
func (cpmService *CpmService) GetProject(info request.SysDictionaryDetailSearch) (cpmProject []response.ProjectList, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.CPM_DB.Model(&cpm.CpmProject{})
	db = db.Select("cpm_projects.uuid as id, cpm_projects.name, sys_users.nick_name as author_name, cpm_languages.name as language_name, cpm_types.name as type_name")
	db = db.Joins("left join sys_users on cpm_projects.author_id = sys_users.id")
	db = db.Joins("left join cpm_languages on cpm_projects.language_id = cpm_languages.id")
	db = db.Joins("left join cpm_types on cpm_projects.type_id = cpm_types.id")
	if info.Name != "" {
		db = db.Where("name = ?", info.Name)
	}
	if info.UUID != uuid.Nil {
		db = db.Where("uuid = ?", info.UUID)
	}
	if info.TypeId != "" {
		db = db.Where("type_id = ?", info.TypeId)
	}
	if info.LanguageId != "" {
		db = db.Where("language_id = ?", info.LanguageId)
	}
	db = db.Scan(&cpmProject)

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&cpmProject).Error
	return cpmProject, total, err
}
