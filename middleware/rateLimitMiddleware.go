package middleware

import (
    "net/http"
    "sync"
    "time"

    "github.com/gin-gonic/gin"
)


type RateLimiter struct {
    visits map[string]time.Time
    lock   sync.Mutex
}


func NewRateLimiter() *RateLimiter {
    return &RateLimiter{
        visits: make(map[string]time.Time),
    }
}


func RateLimitMiddleware(limiter *RateLimiter, limit time.Duration) gin.HandlerFunc {
    return func(c *gin.Context) {
        limiter.lock.Lock()
        defer limiter.lock.Unlock()

        
        ip := c.ClientIP()

        
        if lastVisit, ok := limiter.visits[ip]; ok {
            
            if time.Since(lastVisit) < limit {
                c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests"})
                return
            }
        }

        
        limiter.visits[ip] = time.Now()
        c.Next()
    }
}
