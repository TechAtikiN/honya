package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/techatikin/backend/config"
	"github.com/techatikin/backend/utils"
)

func SeedBooksAPI(ctx *fiber.Ctx) error {
	if err := utils.SeedBooks(config.DB.Db); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Books seeded successfully or already exist",
	})
}
