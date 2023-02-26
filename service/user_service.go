package service

import (
	"MyProject/dao"
	"MyProject/dto"
	"MyProject/entity"
	"errors"

	"MyProject/utils"
)

const (
	MaxUsernameLength = 32 //Maximum length of user name
	MaxPasswordLength = 32 //Maximum password length
	MinPasswordLength = 6  //Minimum password length
)

//UserService接口用来定义行为类型，接口中声明完方法后，
//userService结构体重写UserService接口中所有方法，认为userService结构体实现了接口UserService

type UserService interface {
	Register(username string, password string) (*dto.UserDTO, error)
	ParamValid(username string, password string) error
}

// 接口的实例
type userService struct {
}

func NewUserService() UserService {
	return &userService{}
}

func (c *userService) Register(username string, password string) (*dto.UserDTO, error) {
	//0.用户名和密码长度验证
	if err := c.ParamValid(username, password); err != nil {
		return nil, errors.New("username or password length error")
	}

	//1、查询数据库是否存在
	uq := dao.Q.User
	_, err := uq.Where(uq.Username.Eq(username)).First()
	if err == nil {
		return nil, errors.New("username already exists")
	}

	//2、密码加密
	hashPassword, _ := utils.PasswordHash(password)

	//3、插入数据库
	user := entity.User{
		Username: username,
		Password: hashPassword,
	}

	err = uq.Create(&user)
	if err != nil {
		return nil, errors.New("create user failed")
	}
	//token生成
	token := username + "&" + password

	//4、返回结果
	return &dto.UserDTO{UserID: user.ID, Token: token}, nil
}

func (c *userService) ParamValid(username string, password string) error {
	//1.username
	if username == "" {
		return errors.New("username is null")
	}
	if len(username) > MaxUsernameLength {
		return errors.New("username is too long")
	}
	//2.password
	if password == "" {
		return errors.New("password is null")
	}
	if len(password) > MaxPasswordLength || len(password) < MinPasswordLength {
		return errors.New("password length error")
	}
	return nil
}
