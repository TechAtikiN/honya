package main

import (
	"log"
	"os"

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/techatikin/backend/config"
	"github.com/techatikin/backend/middleware"
	"github.com/techatikin/backend/router"
)

// @title Honya API Documentation
// @version 1.0
// @description API documentation for Honya - an online book library.
// @termsOfService http://swagger.io/terms/
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /api
func main() {
	_ = godotenv.Load()

	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		AppName:      "Honya API",
		ErrorHandler: middleware.ErrorHandler,
	})

	// Swagger configuration
	cfg := swagger.Config{
		BasePath: "/",
		FilePath: "./docs/swagger.json",
		Path:     "docs",
		Title:    "Honya | API Documentation",
		CacheAge: 60,
	}

	app.Use(swagger.New(cfg))

	app.Use(cors.New())

	app.Use(config.SetupLogger())

	config.ConnectToDatabase()

	router.Setup(app)

	log.Println("Server starting on port 8080...")
	log.Fatal(app.Listen(os.Getenv("SERVER_PORT")))
}
