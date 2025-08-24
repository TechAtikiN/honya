package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/techatikin/backend/config"
	"github.com/techatikin/backend/controller"
	"github.com/techatikin/backend/service"
)

type UrlRouter struct {
	app  *fiber.App
	ctrl controller.UrlController
}

func NewUrlRouter(app *fiber.App) *UrlRouter {
	originalDomain, _ := config.GetOriginalDomain()

	service := service.NewUrlService(originalDomain)
	ctrl := controller.NewUrlController(service)

	return &UrlRouter{
		app:  app,
		ctrl: ctrl,
	}
}

func (r *UrlRouter) Setup(api fiber.Router) {
	api.Get("/process-url", r.ctrl.ProcessUrl)
}
