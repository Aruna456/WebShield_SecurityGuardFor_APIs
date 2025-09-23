package shields

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthShield() gin.HandlerFunc {

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization Header required"})
			c.Abort()
			return
		}
		if authHeader != "token-placeholder" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		c.Next()
	}

}
