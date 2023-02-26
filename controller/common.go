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
