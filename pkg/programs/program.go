package programs

import (
	"fmt"
	"net/http"
	"otontech/console-api/models"
	"otontech/console-api/utils/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProgramHandler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &ProgramHandler{
		DB: db,
	}

	routes := r.Group("/api/v1/programs")
	routes.GET(
		"/:name",
		middlewares.JwtAuthMiddleware(false),
		h.GetProgram,
	)
	routes.GET(
		"/",
		middlewares.JwtAuthMiddleware(false),
		h.GetPrograms,
	)
}

func (h ProgramHandler) GetPrograms(c *gin.Context) {
	userContext, exists := c.Get("user")
	var programs []models.Program

	h.DB.Model(&models.Program{}).Preload("Role").Find(&programs)

	var userPrograms []models.Program
	var userLevel uint = 0

	if exists {
		u := userContext.(models.User)
		userLevel = u.Role.Level
	}

	for _, p := range programs {
		if p.Role.Level <= userLevel {
			userPrograms = append(userPrograms, p)
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": userPrograms})
}

func (h ProgramHandler) GetProgram(c *gin.Context) {
	userContext, exists := c.Get("user")
	var program models.Program

	if err := h.DB.Preload("Role").Where("name = ?", c.Param("name")).First(&program).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Program not found!"})
		return
	}

	fmt.Println(program)

	var userLevel uint = 0

	if exists {
		u := userContext.(models.User)
		userLevel = u.Role.Level
	}

	if program.Role.Level > userLevel {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Program not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": program})
}
