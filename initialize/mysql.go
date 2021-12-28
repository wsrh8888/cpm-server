package initialize

import (
	"fmt"

	"cpm/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // indirect
	"github.com/spf13/viper"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	// driverName := viper.GetString("datasource.driverName")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local",
		username,
		password,
		host,
		port,
		database,
	)

	db, err := gorm.Open("mysql", args)
	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	}
	db.AutoMigrate(&model.User{}, &model.Task{})

	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
