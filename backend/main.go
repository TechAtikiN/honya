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

func main() {
	_ = godotenv.Load()

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

	app.Use(config.SetupLogger())

	config.ConnectToDatabase()

	router.Setup(app)

	defer config.CloseLogFile()

	log.Println("Server starting on port 8080...")
	log.Fatal(app.Listen(":" + os.Getenv("SERVER_PORT")))
}
