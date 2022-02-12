package controller

import (
	"github.com/RahmaYasser/go-gin/entity"
	"github.com/RahmaYasser/go-gin/service"
	"github.com/gin-gonic/gin"
)

type VideoController interface {
	Save(ctx *gin.Context) entity.Video
	FindAll() []entity.Video
}

type videoController struct {
	service service.VideoService
}

func New(videoService service.VideoService) VideoController {
	return &videoController{service: videoService}
}

func (controller *videoController) Save(ctx *gin.Context) entity.Video {
	//controller.service.Save(ctx)
	var video entity.Video
	ctx.BindJSON(&video)
	controller.service.Save(video)
	return video
}
func (controller *videoController) FindAll() []entity.Video {
	return controller.service.FindAll()
}
