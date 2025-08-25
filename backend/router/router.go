package router

import (
	"honya/backend/middleware"
	"honya/backend/router/api"

	"github.com/gofiber/fiber/v2"
)

type Router struct {
	app             *fiber.App
	healthRouter    *api.HealthRouter
	bookRouter      *api.BookRouter
	reviewRouter    *api.ReviewRouter
	seedRouter      *api.SeedRouter
	urlRouter       *api.UrlRouter
	dashboardRouter *api.DashboardRouter
}

func New(app *fiber.App) *Router {
	return &Router{
		app:             app,
		healthRouter:    api.NewHealthRouter(app),
		bookRouter:      api.NewBookRouter(app),
		reviewRouter:    api.NewReviewRouter(app),
		seedRouter:      api.NewSeedRouter(app),
		urlRouter:       api.NewUrlRouter(app),
		dashboardRouter: api.NewDashboardRouter(app),
	}
}

func Setup(app *fiber.App) {
	router := New(app)

	api := app.Group("/api", middleware.RateLimiter())

	router.healthRouter.Setup(api)
	router.bookRouter.Setup(api)
	router.reviewRouter.Setup(api)
	router.seedRouter.Setup(api)
	router.urlRouter.Setup(api)
	router.dashboardRouter.Setup(api)
}
