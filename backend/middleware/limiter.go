package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/techatikin/backend/utils"
)

func RateLimiter() fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        utils.RateLimitMaxRequests,
		Expiration: utils.RateLimitExpiryDuration,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error": "Rate limit exceeded. Please try again later.",
			})
		},
	})
}
