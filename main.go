package main

import (
	"fmt"
	"github.com/freeCodeCamp-Samples/GoAPI/server"
	"github.com/freeCodeCamp-Samples/GoAPI/storage"
	"github.com/gin-gonic/gin"
)

func main() {
	// create a gin server
	serv := gin.Default()
	// specify the path
	serv.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "This is a URL Shortener API: We will shorten your URL for you.",
		})
		serv.POST("/shorten-url", func(ctx *gin.Context) {
			server.ShortenURL(ctx)
		})
		serv.GET("/:shortUrl", func(ctx *gin.Context) {
			server.RedirectToOriginalUrl()
		})
		storage.InitStore()
	})
	// set the server to a port
	err := serv.Run(":8080")
	if err != nil {
		panic(fmt.Sprintf("Server failed... \nError: %v", err))
	}
}
