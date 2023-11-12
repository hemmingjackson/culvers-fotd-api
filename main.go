package main

import (
    "github.com/gin-gonic/gin"
	"example/gin-service-web/api"
	"example/gin-service-web/logger"
)

func main() {
	logger.InitLogger()

    router := gin.Default()
	router.Use(logger.AttachUUID())

	router.GET("/ping", api.HealthCheck)
    router.GET("/culvers/:town", api.CulversHandler)
	router.Run("localhost:8080")
}
