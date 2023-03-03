package entity

import (
	"github.com/jinzhu/gorm"
)

type Video struct {
	gorm.Model
	UserID        uint
	User          User   `gorm:"foreignkey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	PlayURL       string `gorm:"size:50"`
	CoverURL      string `gorm:"size:50"`
	FavoriteCount uint   `gorm:"default:0"`
	CommentCount  uint   `gorm:"default:0"`
	Title         string `gorm:"size:100"`
}
