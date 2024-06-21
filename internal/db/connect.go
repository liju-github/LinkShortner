package database

import (
	"fmt"
	"linkshortner/internal/config"
	"linkshortner/internal/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// Load database credentials from env file
	cred, err := config.DatabaseCred()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Connect to database
	dsn := fmt.Sprintf("%v:%v@tcp(127.0.0.1:3306)/%v?charset=utf8mb4&parseTime=True&loc=Local", cred.User, cred.Password, cred.Name)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	} else {
		log.Println("Connected to database successfully.")
	}

	// Auto-migrate the table in the database
	fmt.Println(&DB)
    AutoMigrate(DB)

}

func AutoMigrate(db *gorm.DB) {
	if err := db.AutoMigrate(&models.LinkShortner{}); err != nil {
		log.Fatal("unable to automigrate the database")
	}

}

func GetURLsDB() ([]models.LinkShortner, error) {
	fmt.Printf("DB State: %p\n", &DB) // Debug log to check DB state
	var URLs []models.LinkShortner
	if err := DB.Find(&URLs).Error; err != nil {
		return nil, err
	}
	return URLs, nil
}
