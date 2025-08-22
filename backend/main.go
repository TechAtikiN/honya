package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/techatikin/backend/config"
	"github.com/techatikin/backend/controllers"
	"github.com/techatikin/backend/errors"
	"github.com/techatikin/backend/repositories"
	"github.com/techatikin/backend/routers"
	"github.com/techatikin/backend/services"
)

func main() {
	_ = godotenv.Load()

	db, err := config.ConnectToDatabase()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

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

	api := app.Group("/api")

	repo := repositories.BookRepository(db)

	svc := services.BookService(repo)

	ctrl := controllers.BookController(svc)

	api.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("API is running")
	})

	routers.SetupBooksRouter(api, ctrl)

	log.Println("Server starting on port 8080...")
	log.Fatal(app.Listen(":8080"))
}
