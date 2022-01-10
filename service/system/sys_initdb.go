package system

import (
	"cpm/global"
	"cpm/model/cpm"
	model "cpm/model/system"
	source "cpm/source/system"
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
		source.User,
	)
}
