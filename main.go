package main

import (
	"cpm/core"
	"cpm/global"
	"cpm/initialize"
)

func main() {
	// 初始化部分global.CPM_CONFIG 相关数据
	global.CPM_VP = core.Viper()      // 初始化Viper
	global.CPM_DB = initialize.Gorm() // gorm连接数据库
	core.RunWindowsServer()
}

// func main() {
// InitConfig()
// db := initialize.InitDB()

// defer db.Close()

// router := gin.Default()
// router.Use(initialize.CrosMiddleWare)

// all_router.InitRouter(router)
// port := viper.GetString("server.port")

// router.Run(":" + port)
// }

// func InitConfig() {
// 	workDir, _ := os.Getwd()
// 	viper.SetConfigName("application")
// 	viper.SetConfigType("yml")
// 	viper.AddConfigPath(workDir)
// 	err := viper.ReadInConfig()
// 	if err != nil {
// 		panic(err)
// 	}
// }
