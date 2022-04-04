package commands

import (
	"net/http"
	"otontech/console-api/models"
	"otontech/console-api/utils/middlewares"
	"sort"

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
	routes.POST("/",
		middlewares.JwtAuthMiddleware(true),
		middlewares.RoleMiddleware([]string{"admin"}),
		h.CreateCommand,
	)
	routes.GET("/:id", h.GetCommand)
	routes.PATCH("/:id",
		middlewares.JwtAuthMiddleware(true),
		middlewares.RoleMiddleware([]string{"admin"}),
		h.UpdateCommand,
	)
	routes.DELETE(
		"/:id",
		middlewares.JwtAuthMiddleware(true),
		middlewares.RoleMiddleware([]string{"admin"}),
		h.DeleteCommand,
	)
}

func (h CommandHandler) CreateCommand(c *gin.Context) {
	var input models.CreateCommandInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	command := models.Command{Description: input.Description, Activator: input.Activator}

	if results := h.DB.Create(&command); results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": results.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": command})
}

func (h CommandHandler) GetCommands(c *gin.Context) {
	userContext, exists := c.Get("user")
	var commands []models.Command

	h.DB.Model(&models.Command{}).Preload("Roles").Find(&commands)

	var userCommands []models.Command
	var userRoles []string

	if exists {
		u := userContext.(models.User)
		for i := range u.Roles {
			userRoles = append(userRoles, u.Roles[i].Name)
		}
	} else {
		userRoles = append(userRoles, "default")
	}

	for i := range commands {
		for j := range commands[i].Roles {
			idx := sort.Search(len(userRoles), func(k int) bool {
				return userRoles[k] == commands[i].Roles[j].Name
			})

			if idx < len(userRoles) && userRoles[idx] == commands[i].Roles[j].Name {
				userCommands = append(userCommands, commands[i])
				break
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": userCommands})
}

func (h CommandHandler) GetCommand(c *gin.Context) {
	var command models.Command

	if err := h.DB.Where("unique_id = ?", c.Param("id")).First(&command).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": command})
}

func (h CommandHandler) UpdateCommand(c *gin.Context) {
	var command models.Command
	if err := h.DB.Where("unique_id = ?", c.Param("id")).First(&command).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input models.UpdateCommandInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if results := h.DB.Model(&command).Updates(&models.Command{Activator: input.Activator, Description: input.Description, Action: input.Action}); results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": results.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": command})
}

func (h CommandHandler) DeleteCommand(c *gin.Context) {
	var command models.Command
	if err := h.DB.Where("unique_id = ?", c.Param("id")).First(&command).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.DB.Delete(&command)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
