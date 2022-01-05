package global

import (
	"cpm/config"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	CPM_DB     *gorm.DB
	CPM_CONFIG config.Server
	CPM_VP     *viper.Viper
)
