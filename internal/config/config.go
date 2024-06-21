package config

import (
	"errors"
	"fmt"
	"linkshortner/internal/models"
	"log"
	"os"

	"github.com/joho/godotenv"
)


func LoadENV() {
	fmt.Print("hello")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func DatabaseCred() (*models.DatabaseCred, error) {
	User := os.Getenv("DBUSER")
	if User == "" {
		return nil, errors.New("DBUSER environment variable not set")
	}

	Password := os.Getenv("DBPASSWORD")
	if Password == "" {
		return nil, errors.New("DBPASSWORD environment variable not set")
	}

	Name := os.Getenv("DBNAME")
	if Name == "" {
		return nil, errors.New("DBNAME environment variable not set")
	}

	return &models.DatabaseCred{
		Name:     Name,
		User:     User,
		Password: Password,
	}, nil
}
