package api

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/api/middleware"
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humagin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// corsAllowedOrigins reads a comma-separated allow-list from CORS_ALLOWED_ORIGINS
// so it can be set per-environment (dev/staging/prod) without a code change or
// redeploy. Only dev gets a built-in fallback (the local Vite server) - failing
// fast anywhere else beats silently rejecting every cross-origin request, which
// is a confusing way to discover a missing env var.
func corsAllowedOrigins() []string {
	if raw := os.Getenv("CORS_ALLOWED_ORIGINS"); raw != "" {
		origins := strings.Split(raw, ",")
		for i, origin := range origins {
			origins[i] = strings.TrimSpace(origin)
		}
		return origins
	}

	environment := os.Getenv("ENVIRONMENT")
	if environment == "dev" || environment == "development" {
		return []string{"http://localhost:5173"}
	}

	log.Fatal("CORS_ALLOWED_ORIGINS is not set; refusing to start (a missing value here " +
		"would otherwise silently reject every cross-origin request instead of failing loudly)")
	return nil
}

type Server struct {
	*gin.Engine
	API huma.API
}

func NewServer() *Server {
	fmt.Println("Creating Server...")

	err := godotenv.Load(".dev.env")
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	var router *gin.Engine

	environment := os.Getenv("ENVIRONMENT")

	switch environment {
	case "dev", "development":
		router = gin.Default()
		fmt.Println("Running in DEVELOPMENT mode")

	case "test", "testing":
		gin.SetMode(gin.TestMode)
		router = gin.New()
		router.Use(gin.Recovery())
		fmt.Println("Running in TESTING mode")

	default:
		gin.SetMode(gin.ReleaseMode)
		router = gin.New()
		router.Use(gin.Recovery())
		fmt.Println("Running in PRODUCTION mode")
	}
	router.Use(cors.New(cors.Config{
		AllowOrigins:     corsAllowedOrigins(),
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Global, per-IP, in-memory rate limit - a coarse baseline against abuse/
	// scripted hammering. Route-specific tighter limits can be layered on top
	// later if a particular endpoint needs it.
	router.Use(middleware.NewRateLimiter(100, time.Minute))

	// Huma config with OpenAPI spec generation
	config := huma.DefaultConfig("Service Base API", "1.0")
	config.OpenAPI.Info.Description = "A generic service base with JWT authentication"
	config.OpenAPIPath = "/openapi"
	config.DocsPath = "/docs"
	config.SchemasPath = "/schemas"
	backendURL := os.Getenv("BACKEND_URL")

	config.Servers = []*huma.Server{
		{
			URL:         backendURL,
			Description: fmt.Sprintf("%s Environment Server", environment),
		},
	}

	if config.OpenAPI.Components.SecuritySchemes == nil {
		config.OpenAPI.Components.SecuritySchemes = make(map[string]*huma.SecurityScheme)
	}

	config.OpenAPI.Components.SecuritySchemes["BearerAuth"] = &huma.SecurityScheme{
		Type:         "http",
		Scheme:       "bearer",
		BearerFormat: "JWT",
		Description:  "Type \"Bearer\" followed by a space and JWT token",
	}

	api := humagin.New(router, config)

	s := &Server{
		Engine: router,
		API:    api,
	}

	fmt.Println("Server Creation Succeeded")
	return s
}
