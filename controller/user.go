/*
@Time    : 2020/8/1
@Author  : Wangcq
@File    : user.go
@Software: GoLand
*/

package controller

import (
	"github.com/gin-gonic/gin"
	"goinception/logs"
	"goinception/middleware"
	"goinception/model"
	"goinception/response"
	"goinception/utools"
)

// @Summary 用户登录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body execute
// @Success 200 {string} string "{"httpStatus":200,"data":{},"msg":"succeed"}"
// @Router /api/user/login [post]
func UserLogin(ctx *gin.Context) {
	var tmp model.Users
	err := ctx.ShouldBindJSON(&tmp)
	if err != nil {
		logs.Log.Error("Json 解析失败, ", err)
	}
	password := tmp.Password
	model.DB.Where("name = ?", tmp.Name).First(&tmp)
	if res, err := utools.DecodeUserPassword(tmp.Password, password); res && err == nil {
		token, err := middleware.ReleaseToken(tmp)
		if err == nil {
			logs.Log.Infof("用户%s 登录成功，发放token: %s ", tmp.Name, token)
			response.Response(ctx, 200, gin.H{"token": token}, "登录成功")
			return
		} else {
			logs.Log.Errorf("Token发放失败, ", err)
			response.Response(ctx, 500, gin.H{"error": err}, "用户登录失败")
			return
		}

	} else {
		logs.Log.Errorf("密码解密验证失败, ", tmp)
		response.Response(ctx, 500, gin.H{"error": err}, "用户登录失败")
		return
	}

}

// @Summary 用户注册
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body execute
// @Success 200 {string} string "{"httpStatus":200,"data":{},"msg":"succeed"}"
// @Router /api/user/register [post]
func UserRegister(ctx *gin.Context) {
	var tmp model.Users
	err := ctx.ShouldBindJSON(&tmp)
	if err != nil {
		logs.Log.Error("Json 解析失败, ", err)
	}
	model.DB.Create(&tmp)
	response.Response(ctx, 200, gin.H{"data": tmp}, "succeed")
}
