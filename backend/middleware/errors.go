package middleware

import (
	"honya/backend/errors"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	if appErr, ok := err.(*errors.AppError); ok {
		return c.Status(appErr.Code).JSON(fiber.Map{
			"error":   appErr.Message,
			"details": appErr.Err.Error(),
		})
	}

	if err.Error() == "Cannot "+c.Method()+" "+c.Path() {
		return c.Status(fiber.StatusMethodNotAllowed).JSON(fiber.Map{
			"error": "Method Not Allowed",
		})
	}

	// Handle default error
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error":   "Something went wrong",
		"details": err.Error(),
	})
}
