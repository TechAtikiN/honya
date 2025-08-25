package api

import (
	"honya/backend/controller"
	"honya/backend/repository"
	"honya/backend/service"

	"github.com/gofiber/fiber/v2"
)

type ReviewRouter struct {
	app  *fiber.App
	ctrl controller.ReviewController
}

func NewReviewRouter(app *fiber.App) *ReviewRouter {
	repo := repository.NewReviewRepository()
	service := service.NewReviewService(repo)
	ctrl := controller.NewReviewController(service)

	return &ReviewRouter{
		app:  app,
		ctrl: ctrl,
	}
}

func (r *ReviewRouter) Setup(api fiber.Router) {
	reviewRoutes := api.Group("/reviews")

	reviewRoutes.Get("/", r.ctrl.GetAllReviews)
	reviewRoutes.Get("/:id", r.ctrl.GetReviewByID)
	reviewRoutes.Get("/book/:book_id", r.ctrl.GetReviewsByBookID)
	reviewRoutes.Post("/", r.ctrl.CreateReview)
	reviewRoutes.Patch("/:id", r.ctrl.UpdateReview)
	reviewRoutes.Delete("/:id", r.ctrl.DeleteReview)
}
