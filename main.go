package main

import (
	"auth/internal/database"
	"auth/routers"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	router := gin.Default()
	routers.SetupRouter(router, db)
	err = router.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
