package router

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router := gin.Default()
	
	initializeRoutes(router)

	router.Run(port)
}