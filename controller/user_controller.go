package controller

import (
	"MyProject/dto"
	"MyProject/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {

	//结构体内嵌，可构建一种面向对象中的继承关系
	//结构体实例化后，可以直接访问内嵌结构体的所有成员变量和方法

	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return UserController{
		userService: userService,
	}
}

// Register handels `/user/register/`
func (c *UserController) Register(ctx *gin.Context) {
	username := ctx.Query("username")
	password := ctx.Query("password")

	userDTO, err := c.userService.Register(username, password)
	if err != nil {
		ErrorResponse(ctx, err.Error())
	} else {
		ctx.JSON(http.StatusOK, dto.UserResponse{
			Response: dto.Response{StatusCode: 0, StatusMsg: "test"},
			UserDTO:  *userDTO,
			//!!!这里报错，必须是if else,不然这里就会对空指针解引用！！！
		})
	}
}
