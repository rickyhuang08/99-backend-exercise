package middleware

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rickyhuang08/99-backend-exercise/helpers"
	"github.com/rickyhuang08/99-backend-exercise/pkg/auth"
)

type LoggerMiddleware struct {
	Time helpers.TimeProvider
}

func NewLoggerMiddleWare(time helpers.TimeProvider) *LoggerMiddleware {
	return &LoggerMiddleware{
		Time: time,
	}
}



// LoggerMiddleware logs all incoming requests
func(m *LoggerMiddleware) LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := m.Time.Now()
		requestID := m.Time.Now().UnixNano() // Generate simple request ID

		// Store request ID in context
		ctx := context.WithValue(c.Request.Context(), auth.RequestIDKey, requestID)
		c.Request = c.Request.WithContext(ctx)

		// Process request
		c.Next()

		// Calculate execution time
		duration := m.Time.Since(startTime)
		status := c.Writer.Status()

		// Log request details
		log.Printf(
			"[LoggerMiddleware][Request ID: %d] %s %s - %d %s", 
			requestID, 
			c.Request.Method, 
			c.Request.URL.Path, 
			status, 
			duration,
		)
	}
}