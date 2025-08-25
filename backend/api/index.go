package handler

import (
	_ "embed"
	"log"
	"net/http"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"honya/backend/config"
	"honya/backend/middleware"
	"honya/backend/router"
)

//go:embed swagger/swagger.json
var swaggerJson []byte

// Handler is the entry point for Vercel Serverless Functions.
func Handler(w http.ResponseWriter, r *http.Request) {
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
		BasePath:    "/",
		FileContent: swaggerJson,
		Path:        "docs",
		Title:       "Honya | API Documentation",
		CacheAge:    60,
	}

	app.Use(swagger.New(cfg))
	app.Use(cors.New())

	app.Use(config.SetupLogger(env.LogStack, env.LogRetention))

	config.ConnectToDatabase(env.DatabaseURL)

	router.Setup(app)

	adaptor.FiberApp(app)(w, r)
}
