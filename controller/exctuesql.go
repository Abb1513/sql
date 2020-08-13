/*
@Time    : 2020/5/28
@Author  : Wangcq
@File    : exctuesql.go
@Software: GoLand
*/

package controller

import (
	"github.com/gin-gonic/gin"
	log "goinception/logs"
	"goinception/model"
	"goinception/response"
	"goinception/utools"
)

type execute struct {
	SQL string `json:"sql"`
	ID  int    `json:"id"`
}

// @Summary 执行sql
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body execute
// @Success 200 {string} string "{"httpStatus":200,"data":{},"msg":"succeed"}"
// @Router /api/execute/ [post]
func Execute(ctx *gin.Context) {
	var d execute
	var reslut []model.ResultMessage
	err := ctx.ShouldBindJSON(&d)
	log.Log.Infoln("Execute入参, ", d)
	if err != nil {
		response.Response(ctx, 200, nil, err)
		return
	}
	var tmp model.DbInfo
	model.DB.First(&tmp, d.ID)
	res := utools.Exec(tmp, d.SQL)
	for _, i := range res {
		if i.ErrorMessage != "" {
			response.Response(ctx, 200, gin.H{"sql": i.Sql, "error": i.ErrorMessage}, "execute fail")
			return
		}
	}
	for _, i := range res {
		i.Rollsql = utools.GetRollSql(i.BackupDbname, i.OpidTime)
		model.DB.Save(&i)
		reslut = append(reslut, i)
	}
	response.Response(ctx, 200, reslut, "succeed")
}

// @Summary 执行回滚sql语句
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ResultMessage
// @Success 200 {string} string "{"httpStatus":200,"data":{},"msg":"succeed"}"
// @Router /api/executerollsql/ [post]
func ExecuteRollSql(ctx *gin.Context) {
	var res model.ResultMessage
	var dbinfo model.DbInfo
	err := ctx.ShouldBindJSON(&res)
	log.Log.Infoln("ExecuteRollSql入参, ", res)
	if err != nil {
		response.Response(ctx, 200, "Json解析失败", err)
		return
	}
	model.DB.First(&dbinfo, res.ExcuteDb)
	model.DB.First(&res, res.ID)
	if !res.IsExcuteRollsql {
		if err = utools.ExcuteRollBak(res, dbinfo); err != nil {
			response.Response(ctx, 200, "执行回滚Sql失败", err)
			return
		}
	}
	res.IsExcuteRollsql = true
	model.DB.Model(&model.ResultMessage{}).Updates(&res)
	response.Response(ctx, 200, "执行回滚Sql成功", "succeed")
}

// @Summary 获取所有执行历史
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"httpStatus":200,"data":{},"msg":"succeed"}"
// @Router /api/execute/ [get]
func ExecuteGetAll(ctx *gin.Context) {
	var resList []model.ResultMessage
	model.DB.Find(&resList)
	response.Response(ctx, 200, resList, "succeed")
}
