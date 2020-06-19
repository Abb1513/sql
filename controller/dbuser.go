/*
@Time    : 2020/5/28
@Author  : Wangcq
@File    : dbuser.go
@Software: GoLand
*/

package controller

import (
	"github.com/gin-gonic/gin"
	"goinception/model"
	"goinception/response"
	"goinception/utools"
)

func CreateDbuser(ctx *gin.Context) {
	var d model.DbInfo
	err := ctx.ShouldBindJSON(&d)
	if err != nil {
		response.Response(ctx, 200, nil, err)
		return
	}
	d.DbPassword = utools.EncryptDbPasswd(d.DbPassword)
	model.DB.Table("db_info").Create(&d)
	response.Response(ctx, 200, d, "succeed")
}

func GetAllDbUser(ctx *gin.Context) {
	var resList []model.DbinfoResponse
	model.DB.Table("db_info").Find(&resList)
	// 返回数据中没有 password字段
	response.Response(ctx, 200, resList, "succeed")
}

func DeleteDbUser(ctx *gin.Context) {
	var d model.DbInfo
	err := ctx.ShouldBindJSON(&d)
	if err != nil {
		response.Response(ctx, 200, nil, err)
		return
	}
	model.DB.Delete(&d)
	response.Response(ctx, 200, nil, "succeed")
}

func UpdateDbUser(ctx *gin.Context) {
	var d model.DbInfo
	err := ctx.ShouldBindJSON(&d)
	if err != nil {
		response.Response(ctx, 200, nil, err)
		return
	}

	// post 没有password 就将库中取出来
	// 不更新密码
	if d.DbPassword == "" {
		var tmp model.DbInfo
		model.DB.First(&tmp, d.ID)
		d.DbPassword = tmp.DbPassword
		model.DB.Model(&model.DbInfo{}).Updates(&d)
		response.Response(ctx, 200, d, "succeed")
	} else {
		d.DbPassword = utools.EncryptDbPasswd(d.DbPassword)
		model.DB.Model(&model.DbInfo{}).Updates(&d)
		response.Response(ctx, 200, d, "succeed")
	}
}
