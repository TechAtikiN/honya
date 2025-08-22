package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/techatikin/backend/controller"
	"github.com/techatikin/backend/repositories"
	"github.com/techatikin/backend/service"
)

type BookRouter struct {
	app  *fiber.App
	ctrl controller.BookController
}

func NewBookRouter(app *fiber.App) *BookRouter {
	repo := repositories.NewBookRepository()
	service := service.NewBookService(repo)
	ctrl := controller.NewBookController(service)

	return &BookRouter{
		app:  app,
		ctrl: ctrl,
	}
}

func (r *BookRouter) Setup(api fiber.Router) {
	booksRoutes := api.Group("/books")

	booksRoutes.Get("/", r.ctrl.GetBooks)
	booksRoutes.Get("/:id", r.ctrl.GetBookByID)
	booksRoutes.Post("/", r.ctrl.CreateBook)
	// add put endpoint
	booksRoutes.Patch("/:id", r.ctrl.UpdateBook)
	booksRoutes.Delete("/:id", r.ctrl.DeleteBook)
}
