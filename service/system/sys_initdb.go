package system

import (
	"cpm/global"
	"cpm/model/cpm"
	"cpm/model/system"
	model "cpm/model/system"
)

type InitDBService struct{}

func (initDBService *InitDBService) initTables() error {
	return global.CPM_DB.AutoMigrate(
		system.SysUser{},
		cpm.CpmProject{},
		cpm.CpmImport{},
		cpm.CpmImport{},
	)
}

func (initDBService *InitDBService) InitDB() error {
	if err := initDBService.initTables(); err != nil {
		global.CPM_DB = nil
		return err
	}

	if err := initDBService.initMysqlData(); err != nil {
		global.CPM_DB = nil
		return err
	}
	return nil
}

func (initDBService *InitDBService) initMysqlData() error {
	return model.MysqlDataInitialize(
	// system.User,
	)
}
