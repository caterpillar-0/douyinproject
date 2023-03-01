package dto

import "MyProject/entity"

type UserDTO struct {
	UserID uint   `json:"user_id,omitempty"`
	Token  string `json:"token,omitempty"`
}
type UserResponse struct {
	Response
	UserDTO
}

type UserInfoDTO struct {
	ID             uint   `json:"id"`             // 用户id
	Name           string `json:"name"`           // 用户名称
	FollowCount    uint   `json:"follow_count"`   // 关注总数
	FollowerCount  uint   `json:"follower_count"` // 粉丝总数
	IsFollow       bool   `json:"is_follow"`
	TotalFavorited uint   `json:"total_favorited"` // 获赞数量
	WorkCount      uint   `json:"work_count"`      // 作品数
	FavoriteCount  uint   `json:"favorite_count"`  // 喜欢数
}

type UserInfoResponse struct {
	Response
	UserInfoDTO `json:"user"`
}

func NewUserInfoDTO(user *entity.User, loginID uint) *UserInfoDTO {
	//根据logintID确定is_follow!!!
	//TODO:is follow
	isFollow := false
	if user.ID == loginID {
		isFollow = true
	}
	return &UserInfoDTO{
		ID:             user.ID,
		Name:           user.Username,
		FollowCount:    user.FollowCount,
		IsFollow:       isFollow,
		FollowerCount:  user.FollowerCount,
		TotalFavorited: user.TotalFavorited,
		WorkCount:      user.WorkCount,
		FavoriteCount:  user.FavoriteCount,
	}
}
