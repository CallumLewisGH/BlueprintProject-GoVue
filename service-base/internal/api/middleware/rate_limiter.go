package middleware

import (
	"time"

	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	"github.com/gin-gonic/gin"
)

func keyFunc(c *gin.Context) string {
	return c.ClientIP()
}

func errorHandler(c *gin.Context, info ratelimit.Info) {
	c.JSON(429, gin.H{"error": "Too many requests. Try again in " + time.Until(info.ResetTime).String()})
}

// NewRateLimiter is in-memory and per-instance - no Redis or other shared
// store to run. If you deploy this to multiple instances behind a load
// balancer, the effective limit scales with instance count (each tracks its
// own counters) - an accepted tradeoff for not standing up shared
// infrastructure just for rate limiting.
func NewRateLimiter(requestsPer uint, timeUnit time.Duration) gin.HandlerFunc {
	store := ratelimit.InMemoryStore(&ratelimit.InMemoryOptions{
		Rate:  timeUnit,
		Limit: requestsPer,
	})

	return ratelimit.RateLimiter(store, &ratelimit.Options{
		ErrorHandler: errorHandler,
		KeyFunc:      keyFunc,
	})
}
