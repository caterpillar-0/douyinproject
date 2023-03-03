package controller

import (
	"MyProject/dto"
	"MyProject/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type videoController struct {
	videoService service.VideoService
}

func NewVideoController(videoService service.VideoService) *videoController {
	return &videoController{
		videoService: videoService,
	}
}

func (c *videoController) PublishVideo(ctx *gin.Context) {
	userID := GETID(ctx)
	title := ctx.PostForm("title")
	form, err := ctx.MultipartForm()
	if err != nil {
		ErrorResponse(ctx, err.Error())
		return
	}
	videofile := form.File["data"][0]
	err = c.videoService.PublishAction(userID, title, videofile)
	if err != nil {
		ErrorResponse(ctx, err.Error())
	} else {
		ctx.JSON(http.StatusOK, dto.Response{
			StatusCode: 0,
			StatusMsg:  "success publish video!",
		})
	}

}
