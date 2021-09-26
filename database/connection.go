package database

import (
	"fmt"
	"log"
	"os"

	"github.com/aysf/go_learn-jw_auth/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db_username := os.Getenv("DB_USERNAME")
	db_password := os.Getenv("DB_PASSWORD")
	connectionString := fmt.Sprintf("%s:%s@/goauth", db_username, db_password)
	connection, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic("could not connect database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}
