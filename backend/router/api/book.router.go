package api

import (
	"honya/backend/controller"
	"honya/backend/repository"
	"honya/backend/service"

	"github.com/gofiber/fiber/v2"
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
