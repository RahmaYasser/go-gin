package main

import (
	"github.com/RahmaYasser/go-gin/controller"
	"github.com/RahmaYasser/go-gin/service"
	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.Default()
	server.GET("/posts", func(context *gin.Context) {
		context.JSON(200, videoController.FindAll())
	})
	server.Run("localhost:8080")
}
