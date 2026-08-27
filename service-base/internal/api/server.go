package api

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humagin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

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
	// CORS configuration. Add your deployed frontend origin(s) here once you have one.
	allowedOrigins := []string{"http://localhost:5173", "http://localhost:8000"}
	if extra := os.Getenv("FRONTEND_URL"); extra != "" {
		allowedOrigins = append(allowedOrigins, extra)
	}
	router.Use(cors.New(cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

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
