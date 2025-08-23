package controller

import (
	"mime/multipart"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/techatikin/backend/dto"
	"github.com/techatikin/backend/errors"
	"github.com/techatikin/backend/service"
	"github.com/techatikin/backend/utils"
)

type BookController interface {
	GetBooks(ctx *fiber.Ctx) error
	GetBookByID(ctx *fiber.Ctx) error
	CreateBook(ctx *fiber.Ctx) error
	UpdateBook(ctx *fiber.Ctx) error
	DeleteBook(ctx *fiber.Ctx) error
}

type bookController struct {
	service service.BookService
}

func NewBookController(service service.BookService) BookController {
	return &bookController{service}
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
// @Success 200 {object} dto.BookListResponse "Books fetched successfully"
// @Router /books [get]
func (c *bookController) GetBooks(ctx *fiber.Ctx) error {
	params := dto.BookQueryParams{
		Query:           ctx.Query("query"),
		Offset:          utils.ParseInt(ctx.Query("offset"), utils.DefaultOffset),
		Limit:           utils.ParseInt(ctx.Query("limit"), utils.DefaultLimit),
		Category:        strings.ToLower(ctx.Query("category")),
		PublicationYear: utils.ParseInt(ctx.Query("publication_year"), utils.DefaultPublicationYear),
		Rating:          utils.ParseFloat(ctx.Query("rating"), utils.DefaultRating),
		Pages:           utils.ParseInt(ctx.Query("pages"), utils.DefaultPages),
		Sort:            strings.ToLower(ctx.Query("sort")),
	}

	books, meta, err := c.service.GetBooks(params)
	if err != nil {
		return err
	}

	result := dto.ToBookListResponse(books, *meta)
	return ctx.Status(fiber.StatusOK).JSON(result)
}

// GetBookByID godoc
// @Summary Get book by ID
// @Description Get a single book by its ID
// @Tags books
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Success 200 {object} dto.BookResponse "Book fetched successfully"
// @Failure 400 {object} errors.ErrorResponse "Invalid ID format"
// @Failure 404 {object} errors.ErrorResponse "Book not found"
// @Router /books/{id} [get]
func (c *bookController) GetBookByID(ctx *fiber.Ctx) error {
	id, err := utils.ParseUUIDParam(ctx, "id")
	if err != nil {
		return err
	}

	book, err := c.service.GetBookByID(id)
	if err != nil {
		return err
	}

	result := dto.ToBookResponse(book)
	return ctx.Status(fiber.StatusOK).JSON(result)
}

// CreateBook godoc
// @Summary Create a new book
// @Description Create a new book with the provided details
// @Tags books
// @Accept json
// @Produce json
// @Param book body dto.BookCreateRequest true "Book creation payload"
// @Success 201 {object} dto.BookResponse "Book created successfully"
// @Failure 400 {object} errors.ErrorResponse "Invalid input data"
// @Router /books [post]
func (c *bookController) CreateBook(ctx *fiber.Ctx) error {
	// Extract form values instead of BodyParser
	var reqData dto.BookCreateRequest
	reqData.Title = ctx.FormValue("title")
	reqData.Description = ctx.FormValue("description")
	reqData.Category = ctx.FormValue("category")
	reqData.PublicationYear, _ = strconv.Atoi(ctx.FormValue("publication_year"))
	reqData.Rating, _ = strconv.ParseFloat(ctx.FormValue("rating"), 64)
	reqData.Pages, _ = strconv.Atoi(ctx.FormValue("pages"))
	reqData.Isbn = ctx.FormValue("isbn")
	reqData.AuthorName = ctx.FormValue("author_name")

	// Get uploaded file
	var fileHeader *multipart.FileHeader
	file, err := ctx.FormFile("image")
	if err == nil {
		fileHeader = file
	}

	// Call service
	book, err := c.service.CreateBook(&reqData, fileHeader)
	if err != nil {
		return err
	}

	result := dto.ToBookResponse(book)
	return ctx.Status(fiber.StatusCreated).JSON(result)
}

// UpdateBook godoc
// @Summary Update an existing book
// @Description Update book details by its ID
// @Tags books
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Param book body dto.BookUpdateRequest true "Book update payload"
// @Success 200 {object} dto.BookResponse "Book updated successfully"
// @Failure 400 {object} errors.ErrorResponse "Invalid input data"
// @Failure 404 {object} errors.ErrorResponse "Book not found"
// @Router /books/{id} [put]
func (c *bookController) UpdateBook(ctx *fiber.Ctx) error {
	id, err := utils.ParseUUIDParam(ctx, "id")
	if err != nil {
		return err
	}

	var requestData dto.BookUpdateRequest
	var fileHeader *multipart.FileHeader

	contentType := ctx.Get("Content-Type")

	if strings.HasPrefix(contentType, "multipart/form-data") {
		// Parse text fields from form-data
		if title := ctx.FormValue("title"); title != "" {
			requestData.Title = &title
		}
		if description := ctx.FormValue("description"); description != "" {
			requestData.Description = &description
		}
		if category := ctx.FormValue("category"); category != "" {
			requestData.Category = &category
		}
		if year := ctx.FormValue("publication_year"); year != "" {
			yearInt, _ := strconv.Atoi(year)
			requestData.PublicationYear = &yearInt
		}
		if rating := ctx.FormValue("rating"); rating != "" {
			ratingFloat, _ := strconv.ParseFloat(rating, 64)
			requestData.Rating = &ratingFloat
		}
		if pages := ctx.FormValue("pages"); pages != "" {
			pagesInt, _ := strconv.Atoi(pages)
			requestData.Pages = &pagesInt
		}
		if author := ctx.FormValue("author_name"); author != "" {
			requestData.AuthorName = &author
		}

		// Parse file
		file, err := ctx.FormFile("image")
		if err == nil {
			fileHeader = file
		}
	} else {
		// JSON body
		if err := ctx.BodyParser(&requestData); err != nil {
			return errors.NewBadRequestError("Invalid JSON body")
		}
	}

	// ISBN cannot be updated
	if requestData.Isbn != nil {
		return errors.NewBadRequestError("ISBN cannot be updated once set")
	}

	// ✅ Run your validator
	if err := utils.ValidateBookUpdateRequest(&requestData); err != nil {
		return errors.NewBadRequestError(err.Error())
	}

	updatedBook, err := c.service.UpdateBook(id, &requestData, fileHeader)
	if err != nil {
		return err
	}

	result := dto.ToBookResponse(updatedBook)
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
func (c *bookController) DeleteBook(ctx *fiber.Ctx) error {
	id, err := utils.ParseUUIDParam(ctx, "id")
	if err != nil {
		return err
	}

	if err := c.service.DeleteBook(id); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Book deleted successfully",
	})
}
