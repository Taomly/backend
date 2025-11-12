package main

import (
	"auth/internal/database"
	"auth/internal/database/queries"
	"auth/routers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("CI") == "" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	err = queries.CreateTable(db)
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
