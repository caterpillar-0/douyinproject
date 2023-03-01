package router

import (
	"MyProject/controller"
	"MyProject/middleware"
	"MyProject/service"

	"github.com/gin-gonic/gin"
)

var (
	userService = service.NewUserService()

	userController = controller.NewUserController(userService)
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	//change to `gin.ReleaseMode` in production.
	gin.SetMode(gin.DebugMode)

	apiRouter := r.Group("/douyin/")

	//basic apis
	apiRouter.POST("user/register/", userController.Register)
	apiRouter.POST("user/login/", userController.Login)
	apiRouter.GET("user/", middleware.JWTMiddleware(), userController.GetUserInfo)

	return r

}
