package main

import (
	"auth/routers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routers.SetupRouter(router)
	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
