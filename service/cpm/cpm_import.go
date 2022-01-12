package cpm

import (
	"cpm/global"
	"cpm/model/cpm"
)

type CpmImportService struct{}

func (cpmService *CpmImportService) GetImport(cpmImport cpm.CpmImport) (cpmInter []cpm.CpmImport, err error) {
	println(cpmImport.ProjectId, cpmImport.Version, "@222")
	var importList []cpm.CpmImport
	err = global.CPM_DB.Where("project_id = ? and version = ?", cpmImport.ProjectId, cpmImport.Version).Find(&importList).Error
	return importList, err
}
