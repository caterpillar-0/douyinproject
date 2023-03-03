package service

import (
	"MyProject/dao"
	"MyProject/entity"
	"mime/multipart"
)

type VideoService interface {
	PublishAction(userID uint, title string, videoFile *multipart.FileHeader) error
}

type videoservice struct {
}

func NewVideoService() *videoservice {
	return &videoservice{}
}

func (c *videoservice) PublishAction(userID uint, title string, videoFile *multipart.FileHeader) error {
	video := entity.Video{
		UserID:   userID,
		PlayURL:  "",
		CoverURL: "",
		Title:    title,
	}
	vq := dao.Q.Video
	if err := vq.Create(&video); err != nil {
		return err
	}

	uq := dao.Q.User
	if _, err := uq.Where(uq.ID.Eq(userID)).UpdateSimple(uq.WorkCount.Add(1)); err != nil {
		return err
	}

	return nil
}
