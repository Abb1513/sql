/*
@Time    : 2020/5/28
@Author  : Wangcq
@File    : response.go
@Software: GoLand
*/

package response

import "github.com/gin-gonic/gin"

func Response(ctx *gin.Context, httpStatus int, data interface{}, msg interface{}) {
	ctx.JSON(httpStatus, gin.H{"data": data, "msg": msg})
}
