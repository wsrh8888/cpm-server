package initialize

import (
	"cpm/global"

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
