package middleware

import (
	"MyProject/controller"
	"MyProject/utils"

	"github.com/gin-gonic/gin"
)

// verify token
func JWTMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// get token string
		tokenString := ctx.Query("token")
		claims, err := utils.ValidToken(tokenString)
		if err != nil {
			controller.ErrorResponse(ctx, err.Error())
			ctx.Abort()
			return
		}
		ctx.Set("token_id", claims.UserID)
		ctx.Next()
	}
}
