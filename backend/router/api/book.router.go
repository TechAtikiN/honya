package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/techatikin/backend/controller"
	"github.com/techatikin/backend/repository"
	"github.com/techatikin/backend/service"
)

type BookRouter struct {
	app  *fiber.App
	ctrl controller.BookController
}

func NewBookRouter(app *fiber.App) *BookRouter {
	repo := repository.NewBookRepository()
	s3repo := repository.NewS3Repository()
	service := service.NewBookService(repo, s3repo)
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
	booksRoutes.Patch("/:id", r.ctrl.UpdateBook)
	booksRoutes.Delete("/:id", r.ctrl.DeleteBook)
}
