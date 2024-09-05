package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tobyscott25/gin-air-docker-boilerplate/src/config"
	"github.com/tobyscott25/gin-air-docker-boilerplate/src/controller"
	"github.com/tobyscott25/gin-air-docker-boilerplate/src/di"
)

func main() {
	config := config.Load()

	dependencies, err := di.InitDependencies(config)
	if err != nil {
		log.Fatal(err.Error())
	}

	router := gin.Default()

	controller.InitialiseAppRouter(router, dependencies)

	portStr := strconv.Itoa(int(config.Port))
	router.Run(fmt.Sprintf(":%s", portStr))
}
