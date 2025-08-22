package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/techatikin/backend/config"
	"github.com/techatikin/backend/errors"
	"github.com/techatikin/backend/router"
)

// @title Honya API
// @version 1.0
// @description API documentation for Honya - an online book library.
// @termsOfService http://swagger.io/terms/
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /api
func main() {
	_ = godotenv.Load()

	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			if appErr, ok := err.(*errors.AppError); ok {
				return ctx.Status(appErr.Code).JSON(fiber.Map{
					"error":   appErr.Message,
					"details": appErr.Err.Error(),
				})
			}

			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Something went wrong",
			})
		},
	})

	app.Use(cors.New())

	config.ConnectToDatabase()

	router.Setup(app)

	log.Println("Server starting on port 8080...")
	log.Fatal(app.Listen(":8080"))
}
