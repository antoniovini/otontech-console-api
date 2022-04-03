package main

import (
	commands "otontech/console-api/controllers"
	"otontech/console-api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	db := models.ConnectDatabase()
	commands.RegisterRoutes(router, db)

	router.Run("localhost:8000")
}
