package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/techatikin/backend/controllers"
)

func SetupBooksRouter(api fiber.Router, ctrl *controllers.TBookController) {
	booksRoutes := api.Group("/books")

	booksRoutes.Get("/", ctrl.GetBooks)
	booksRoutes.Get("/:id", ctrl.GetBookByID)
	booksRoutes.Post("/", ctrl.CreateBook)
	booksRoutes.Put("/:id", ctrl.UpdateBook)
	booksRoutes.Delete("/:id", ctrl.DeleteBook)
}
