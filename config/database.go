package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Database() {
	err := godotenv.Load()
	
	if err != nil {
		log.Fatal("Failed to load .env file")
	}

	dsn := os.Getenv("DATABASE_URL")
  	db,err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect databse")
	}

	DB = db 
}