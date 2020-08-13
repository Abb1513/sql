/*
@Time    : 2020/5/28
@Author  : Wangcq
@File    : dbuser.go
@Software: GoLand
*/
// 数据用户创建与修改
package controller

import (
	log "goinception/logs"
	"goinception/model"
	"goinception/response"
	"goinception/utools"

	"github.com/gin-gonic/gin"
)

// @Summary 创建db user
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DbInfo
// @Success 200 {string} string "{"httpStatus":200,"data":{},"msg":"succeed"}"
// @Router /api/dbuser/ [post]
func CreateDbuser(ctx *gin.Context) {
	var d model.DbInfo
	err := ctx.ShouldBindJSON(&d)
	log.Log.Infoln("CreateDbuser入参, ", d)
	if err != nil {
		response.Response(ctx, 200, nil, err)
		return
	}
	d.DbPassword = utools.EncryptDbPasswd(d.DbPassword)
	model.DB.Table("db_info").Create(&d)
	response.Response(ctx, 200, d, "succeed")
}

// @Summary 获取所有 db user
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"httpStatus":200,"data":{},"msg":"succeed"}"
// @Router /api/dbuser/ [get]
func GetAllDbUser(ctx *gin.Context) {
	var resList []model.DbinfoResponse
	model.DB.Table("db_info").Find(&resList)
	// 返回数据中没有 password字段
	response.Response(ctx, 200, resList, "succeed")
}

// 屏蔽
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

// @Summary 更新db user
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DbInfo
// @Success 200 {string} string "{"httpStatus":200,"data":{},"msg":"succeed"}"
// @Router /api/updatedbuser/ [post]
func UpdateDbUser(ctx *gin.Context) {
	var d model.DbInfo
	err := ctx.ShouldBindJSON(&d)
	log.Log.Infoln("UpdateDbUser入参, ", d)
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
