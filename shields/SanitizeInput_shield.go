package shields

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/microcosm-cc/bluemonday"
)

func SanitizeInput() gin.HandlerFunc {
	return func(c *gin.Context) {

		if c.Request.Method != "POST" || c.ContentType() != "application/json" {
			c.Next()
			return
		}
		var input map[string]interface{}
		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			c.Abort()
			return
		}
		//Initializing bluemonday sanitizer
		p := bluemonday.StrictPolicy()
		//Check all string fields
		for key, value := range input {
			if str, ok := value.(string); ok {
				//Check for sql injection pattern
				dangerous := []string{"SELECT", "DROP", "INSERT", "DELETE", ";", "--"}
				for _, bad := range dangerous {
					if strings.Contains(strings.ToUpper(str), strings.ToUpper(bad)) {
						c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
						c.Abort()
						return
					}
				}
				//sanitiza HTML
				sanitized := p.Sanitize(str)
				if sanitized != str {
					c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input: Deteced a HTML content"})
					c.Abort()
					return
				}
				input[key] = sanitized
			}
		}

		//update request with sanitized input
		c.Set("sanitizedInput", input)
		c.Next()

	}
}
