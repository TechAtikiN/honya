package controller

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/techatikin/backend/dtos"
	"github.com/techatikin/backend/errors"
	"github.com/techatikin/backend/service"
	"github.com/techatikin/backend/utils"
)

type TBookController struct {
	bookService service.TBookService
}

func BookController(service service.TBookService) TBookController {
	return TBookController{service}
}

// GetBooks godoc
// @Summary Get list of all books
// @Description Get paginated list of books with optional filters
// @Tags books
// @Accept json
// @Produce json
// @Param query query string false "Search query"
// @Param offset query integer false "Offset for pagination" default(0)
// @Param limit query integer false "Limit for pagination" default(10)
// @Param category query string false "Filter by category"
// @Param publication_year query integer false "Filter by publication year"
// @Param rating query number false "Filter by rating"
// @Param pages query integer false "Filter by number of pages"
// @Param sort query string false "Sort order (asc/desc)"
// @Success 200 {object} dtos.BookListResponse "Books fetched successfully"
// @Router /books [get]
func (c *TBookController) GetBooks(ctx *fiber.Ctx) error {
	params := dtos.BookQueryParams{
		Query:           ctx.Query("query"),
		Offset:          utils.ParseInt(ctx.Query("offset"), utils.DefaultOffset),
		Limit:           utils.ParseInt(ctx.Query("limit"), utils.DefaultLimit),
		Category:        strings.ToLower(ctx.Query("category")),
		PublicationYear: utils.ParseInt(ctx.Query("publication_year"), utils.DefaultPublicationYear),
		Rating:          utils.ParseFloat(ctx.Query("rating"), utils.DefaultRating),
		Pages:           utils.ParseInt(ctx.Query("pages"), utils.DefaultPages),
		Sort:            strings.ToLower(ctx.Query("sort")),
	}

	books, meta, err := c.bookService.GetBooks(params)
	if err != nil {
		return err
	}

	result := dtos.ToBookListResponse(books, *meta)
	return ctx.Status(fiber.StatusOK).JSON(result)
}

// GetBookByID godoc
// @Summary Get book by ID
// @Description Get a single book by its ID
// @Tags books
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Success 200 {object} dtos.BookResponse "Book fetched successfully"
// @Failure 400 {object} errors.ErrorResponse "Invalid ID format"
// @Failure 404 {object} errors.ErrorResponse "Book not found"
// @Router /books/{id} [get]
func (c *TBookController) GetBookByID(ctx *fiber.Ctx) error {
	id, err := utils.ParseUUIDParam(ctx, "id")
	if err != nil {
		return err
	}

	book, err := c.bookService.GetBookByID(id)
	if err != nil {
		return err
	}

	result := dtos.ToBookResponse(book)
	return ctx.Status(fiber.StatusOK).JSON(result)
}

// CreateBook godoc
// @Summary Create a new book
// @Description Create a new book with the provided details
// @Tags books
// @Accept json
// @Produce json
// @Param book body dtos.BookCreateRequest true "Book creation payload"
// @Success 201 {object} dtos.BookResponse "Book created successfully"
// @Failure 400 {object} errors.ErrorResponse "Invalid input data"
// @Router /books [post]
func (c *TBookController) CreateBook(ctx *fiber.Ctx) error {
	var reqData dtos.BookCreateRequest
	if err := ctx.BodyParser(&reqData); err != nil {
		return errors.NewBadRequestError("Invalid JSON body")
	}

	book, err := c.bookService.CreateBook(&reqData)
	if err != nil {
		return err
	}

	result := dtos.ToBookResponse(book)
	return ctx.Status(fiber.StatusCreated).JSON(result)
}

// UpdateBook godoc
// @Summary Update an existing book
// @Description Update book details by its ID
// @Tags books
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Param book body dtos.BookUpdateRequest true "Book update payload"
// @Success 200 {object} dtos.BookResponse "Book updated successfully"
// @Failure 400 {object} errors.ErrorResponse "Invalid input data"
// @Failure 404 {object} errors.ErrorResponse "Book not found"
// @Router /books/{id} [put]
func (c *TBookController) UpdateBook(ctx *fiber.Ctx) error {
	id, err := utils.ParseUUIDParam(ctx, "id")
	if err != nil {
		return err
	}

	var requestData dtos.BookUpdateRequest
	if err := ctx.BodyParser(&requestData); err != nil {
		return errors.NewBadRequestError("Invalid JSON body")
	}

	if requestData.Isbn != nil {
		return errors.NewBadRequestError("ISBN cannot be updated once set")
	}

	updatedBook, err := c.bookService.UpdateBook(id, &requestData)
	if err != nil {
		return err
	}

	result := dtos.ToBookResponse(updatedBook)
	return ctx.Status(fiber.StatusOK).JSON(result)
}

// DeleteBook godoc
// @Summary Delete a book
// @Description Delete a book by its ID
// @Tags books
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Success 200 {object} map[string]string "Book deleted successfully"
// @Failure 400 {object} errors.ErrorResponse "Invalid ID format"
// @Failure 404 {object} errors.ErrorResponse "Book not found"
// @Router /books/{id} [delete]
func (c *TBookController) DeleteBook(ctx *fiber.Ctx) error {
	id, err := utils.ParseUUIDParam(ctx, "id")
	if err != nil {
		return err
	}

	if err := c.bookService.DeleteBook(id); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Book deleted successfully",
	})
}
