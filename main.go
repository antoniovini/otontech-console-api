package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	// router.GET("/test", ())
	router.Run("localhost:8000")
}
