/*
@Time    : 2020/5/28
@Author  : Wangcq
@File    : route.go
@Software: GoLand
*/

package router

import (
	"github.com/gin-gonic/gin"
	"goinception/controller"
	"goinception/middleware"
)

func Router(r *gin.Engine) {
	r.Use(middleware.CORSMiddleware())

	// SQL 操作
	db := r.Group("/api/db")
	db.Use(middleware.AuthMiddleware())
	db.POST("/dbuser/", controller.CreateDbuser)
	db.GET("/dbuser/", controller.GetAllDbUser)
	//	r.POST("/api/deletedbuser/", controller.DeleteDbUser)
	db.POST("/updatedbuser/", controller.UpdateDbUser)
	db.POST("/execute/", controller.Execute)
	db.POST("/executerollsql/", controller.ExecuteRollSql)
	db.GET("/execute/", controller.ExecuteGetAll)

	// 用户相关
	users := r.Group("/api/user")
	users.POST("/login/", controller.UserLogin)
	users.POST("/register/", controller.UserRegister)
}
