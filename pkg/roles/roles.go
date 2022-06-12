package roles

// import (
// 	"net/http"
// 	"otontech/console-api/models"
// 	"otontech/console-api/utils/middlewares"
// 	"sort"

// 	"github.com/gin-gonic/gin"
// 	"gorm.io/gorm"
// )

// type RolesHandler struct {
// 	DB *gorm.DB
// }

// func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
// 	h := &RolesHandler{
// 		DB: db,
// 	}

// 	routes := r.Group("/api/v1/roles")
// 	routes.POST(
// 		"/give",
// 		middlewares.JwtAuthMiddleware(true),
// 		middlewares.RoleMiddleware([]string{"admin"}),
// 		h.GiveRole,
// 	)
// 	routes.POST(
// 		"/revoke",
// 		middlewares.JwtAuthMiddleware(true),
// 		middlewares.RoleMiddleware([]string{"admin"}),
// 		h.RevokeRole,
// 	)
// }

// func (h RolesHandler) GiveRole(c *gin.Context) {
// 	var input models.RoleManagmentInput
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	u, err := models.GetUserByUsername(input.Username)

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	roles := u.GetRolesString()

// 	sort.Strings(roles)
// 	idx := sort.SearchStrings(roles, input.Role)

// 	if len(roles) > idx && roles[idx] == input.Role {
// 		c.JSON(http.StatusOK, gin.H{"data": "role given to user."})
// 		return
// 	}

// 	var role models.Role
// 	if err := h.DB.Model(&models.Role{}).Where("name = ?", input.Role).First(&role).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if role.Level > u.MaxRoleLevel() {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Action cannot be completed"})
// 		return
// 	}

// 	h.DB.Model(&u).Association("Roles").Append(&role)

// 	c.JSON(http.StatusOK, gin.H{"data": "role given to user."})
// }

// func (h RolesHandler) RevokeRole(c *gin.Context) {
// 	var input models.RoleManagmentInput
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	u, err := models.GetUserByUsername(input.Username)

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	roles := u.GetRolesString()

// 	sort.Strings(roles)
// 	idx := sort.SearchStrings(roles, input.Role)

// 	if !(len(roles) > idx && roles[idx] == input.Role) {
// 		c.JSON(http.StatusOK, gin.H{"data": "revoked role for user."})
// 		return
// 	}

// 	var role models.Role
// 	if err := h.DB.Model(&models.Role{}).Where("name = ?", input.Role).First(&role).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if role.Level >= u.MaxRoleLevel() {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Action cannot be completed"})
// 		return
// 	}

// 	h.DB.Model(&u).Association("Roles").Delete(&role)

// 	c.JSON(http.StatusOK, gin.H{"data": "role given to user."})
// }
