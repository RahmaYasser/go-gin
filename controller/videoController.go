package controller

import (
	"github.com/RahmaYasser/go-gin/entity"
	"github.com/RahmaYasser/go-gin/service"
	"github.com/RahmaYasser/go-gin/validators"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type VideoController interface {
	Save(ctx *gin.Context) error
	FindAll() []entity.Video
	ShowAll(ctx *gin.Context)
}

type videoController struct {
	service service.VideoService
}

var validate *validator.Validate

func New(videoService service.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	return &videoController{service: videoService}
}

func (controller *videoController) Save(ctx *gin.Context) error {
	//controller.service.Save(ctx)
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}
	err = validate.Struct(video)
	if err != nil {
		return err
	}
	controller.service.Save(video)
	return nil
}
func (controller *videoController) FindAll() []entity.Video {
	return controller.service.FindAll()
}
func (controller *videoController) ShowAll(ctx *gin.Context) {
	videos := controller.service.FindAll()
	data := gin.H{
		"title":  "video page",
		"videos": videos,
	}
	ctx.HTML(200, "index.html", data)
}
