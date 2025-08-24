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
// @Summary Get a list of books
// @Description Retrieve a list of books with optional filtering, sorting, and pagination
// @Tags books
// @Accept json
// @Produce json
// @Param query query string false "Search query"
// @Param offset query int false "Pagination offset" default(0)
// @Param limit query int false "Pagination limit" default(10)
// @Param category query string false "Book category (Available categories: fiction, non_fiction, science, history, fantasy, mystery, thriller, cooking, travel, classics)"
// @Param publication_year query int false "Publication year"
// @Param rating query number false "Minimum rating"
// @Param pages query int false "Number of pages"
// @Param sort query string false "Sort by field (e.g., title, publication_year, rating)"
// @Success 200 {object} dto.BookListResponse "List of books fetched successfully"
// @Failure 400 {object} errors.ErrorResponse "Invalid query parameters"
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
// @Summary Get a book by ID
// @Description Retrieve detailed information about a specific book by its ID
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
// @Accept multipart/form-data
// @Produce json
// @Param title formData string true "Book title"
// @Param description formData string false "Book description"
// @Param category formData string true "Book category (Available categories: fiction, non_fiction, science, history, fantasy, mystery, thriller, cooking, travel, classics)"
// @Param publication_year formData int true "Publication year"
// @Param rating formData number true "Book rating"
// @Param pages formData int true "Number of pages"
// @Param isbn formData string true "ISBN number"
// @Param author_name formData string true "Author name"
// @Param image formData file false "Book cover image"
// @Success 201 {object} dto.BookResponse "Book created successfully"
// @Failure 400 {object} errors.ErrorResponse "Invalid input data"
// @Router /books [post]
func (c *bookController) CreateBook(ctx *fiber.Ctx) error {
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

	book, err := c.service.CreateBook(&reqData, fileHeader)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") && strings.Contains(err.Error(), "uni_books_isbn") {
			return errors.NewConflictError("A book with this ISBN already exists")
		}
		return err
	}

	result := dto.ToBookResponse(book)

	return ctx.Status(fiber.StatusCreated).JSON(result)
}

// UpdateBook godoc
// @Summary Update an existing book
// @Description Update the details of an existing book by its ID
// @Tags books
// @Accept multipart/form-data
// @Produce json
// @Param id path string true "Book ID"
// @Param title formData string false "Book title"
// @Param description formData string false "Book description"
// @Param category formData string false "Book category (Available categories: fiction, non_fiction, science, history, fantasy, mystery, thriller, cooking, travel, classics)"
// @Param publication_year formData int false "Publication year"
// @Param rating formData number false "Book rating"
// @Param pages formData int false "Number of pages"
// @Param author_name formData string false "Author name"
// @Param image formData file false "Book cover image"
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

		file, err := ctx.FormFile("image")
		if err == nil {
			fileHeader = file
		}
	} else {
		if err := ctx.BodyParser(&requestData); err != nil {
			return errors.NewBadRequestError("Invalid JSON body")
		}
	}

	if requestData.Isbn != nil {
		return errors.NewBadRequestError("ISBN cannot be updated once set")
	}

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
