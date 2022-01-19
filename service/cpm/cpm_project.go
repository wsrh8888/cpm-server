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

type result struct {
	Name         string `json:"name"`
	Author       string `json:"author"`
	ServiceId    string `json:"serviceId"`
	ServiceName  string `json:"serviceName"`
	LanguageName string `json:"languageName"`
	TypeName     string `json:"typeName"`
}

func (cpmService *CpmService) GetImportInfo(info request.CpmProjectAllInfo) (result response.CpmProjectAllInfo, err error) {
	println("*********")

	println(info.Name)
	var cpmAll response.CpmProjectAllInfo
	db := global.CPM_DB.Table("cpm_projects")
	db = db.Select("cpm_projects.uuid as id, cpm_projects.name, sys_users.nick_name as author_name, cpm_languages.name as language_name, cpm_types.name as type_name")
	db = db.Where("cpm_projects.name = ?", info.Name)
	db = db.Joins("left join sys_users on cpm_projects.author_id = sys_users.id")
	db = db.Joins("left join cpm_languages on cpm_projects.language_id = cpm_languages.id")
	db = db.Joins("left join cpm_types on cpm_projects.type_id = cpm_types.id")
	db = db.Scan(&cpmAll.CpmProject)

	global.CPM_DB.Model(&cpm.CpmVersion{}).Where("project_id = ?", cpmAll.CpmProject.Id)
	// global.CPM_DB.Table("cpm_projects")
	// global.CPM_DB.Table("cpm_projects").Select("cpm_projects.name as name, sys_users.author_name as author_name").Joins("left join sys_users on cpm_projects.author_id = sys_users.id").Scan(&result)
	// global.CPM_DB.Where("project_id = ? and version = ?", info.Name, info.Version).Preload("Publish").First(&version)
	// global.CPM_DB.Where("project_id = ?", info.Name).Preload("Author").Preload("Type").Preload("Language").First(&project)
	// if info.Version != "" {

	// }

	return cpmAll, err
}

/**
 * @description: 分页获取项目详情列表
 * @param {request.SysDictionaryDetailSearch} info
 * @return {*}
 */
func (cpmService *CpmService) GetProject(info request.SysDictionaryDetailSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	global.CPM_DB.Table("cpm_projects").Joins("left join sys_users on cpm_projects.author_id = sys_users.Id")
	db := global.CPM_DB.Model(&cpm.CpmProject{})
	var sysDictionaryDetails []cpm.CpmProject
	var project []response.ProjectList
	// db.Table("go_service_info").Select("go_service_info.serviceId as service_id, go_service_info.serviceName as service_name, go_system_info.systemId as system_id, go_system_info.systemName as system_name").Joins("left join go_system_info on go_service_info.systemId = go_system_info.systemId").Scan(&results)
	// print(info.TypeId, "2222")
	// if info.Name != "" {
	// 	db = db.Where("name = ?", info.Name)
	// }
	// if info.UUID != uuid.Nil {
	// 	db = db.Where("uuid = ?", info.UUID)
	// }
	// if info.TypeId != "" {
	// 	db = db.Where("type_id = ?", info.TypeId)
	// }
	// if info.LanguageId != "" {
	// 	db = db.Where("status = ?", info.LanguageId)
	// }
	// err = db.Count(&total).Error
	// if err != nil {
	// 	return
	// }
	err = db.Limit(limit).Offset(offset).Preload("Author").Preload("Type").Preload("Language").Find(&sysDictionaryDetails).Error
	// for _, v := range sysDictionaryDetails {
	// 	println("测试测试")
	// project = append(project, response.ProjectList{
	// 	Id:           v.UUID,
	// 	Name:         v.Name,
	// 	AuthorName:   v.Author.NickName,
	// 	LanguageName: v.Language.LanguageName,
	// 	TypeName:     v.Type.TypeName,
	// })
	// }

	return project, total, err
}
