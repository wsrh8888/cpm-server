package initialize

import (
	"cpm/global"
	"cpm/model/cpm"
	"cpm/model/system"
	"os"

	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	switch global.CPM_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	default:
		return GormMysql()
	}
}

func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		cpm.CpmLanguage{},
		cpm.CpmType{},
		cpm.CpmProject{},
		cpm.CpmImport{},
		cpm.CpmVersion{},

		system.SysUser{},
	)
	if err != nil {
		// global.GVA_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
}
