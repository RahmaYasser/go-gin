package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.Default()
	server.GET("/test", func(context *gin.Context) {
		context.JSONP(200, gin.H{
			"message": "OK",
		})
	})
	server.Run("localhost:8080")
}
