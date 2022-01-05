package core

import (
	"cpm/global"
	"cpm/utils"
	"flag"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func Viper(path ...string) *viper.Viper {
	var config string
	if len(path) == 0 {
		// 打包完成后的命令行参数
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.Parse()
		if config == "" {

			if configEnv := os.Getenv(utils.ConfigEnv); configEnv == "" {
				config = utils.ConfigFile
				fmt.Printf("您正在使用config的默认值,config的路径为%v\n", utils.ConfigFile)
			}
		}
	}
	v := viper.New()
	v.SetConfigFile(config) //配置文件名
	v.SetConfigType("yaml") //# 配置文件类型
	err := v.ReadInConfig() // 读取配置文件信息
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	// 填充结构体
	if err := v.Unmarshal(&global.CPM_CONFIG); err != nil {
		fmt.Println(err)
	}
	println(global.CPM_CONFIG.System.DbType)
	return v
}
