package main

import (
	"log"
	"otontech/console-api/models"
	"otontech/console-api/pkg/auth"
	"otontech/console-api/pkg/commands"
	"otontech/console-api/pkg/programs"
	"otontech/console-api/pkg/roles"
	"otontech/console-api/utils/middlewares"

	_ "otontech/console-api/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Console API
// @version         1.0
// @description     API for console management.

// @host      localhost:8080
// @BasePath  /api/v1
func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	router := gin.Default()

	db := models.ConnectDatabase()

	router.Use(middlewares.CORSMiddleware())
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	commands.RegisterRoutes(router, db)
	auth.RegisterRoutes(router, db)
	roles.RegisterRoutes(router, db)
	programs.RegisterRoutes(router, db)

	router.Run("localhost:8000")
}
