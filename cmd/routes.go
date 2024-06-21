package main

import (
	"linkshortner/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ServerStart() {
	server := gin.Default()

	server.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
	server.GET("/url/all",service.GetURLs)
	server.GET("/url/short/:url")
	server.GET("/:key",service.GetRedirectURL)
	server.GET("/url/stats/:shorturl")
	server.GET("/env", service.Env)

	err := server.Run(":8080")
	if err != nil {
		panic(err)
	}

}
