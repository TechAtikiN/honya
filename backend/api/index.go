package handler

import (
	"honya/backend/config"
	"honya/backend/middleware"
	"honya/backend/router"
	"log"
	"net/http"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// Handler is the entry point for Vercel Serverless Functions.
func Handler(w http.ResponseWriter, r *http.Request) {
	// The logic is similar to your main.go, but without the server startup.
	// You may want to handle environment config and database connections more gracefully
	// in a serverless context (e.g., connect on-demand or use connection pooling).
	env, err := config.GetEnvConfig()
	if err != nil {
		log.Printf("Failed to get environment configuration: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	app := fiber.New(fiber.Config{
		AppName:      "Honya API",
		ErrorHandler: middleware.ErrorHandler,
	})

	cfg := swagger.Config{
		BasePath: "/",
		FilePath: "./docs/swagger.json",
		Path:     "docs",
		Title:    "Honya | API Documentation",
		CacheAge: 60,
	}

	app.Use(swagger.New(cfg))
	app.Use(cors.New())

	// Note: Centralized logging like this may not be necessary on Vercel
	// as it handles logging automatically.
	app.Use(config.SetupLogger(env.LogStack, env.LogRetention))

	config.ConnectToDatabase(env.DatabaseURL)

	router.Setup(app)

	// Use the Fiber adapter to process the request and write the response
	adaptor.FiberApp(app)(w, r)
}
