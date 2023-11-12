package logger

import (
    "go.uber.org/zap"
	"github.com/google/uuid"
	"github.com/gin-gonic/gin"
)

var Logger *zap.Logger

func InitLogger() {
    var err error
    Logger, err = zap.NewProduction()
    if err != nil {
        panic(err)
    }
}

func AttachUUID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := uuid.New().String()
		c.Set("requestID", requestID)

		// Add the request ID to the response headers if needed
		c.Writer.Header().Set("X-Request-ID", requestID)

		c.Next()
	}
}