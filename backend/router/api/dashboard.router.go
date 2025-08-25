package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/techatikin/backend/controller"
	"github.com/techatikin/backend/repository"
	"github.com/techatikin/backend/service"
)

type DashboardRouter struct {
	app  *fiber.App
	ctrl controller.DashboardController
}

func NewDashboardRouter(app *fiber.App) *DashboardRouter {
	bookRepo := repository.NewBookRepository()
	repoReview := repository.NewReviewRepository()
	service := service.NewDashboardService(bookRepo, repoReview)
	ctrl := controller.NewDashboardController(service)

	return &DashboardRouter{
		app:  app,
		ctrl: ctrl,
	}
}

func (r *DashboardRouter) Setup(api fiber.Router) {
	dashboardRoutes := api.Group("/dashboard")

	dashboardRoutes.Get("/books-data", r.ctrl.GetBooksData)
	dashboardRoutes.Get("/reviews-data", r.ctrl.GetReviewsData)
}
