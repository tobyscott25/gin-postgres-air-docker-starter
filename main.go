package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", healthCheckHandler)
	router.Run(":8080")
}

func healthCheckHandler(c *gin.Context) {
	var superSecretKey string = os.Getenv("SUPER_SECRET_KEY")

	if superSecretKey == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Missing environment variables",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"healthy": true,
		"secret":  superSecretKey,
	})
}
