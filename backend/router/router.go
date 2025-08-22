package router

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/techatikin/backend/router/api"
)

type Router struct {
	app          *fiber.App
	healthRouter *api.HealthRouter
	bookRouter   *api.BookRouter
}

func New(app *fiber.App) *Router {
	return &Router{
		app:          app,
		healthRouter: api.NewHealthRouter(app),
		bookRouter:   api.NewBookRouter(app),
	}
}

func Setup(app *fiber.App) {
	router := New(app)

	api := app.Group("/api", limiter.New(limiter.Config{
		Max:        20,
		Expiration: 1 * time.Minute,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error": "Rate limit exceeded. Please try again later.",
			})
		},
	}))

	router.healthRouter.Setup(api)
	router.bookRouter.Setup(api)
}
