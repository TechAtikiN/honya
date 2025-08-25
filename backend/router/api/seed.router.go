package api

import (
	"honya/backend/controller"

	"github.com/gofiber/fiber/v2"
)

type SeedRouter struct {
	app  *fiber.App
	ctrl *SeedController
}

type SeedController struct{}

func NewSeedController() *SeedController {
	return &SeedController{}
}

func (s *SeedController) SeedData(ctx *fiber.Ctx) error {
	return controller.SeedBooksAPI(ctx)
}

func NewSeedRouter(app *fiber.App) *SeedRouter {
	ctrl := NewSeedController()
	return &SeedRouter{
		app:  app,
		ctrl: ctrl,
	}
}

func (r *SeedRouter) Setup(api fiber.Router) {
	seedRoutes := api.Group("/seed")

	seedRoutes.Post("/", r.ctrl.SeedData)
}
