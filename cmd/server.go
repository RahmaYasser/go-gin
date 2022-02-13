package main

import (
	"github.com/RahmaYasser/go-gin/controller"
	"github.com/RahmaYasser/go-gin/middleware"
	"github.com/RahmaYasser/go-gin/service"
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
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
	server.Use(gin.Recovery(), middleware.Logger(), middleware.BasicAuth(), gindump.Dump())

	server.Static("/css", "./template/css")
	server.LoadHTMLGlob("templates/*.html")

	apiRoutes := server.Group("/api")
	{
		apiRoutes.GET("/videos", func(context *gin.Context) {
			context.JSON(200, videoController.FindAll())
		})
		apiRoutes.POST("/videos", func(context *gin.Context) {

			err := videoController.Save(context)
			if err != nil {
				context.JSON(400, gin.H{"error": err.Error()})
				return
			}
			context.JSON(200, gin.H{
				"message": "ok",
			})
		})
	}
	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}
	server.Run("localhost:8080")
}
