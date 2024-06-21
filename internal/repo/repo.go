package repo

import (
	"errors"
	"fmt"
	database "linkshortner/internal/db"
	"linkshortner/internal/models"
)

func GetURLsDB() ([]models.LinkShortner, error) {
	fmt.Printf("DB State: %p\n", &database.DB) // Debug log to check DB state
	var URLs []models.LinkShortner
	if err := database.DB.Find(&URLs).Error; err != nil {
		return nil, err
	}
	return URLs, nil
}


func GetRedirectURL(Key string) (string,error) {
	var result models.LinkShortner
	if err := database.DB.Where("url_key =?", Key).First(&result).Error; err!= nil {
		return "", errors.New("failed to get redirect url")
	}
	return result.OriginalLink, nil
}