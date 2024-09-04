package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tobyscott25/gin-air-docker-boilerplate/src/di"
)

func InitialiseAppRouter(router *gin.Engine, deps *di.Dependencies) {
	v1 := router.Group("/v1")
	v1.Use(gin.Recovery())

	v1.GET("/health", HandleHealthCheck(deps))
}
