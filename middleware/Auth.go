/*
@Time    : 2020/8/1
@Author  : Wangcq
@File    : Auth.go
@Software: GoLand
*/

package middleware

import (
	"github.com/gin-gonic/gin"
	"goinception/logs"
	"goinception/model"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取authorization header
		tokenString := ctx.GetHeader("Authorization")
		// validate token formate
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			logs.Log.Infoln("Token格式错误, ", tokenString)
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}
		tokenString = tokenString[7:]
		token, claims, err := ParseToken(tokenString)
		if err != nil || !token.Valid {
			logs.Log.Infoln("Token解析失败 Or Token 失效, ", tokenString)
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}
		// 验证通过后获取claim 中的userId
		userId := claims.UserId
		var user model.Users
		model.DB.First(&user, userId)
		// 用户是否存在
		if user.ID == 0 {
			logs.Log.Infoln("user不存在 ", user)
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}
		// 用户存在 将user 的信息写入上下文
		ctx.Set("user", user)
		ctx.Next()
	}
}
