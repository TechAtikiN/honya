package api

import (
	"honya/backend/config"
	"honya/backend/controller"
	"honya/backend/service"

	"github.com/gofiber/fiber/v2"
)

type UrlRouter struct {
	app  *fiber.App
	ctrl controller.UrlController
}

func NewUrlRouter(app *fiber.App) *UrlRouter {
	env, _ := config.GetEnvConfig()

	service := service.NewUrlService(env.UrlCleanupOriginalDomain)
	ctrl := controller.NewUrlController(service)

	return &UrlRouter{
		app:  app,
		ctrl: ctrl,
	}
}

func (r *UrlRouter) Setup(api fiber.Router) {
	urlRoutes := api.Group("/url")

	urlRoutes.Post("/process-url", r.ctrl.ProcessUrl)
}
