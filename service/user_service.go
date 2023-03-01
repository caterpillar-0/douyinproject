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
	Login(username string, password string) (*dto.UserDTO, error)
	GetUser(userid uint) (*dto.UserInfoDTO, error)

	ParamValid(username string, password string) error
	FindByName(username string) (*entity.User, error)
	FindByID(userid uint) (*entity.User, error)
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
	if _, err := c.FindByName(username); err == nil {
		return nil, errors.New("username has exits")
	}
	//2、密码加密
	hashPassword, _ := utils.PasswordHash(password)
	//3、插入数据库
	uq := dao.Q.User
	user := entity.User{
		Username: username,
		Password: hashPassword,
	}
	if err := uq.Create(&user); err != nil {
		return nil, errors.New("create user failed")
	}
	//token生成
	token, _ := utils.GenToken(user.ID)
	//4、返回结果
	return &dto.UserDTO{UserID: user.ID, Token: token}, nil
}

func (c *userService) Login(username string, password string) (*dto.UserDTO, error) {
	if err := c.ParamValid(username, password); err != nil {
		return nil, errors.New("username or password length error")
	}
	user, err := c.FindByName(username)
	if err != nil {
		return nil, errors.New("username doesn't exits")
	}
	if !utils.PasswordValid(user.Password, password) {
		return nil, errors.New("username or password error")
	}
	token, _ := utils.GenToken(user.ID)
	return &dto.UserDTO{UserID: user.ID, Token: token}, nil
}

func (c *userService) GetUser(userid uint) (*dto.UserInfoDTO, error) {
	user, err := c.FindByID(userid)
	if err != nil {
		return nil, errors.New("userid find failed")
	}
	return dto.NewUserInfoDTO(user, userid), nil
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

func (c *userService) FindByName(username string) (*entity.User, error) {
	uq := dao.Q.User
	if user, err := uq.Where(uq.Username.Eq(username)).First(); err == nil {
		return user, nil
	} else {
		return nil, err
	}
}

func (c *userService) FindByID(userid uint) (*entity.User, error) {
	uq := dao.Q.User
	if user, err := uq.Where(uq.ID.Eq(userid)).First(); err == nil {
		return user, nil
	} else {
		return nil, err
	}
}
