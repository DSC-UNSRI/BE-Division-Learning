package middleware

import (
	"net/http"
	"playlist-app/database"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.GetHeader("Auth-Name")
		pass := c.GetHeader("Auth-Password")

		if name == "" || pass == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Missing credentials"})
			c.Abort()
			return
		}

		var user struct {
			ID uint
		}

		err := database.DB.
			Table("users").
			Select("id").
			Where("name = ? AND password = ?", name, pass).
			Take(&user).Error

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}
