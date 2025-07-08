package middleware

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rickyhuang08/99-backend-exercise/helpers"
)

// RateLimiter struct
type RateLimiter struct {
	Time     helpers.TimeProvider
	Visitors map[string]*visitor
	Mu       sync.Mutex
}

// Visitor data
type visitor struct {
	lastSeen time.Time
	tokens   int
}

// NewRateLimiter creates a new rate limiter instance
func NewRateLimiter(time helpers.TimeProvider) *RateLimiter {
	return &RateLimiter{
		Time: time,
		Visitors: make(map[string]*visitor),
	}
}

// RateLimitMiddleware limits request rate
func (rl *RateLimiter) RateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		rl.Mu.Lock()
		defer rl.Mu.Unlock()

		ip := c.ClientIP()
		v, exists := rl.Visitors[ip]

		if !exists {
			v = &visitor{lastSeen: rl.Time.Now(), tokens: 10}
			rl.Visitors[ip] = v
		}

		// Refill tokens
		elapsed := rl.Time.Since(v.lastSeen).Seconds()
		v.tokens += int(elapsed)
		if v.tokens > 10 {
			v.tokens = 10
		}

		// Check rate limit
		if v.tokens <= 0 {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests"})
			c.Abort()
			return
		}

		v.tokens--
		v.lastSeen = rl.Time.Now()

		// Store rate limiter data in context
		ctx := context.WithValue(c.Request.Context(), helpers.RateLimit, v.tokens)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
