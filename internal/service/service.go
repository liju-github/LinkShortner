package service

import (
	"linkshortner/internal/config"
	database "linkshortner/internal/db"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func GetURLs(c *gin.Context) {
	urls, err := database.GetURLsDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get all URLs"})
		return
	}
	c.JSON(http.StatusOK, urls)
}

func Env(ctx *gin.Context) {
	env, _ := config.DatabaseCred()
	ctx.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   env,
	})
}

func GetRedirectURL(c *gin.Context) {
	Key := c.Param("key")
	var RedirectURL *url.URL
	var err error

	// Attempt to parse the URL with the protocol
	RedirectURL, err = url.Parse("http://" + Key)
	if err != nil {
		// Handle error, possibly log it
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid URL format"})
		return
	}

	// Now RedirectURL is guaranteed to start with "http://"
	c.Redirect(http.StatusPermanentRedirect, RedirectURL.String())
}


func LoadHomeHTML(c *gin.Context)  {
	c.HTML(http.StatusOK,"index.html",nil)
}