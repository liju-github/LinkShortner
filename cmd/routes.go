package main

import (
	"linkshortner/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

var Server *gin.Engine

func ServerStart() {
	Server = gin.Default()
	Server.LoadHTMLGlob("../template/*.html")
	Routes()
	err := Server.Run(":8080")
	if err != nil {
		panic(err)
	}

}

func Routes() {
	Server.GET("/home",service.LoadHomeHTML)
	Server.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
	Server.GET("/url/all", service.GetURLs)
	Server.GET("/url/short/:url")
	Server.GET("/:key", service.GetRedirectURL)
	Server.GET("/url/stats/:shorturl")
	Server.GET("/env", service.Env)
}
