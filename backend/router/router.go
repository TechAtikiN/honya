package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/techatikin/backend/middleware"
	"github.com/techatikin/backend/router/api"
)

type Router struct {
	app          *fiber.App
	healthRouter *api.HealthRouter
	bookRouter   *api.BookRouter
	reviewRouter *api.ReviewRouter
	seedRouter   *api.SeedRouter
}

func New(app *fiber.App) *Router {
	return &Router{
		app:          app,
		healthRouter: api.NewHealthRouter(app),
		bookRouter:   api.NewBookRouter(app),
		reviewRouter: api.NewReviewRouter(app),
		seedRouter:   api.NewSeedRouter(app),
	}
}

func Setup(app *fiber.App) {
	router := New(app)

	api := app.Group("/api", middleware.RateLimiter())

	router.healthRouter.Setup(api)
	router.bookRouter.Setup(api)
	router.reviewRouter.Setup(api)
	router.seedRouter.Setup(api)
}
