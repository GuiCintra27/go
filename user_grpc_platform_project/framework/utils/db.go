package utils

import (
	"log"
	"os"

	"github.com/GuiCintra27/go/user_grpc_platform_project/domain"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func ConnectDB() *gorm.DB {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dsn := os.Getenv("DSN")

	db, err := gorm.Open("postgres", dsn)

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		panic(err)
	}

	db.AutoMigrate(&domain.User{})

	return db
}