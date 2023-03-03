package controller

import (
	"MyProject/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorResponse(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, dto.Response{
		StatusCode: 1,
		StatusMsg:  msg,
	})
}

// 获取token中ID
func GETID(ctx *gin.Context) uint {
	tokenRawId, _ := ctx.Get("token_id") //interface{}uint类型
	tokenID := tokenRawId.(uint)
	return tokenID
}
