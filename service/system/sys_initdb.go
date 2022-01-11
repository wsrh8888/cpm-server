package system

import (
	"cpm/global"
	"cpm/model/cpm"
	model "cpm/model/system"
	cpmSource "cpm/source/cpm"
	cpmSystem "cpm/source/system"
)

type InitDBService struct{}

func (initDBService *InitDBService) initTables() error {
	return global.CPM_DB.AutoMigrate(
		model.SysUser{},
		cpm.CpmProject{},
		cpm.CpmImport{},
		cpm.CpmImport{},
	)
}

func (initDBService *InitDBService) InitDB() error {

	if err := initDBService.initMysqlData(); err != nil {
		global.CPM_DB = nil
		return err
	}
	return nil
}

func (initDBService *InitDBService) initMysqlData() error {
	return model.MysqlDataInitialize(
		cpmSystem.User,
		cpmSource.Language,
		cpmSource.Type,
	)
}
