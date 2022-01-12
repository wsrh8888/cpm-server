package cpm

import (
	"cpm/global"
	"cpm/model/cpm"
)

type CpmImportService struct{}

func (cpmService *CpmImportService) GetImport(cpmImport cpm.CpmImport) (cpmInter []cpm.CpmImport, err error) {
	var importList []cpm.CpmImport
	err = global.CPM_DB.Where("project_id = ? and version = ?", cpmImport.ProjectId, cpmImport.Version).Find(&importList).Error
	return importList, err
}
