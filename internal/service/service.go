package service

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"linkshortner/internal/config"
	database "linkshortner/internal/db"
	"linkshortner/internal/models"
	"linkshortner/internal/repo"
	"net/http"
	"strings"

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

	URL, _ := repo.GetRedirectURL(Key)
	c.Redirect(http.StatusPermanentRedirect, URL)
	
}

func LoadHomeHTML(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func ProcessURL(c *gin.Context) {
	var url models.Request
	// Check Content-Type header to decide between ShouldBind or ShouldBindJSON
	contentType := c.GetHeader("Content-Type")
	if strings.HasPrefix(contentType, "application/json") {
		if err := c.ShouldBindJSON(&url); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	} else {
		// Assuming form data, bind accordingly
		if err := c.ShouldBind(&url); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	var key string
	for {
		key = generateRandomString(5)
		if !repo.KeyExistence(key) {
			break
		}
	}

	err := repo.AddURL(key, url.Url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// Constructing the short URL without the remote IP
	shortURL := fmt.Sprintf("http://localhost:8080"+"/%s", key)

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   shortURL,
	})
}

func generateRandomString(length int) string {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(b)
}
