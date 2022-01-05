package main

import (
	"cpm/initialize"
	all_router "cpm/routers"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	InitConfig()
	db := initialize.InitDB()

	defer db.Close()

	router := gin.Default()
	router.Use(initialize.CrosMiddleWare)

	all_router.InitRouter(router)
	port := viper.GetString("server.port")

	router.Run(":" + port)
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
