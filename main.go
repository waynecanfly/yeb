package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
	"yeb/common"
)

func main() {
	InitConfig()
	common.InitDB()
	r := gin.Default()
	r = CollectRouter(r)
	port := viper.GetString("server.port")
	if port != ""{
		panic(r.Run(":" + port))
	}
	_ = r.Run()
}

// InitConfig 使用viper读取yml配置文件
func InitConfig()  {
	workDir,_ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}