package api

import (
	"github.com/gin-gonic/gin"
	"example/gin-service-web/logger"
	"example/gin-service-web/middleware"
	"example/gin-service-web/exceptions"
	"go.uber.org/zap"
	"net/http"
)

func HealthCheck(c *gin.Context) {
	requestID, _ := c.Get("requestID")
	logger.Logger.Info("Request ID", zap.String("requestID", requestID.(string)))	
	c.JSON(200, gin.H{
		"requestID": requestID.(string),
		"message": "pong",
	})
}

func CulversHandler(c *gin.Context) {
	requestID, _ := c.Get("requestID")
	logger.Logger.Info("Request ID", zap.String("requestID", requestID.(string)))

	town := c.Param("town")
	if town == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing town parameter"})
		return
	}

	url := "https://www.culvers.com/restaurants/" + town

	logger.Logger.Info("Scraping Culver's website", zap.String("url", url))
    data, err := middleware.ScrapeWebsite(url)
    if err != nil {
		logger.Logger.Error("Failed to scrape Culver's website", zap.Error(err))
		if exceptions.IsNotFoundError(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Culver's website not found for the specified town"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scrape Culver's website"})
		return
    }
	logger.Logger.Info("Successfully scraped Culver's website", zap.String("flavor", data.Flavor))
	
    // c.JSON(http.StatusOK, data)	
	c.JSON(http.StatusOK, gin.H{
		"CulversData": data,
		"requestID": requestID.(string),
	})
}