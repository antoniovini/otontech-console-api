package programs

import (
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
	// userContext, exists := c.Get("user")
	var programs []models.Program

	h.DB.Model(&models.Program{}).Preload("Roles").Find(&programs)

	// var userPrograms []models.Program
	// var userRoles []string

	// if exists {
	// 	u := userContext.(models.User)
	// 	for i := range u.Roles {
	// 		userRoles = append(userRoles, u.Roles[i].Name)
	// 	}
	// } else {
	// 	userRoles = append(userRoles, "default")
	// }

	c.JSON(http.StatusOK, gin.H{"data": programs})
}

func (h ProgramHandler) GetProgram(c *gin.Context) {
	// userContext, exists := c.Get("user")
	var program models.Program

	if err := h.DB.Preload("Roles").Where("name = ?", c.Param("name")).First(&program).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Program not found!"})
		return
	}

	// var userRoles []string

	// if exists {
	// 	u := userContext.(models.User)
	// 	for i := range u.Roles {
	// 		userRoles = append(userRoles, u.Roles[i].Name)
	// 	}
	// } else {
	// 	userRoles = append(userRoles, "default")
	// }

	// for i := range program.Roles {
	// 	idx := sort.Search(len(userRoles), func(j int) bool {
	// 		return userRoles[j] == program.Roles[i].Name
	// 	})

	// 	if !(idx < len(userRoles) && userRoles[idx] == program.Roles[i].Name) {
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Program not found!"})
	// 		return
	// 	}
	// }

	// TODO - check user programs
	c.JSON(http.StatusOK, gin.H{"data": program})
}
