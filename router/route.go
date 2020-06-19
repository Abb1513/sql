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
)

func Router(r *gin.Engine) {
	r.POST("/api/dbuser/", controller.CreateDbuser)
	r.GET("/api/dbuser/", controller.GetAllDbUser)
	r.POST("/api/deletedbuser/", controller.DeleteDbUser)
	r.POST("/api/updatedbuser/", controller.UpdateDbUser)
	r.POST("/api/execute/", controller.Execute)
	r.POST("/api/executerollsql/", controller.ExecuteRollSql)
	r.GET("/api/execute/", controller.ExecuteGetAll)
}
