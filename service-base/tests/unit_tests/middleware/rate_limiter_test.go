package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/api/middleware"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRateLimiter(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.New()
	limiter := middleware.NewRateLimiter(1, time.Minute)
	router.Use(limiter)

	// Add a test endpoint
	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "success"})
	})

	// First request should succeed
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/test", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// Second request should be rate limited
	w = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodGet, "/test", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 429, w.Code)
}
