package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tobyscott25/gin-air-docker-boilerplate/src/di"
)

// Swagger spec docs here

func HandleHealthCheck(dependencies *di.Dependencies) gin.HandlerFunc {
	return func(c *gin.Context) {
		var superSecretKey string = dependencies.Config.SuperSecretKey

		if superSecretKey == "" {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"healthy": true,
			"secret":  superSecretKey,
		})
	}
}
