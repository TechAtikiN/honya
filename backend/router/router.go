package router

import (
	"github.com/gofiber/fiber/v2"
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

	api := app.Group("/api")

	router.healthRouter.Setup(api)
	router.bookRouter.Setup(api)
}
