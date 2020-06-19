/*
@Time    : 2020/5/28
@Author  : Wangcq
@File    : main.go
@Software: GoLand
*/

package main

import (
	"goinception/model"
	"goinception/router"
	"goinception/utools"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func initConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	initConfig()
	utools.InitGoinception()
	model.InitMysql()
	defer model.DB.Close()
	r := gin.Default()
	router.Router(r)
	if port := viper.GetString("app.port"); port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())
}
