package database

import (
	"Go-RestApi-Books/models"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var db *gorm.DB
var err error

func StartDB() {
	//Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	//Get database configuration from environment variables
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Create database connection string
	config := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	//Connect to database
	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to database: ", err)
	}
	fmt.Println("Database connected")
	db.AutoMigrate(models.Book{})
}

func GetDB() *gorm.DB {
	return db
}
