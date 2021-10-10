package main

import (
	"github.com/spf13/viper"
	"os"
	"yeb/common"
	"yeb/router"
)

func main() {
	InitConfig()
	common.InitDB()

	router.InitRouter()
}

// InitConfig 使用viper读取yml配置文件
func InitConfig()  {
	workDir,_ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}