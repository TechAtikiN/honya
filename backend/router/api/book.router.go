package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/techatikin/backend/controller"
	"github.com/techatikin/backend/repositories"
	"github.com/techatikin/backend/service"
)

type BookRouter struct {
	app            *fiber.App
	bookController controller.TBookController
}

func NewBookRouter(app *fiber.App) *BookRouter {
	bookRepo := repositories.BookRepository()
	bookService := service.BookService(bookRepo)
	bookController := controller.BookController(bookService)

	return &BookRouter{
		app:            app,
		bookController: bookController,
	}
}

func (r *BookRouter) Setup(api fiber.Router) {
	booksRoutes := api.Group("/books")

	booksRoutes.Get("/", r.bookController.GetBooks)
	booksRoutes.Get("/:id", r.bookController.GetBookByID)
	booksRoutes.Post("/", r.bookController.CreateBook)
	booksRoutes.Put("/:id", r.bookController.UpdateBook)
	booksRoutes.Delete("/:id", r.bookController.DeleteBook)
}
