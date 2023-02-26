package entity

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username       string `gorm:"size:50"`
	Password       string `gorm:"size:200"`
	FollowCount    uint   `gorm:"default:0"`
	FollowerCount  uint   `gorm:"default:0"`
	TotalFavorited uint   `gorm:"default:0"`
	WorkCount      uint   `gorm:"default:0"`
	FavoriteCount  uint   `gorm:"default:0"`
}
