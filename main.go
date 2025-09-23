package main

import (
	"net/http"

	"github.com/Aruna456/WebShield/shields"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(shields.AuthShield())
	router.GET("/todos", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"todos": []string{"Learn Go", "Know about webshield", "Prepare demo"},
		})
	})
	router.POST("/todos", func(c *gin.Context) {
		var todo struct {
			Task string `json:"task"`
		}

		if err := c.BindJSON(&todo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Added task " + todo.Task})
	})
	router.Run(":7777")
}
