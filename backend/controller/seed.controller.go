package controller

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/techatikin/backend/config"
	"github.com/techatikin/backend/utils"
)

func SeedBooksAPI(ctx *fiber.Ctx) error {
	// Check if database connection is valid
	if config.DB.Db == nil {
		log.Println("Database connection is not initialized.")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Database connection is not initialized.",
		})
	}

	// Call SeedBooks function
	if err := utils.SeedBooks(config.DB.Db); err != nil {
		log.Printf("Error seeding books: %v", err) // Log error for debugging
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Books seeded successfully or already exist",
	})
}
