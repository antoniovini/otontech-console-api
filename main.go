package main

import (
	"log"
	"otontech/console-api/models"
	"otontech/console-api/pkg/auth"
	"otontech/console-api/pkg/programs"
	"otontech/console-api/utils/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	router := gin.Default()

	db := models.ConnectDatabase()

	router.Use(middlewares.CORSMiddleware())

	auth.RegisterRoutes(router, db)
	// roles.RegisterRoutes(router, db)
	programs.RegisterRoutes(router, db)

	router.Run(":8081")
}
