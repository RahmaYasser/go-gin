package main

import (
	"github.com/RahmaYasser/go-gin/controller"
	"github.com/RahmaYasser/go-gin/middleware"
	"github.com/RahmaYasser/go-gin/service"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setLoggingOutput() {
	file, _ := os.Create("log")
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
}
func main() {
	setLoggingOutput()
	server := gin.New()
	server.Use(gin.Recovery(), middleware.Logger())
	server.GET("/videos", func(context *gin.Context) {
		context.JSON(200, videoController.FindAll())
	})
	server.POST("/videos", func(context *gin.Context) {
		context.JSON(200, videoController.Save(context))
	})
	server.Run("localhost:8080")
}
