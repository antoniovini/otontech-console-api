package middlewares

import (
	"net/http"
	"otontech/console-api/models"
	"otontech/console-api/pkg/auth"
	"otontech/console-api/utils/token"
	"sort"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func JwtAuthMiddleware(required bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)

		if err != nil && required {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		if err == nil {
			u := models.User{}
			auth.CurrentUser(c, &u)

			c.Set("user", u)
		}

		c.Next()
	}
}

func RoleMiddleware(roles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		u := c.MustGet("user").(models.User)

		var userRoles []string
		for i := range u.Roles {
			userRoles = append(userRoles, u.Roles[i].Name)
		}

		sort.Strings(roles)
		sort.Strings(userRoles)

		var idx int
		for i := range roles {
			idx = sort.SearchStrings(userRoles, roles[i])
			if len(userRoles) > idx && userRoles[idx] == roles[i] {
				c.Next()
				return
			}
		}

		c.String(http.StatusUnauthorized, "Unauthorized")
		c.Abort()
		return
	}
}
