package cpm

import (
	"cpm/global"
	"cpm/model/cpm"
	"cpm/model/cpm/response"

	"gorm.io/gorm"
)

type CpmImportService struct{}

func (cpmService *CpmImportService) GetImport(cpmImport cpm.CpmImport) (cpmInter []cpm.CpmImport, err error) {
	var importList []cpm.CpmImport
	err = global.CPM_DB.Where("project_id = ? and version = ?", cpmImport.ProjectId, cpmImport.Version).Find(&importList).Error
	return importList, err
}

func (cpmService *CpmImportService) GetCpmImportSearchList(cpmImport []cpm.CpmImport) (cpmInter []response.AutoListUrl, err error) {
	err = global.CPM_DB.Transaction(func(tx *gorm.DB) error {
		for _, v := range cpmImport {
			var api cpm.CpmImport
			if err2 := tx.Where("version = ? AND name = ? AND project_id = ?", v.Version, v.Name, v.ProjectId).First(&api).Error; err != nil {
				return err2
			}
			if api.Name != "" {
				cpmInter = append(cpmInter, response.AutoListUrl{
					Name: api.Name,
					List: api.List,
				})
			}
		}
		return nil
	})
	return cpmInter, err
}
