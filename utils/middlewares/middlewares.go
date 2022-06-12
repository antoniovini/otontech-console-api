package middlewares

import (
	"net/http"
	"otontech/console-api/models"
	"otontech/console-api/pkg/auth"
	"otontech/console-api/utils/token"

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

func RoleMiddleware(minLevel uint) gin.HandlerFunc {
	return func(c *gin.Context) {
		u := c.MustGet("user").(models.User)

		if minLevel <= u.Role.Level {
			c.Next()
			return
		}

		c.String(http.StatusUnauthorized, "Unauthorized")
		c.Abort()
	}
}
