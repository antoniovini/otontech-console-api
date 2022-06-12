package commands

import (
	"net/http"
	"otontech/console-api/models"
	"otontech/console-api/utils/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CommandHandler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &CommandHandler{
		DB: db,
	}

	routes := r.Group("/api/v1/commands")
	routes.GET(
		"/",
		middlewares.JwtAuthMiddleware(false),
		h.GetCommands,
	)
	// routes.POST("/",
	// 	middlewares.JwtAuthMiddleware(true),
	// 	middlewares.RoleMiddleware([]string{"admin"}),
	// 	h.CreateCommand,
	// )
	routes.GET(
		"/:id",
		middlewares.JwtAuthMiddleware(false),
		h.GetCommand,
	)
	// routes.PATCH("/:id",
	// 	middlewares.JwtAuthMiddleware(true),
	// 	middlewares.RoleMiddleware([]string{"admin"}),
	// 	h.UpdateCommand,
	// )
	// routes.DELETE(
	// 	"/:id",
	// 	middlewares.JwtAuthMiddleware(true),
	// 	middlewares.RoleMiddleware([]string{"admin"}),
	// 	h.DeleteCommand,
	// )
}

// func (h CommandHandler) CreateCommand(c *gin.Context) {
// 	var input models.CreateCommandInput
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	var role models.Role = models.Role{Name: input.Role}

// 	if err := h.DB.Model(&models.Role{}).Find(&role).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	command := models.Command{
// 		Description: input.Description,
// 		Activator:   input.Activator,
// 		Action:      input.Action,
// 		Role:        role,
// 		Args:        input.Args,
// 	}

// 	if results := h.DB.Create(&command); results.Error != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": results.Error.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": command})
// }

func (h CommandHandler) GetCommands(c *gin.Context) {
	userContext, exists := c.Get("user")
	var commands []models.Command

	h.DB.Model(&models.Command{}).Preload("Role").Preload("Steps").Preload("Steps.Params").Preload("Args").Find(&commands)

	var userCommands []models.Command
	var userLevel uint = 0

	if exists {
		u := userContext.(models.User)
		userLevel = u.Role.Level
	}

	for _, c := range commands {
		if c.Role.Level <= userLevel {
			userCommands = append(userCommands, c)
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": userCommands})
}

func (h CommandHandler) GetCommand(c *gin.Context) {
	userContext, exists := c.Get("user")
	var command models.Command

	if err := h.DB.Preload("Role").Preload("Steps").Preload("Args").Where("unique_id = ?", c.Param("id")).First(&command).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var userLevel uint = 0

	if exists {
		u := userContext.(models.User)
		userLevel = u.Role.Level
	}

	if command.Role.Level > userLevel {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": command})
}

// func (h CommandHandler) UpdateCommand(c *gin.Context) {
// 	var command models.Command
// 	if err := h.DB.Where("unique_id = ?", c.Param("id")).First(&command).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
// 		return
// 	}

// 	var input models.UpdateCommandInput
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	var roles []models.Role

// 	for i := range input.Roles {
// 		roles = append(roles, models.Role{Name: input.Roles[i]})
// 	}

// 	if err := h.DB.Model(&models.Role{}).Find(&roles).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	updatedCommand := models.Command{
// 		Activator:   input.Activator,
// 		Description: input.Description,
// 		Action:      input.Action,
// 		Steps:       input.Steps,
// 		Roles:       roles,
// 		Args:        input.Args,
// 	}

// 	if results := h.DB.Model(&command).Updates(&updatedCommand); results.Error != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": results.Error.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"data": command})
// }

// func (h CommandHandler) DeleteCommand(c *gin.Context) {
// 	var command models.Command
// 	if err := h.DB.Where("unique_id = ?", c.Param("id")).First(&command).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	h.DB.Delete(&command)
// 	c.JSON(http.StatusOK, gin.H{"data": true})
// }
