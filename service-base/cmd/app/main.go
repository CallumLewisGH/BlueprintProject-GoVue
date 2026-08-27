package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/database"
	"github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/api"
	"github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/api/auth"
	"github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/api/routes"
	"github.com/joho/godotenv"
	_ "github.com/lann/builder"
)

func main() {
	err := godotenv.Load(".dev.env")
	if err != nil {
		log.Printf("Warning: .dev.env file not found")
	}

	if os.Getenv("JWT_SECRET") == "" {
		log.Fatal("JWT_SECRET is not set; refusing to start (an empty secret would let anyone forge tokens)")
	}

	//Get Authentication Config
	auth.NewAuth()

	//Database Initialisation
	if err := database.GetDatabase().RunMigrations(); err != nil {
		log.Fatalf("Failed to run database migrations: %v", err)
	}

	//Creates new server instance
	srv := api.NewServer()

	//Route Registry (Huma registers user & database; auth stays as raw Gin for OAuth)
	routes.RegisterDatabaseRoutes(srv)
	routes.RegisterUserRoutes(srv)
	routes.RegisterAuthRoutes(srv)
	routes.RegisterUploadRoutes(srv)

	backendUrl := os.Getenv("BACKEND_URL")
	if os.Getenv("ENVIRONMENT") == "dev" {
		// Huma serves OpenAPI at /openapi.json, /openapi.yaml and docs at /docs
		println(fmt.Sprintf("\x1b]8;;%s\x1b\\%s\x1b]8;;\x1b\\", backendUrl+"/docs", "API Docs"))
	}

	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port, srv)

}

//Implementation Notes
// Queries IE repo.find, .first, all accept pointers &user and data is returned to that object with an err being returned from the function
// Commands IE repo.CreateOne, UpdateMany, DeleteOne, all accept actual values this way the input value is discaurded and another value is supplimented
// New models need to be added to the model_registry in the database package
// Change to prod mode for the database when deploying to prod => Also change GIN server from default

// Adding new stuff process:
// Create models:
//  - Create the db model
//  - Register the db model in the DB registry
//  - Create the DTO
//  - Create a converter from db model to DTO
//  - Create validation for the model
//  - Create requests for interacting with the user IE Create X Update X
//  - These requests should be validated against the validation created
//  - Create a repo if the object is queryable
//  - Create a specification if you want the query builder

// Create query:
//  - Create the query
//  - In the query create a queryFunc. This should have the following profile func(db *gorm.DB, ctx context.Context) (*model.DTO, error)
//  - In the queryFunc initialise the repo repos.NewModelRepo()
//  - Use this repo to query the database
//  - The database query will return the db model. This should be mapped to the DTO and returned
//  - This queryFunc should then be executed with cqrs.DbQuery(queryFunc)

// Create command:
//  - Create the command
//  - In the query create a commandFunc. This should have the following profile func(db *gorm.DB, ctx context.Context) (*model.DTO, error)
//  - In the commandFunc initialise the repo repos.NewModelRepo()
//  - Use this repo to command the database
//  - The database command will return the db model. This should be mapped to the DTO and returned
//  - This commandFunc should then be executed with cqrs.DbExecute(commandFunc)

// Todo:

// - Add roles & permissions
