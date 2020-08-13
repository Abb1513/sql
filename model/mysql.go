/*
@Time    : 2020/5/28
@Author  : Wangcq
@File    : mysql.go
@Software: GoLand
*/

package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	log "goinception/logs"
)

// var DB *xorm.Engine
var DB *gorm.DB

func InitMysql() {
	//driverName := viper.GetString("datasource.driverName")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Local",
		username,
		password,
		host,
		port,
		database,
		charset)

	db, err := gorm.Open("mysql", args)
	db.LogMode(true)
	if err != nil {
		log.Log.Errorln("fail to connect databse,err:" + err.Error())
	}
	db.SingularTable(true)
	db.AutoMigrate(&DbInfo{})
	db.AutoMigrate(&ResultMessage{})
	db.AutoMigrate(&Users{})
	DB = db
}

//初始化 root 用户
