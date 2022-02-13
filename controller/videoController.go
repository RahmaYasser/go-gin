package controller

import (
	"github.com/RahmaYasser/go-gin/entity"
	"github.com/RahmaYasser/go-gin/service"
	"github.com/gin-gonic/gin"
)

type VideoController interface {
	Save(ctx *gin.Context) error
	FindAll() []entity.Video
}

type videoController struct {
	service service.VideoService
}

func New(videoService service.VideoService) VideoController {
	return &videoController{service: videoService}
}

func (controller *videoController) Save(ctx *gin.Context) error {
	//controller.service.Save(ctx)
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}
	controller.service.Save(video)
	return nil
}
func (controller *videoController) FindAll() []entity.Video {
	return controller.service.FindAll()
}
