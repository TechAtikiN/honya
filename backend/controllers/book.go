package controllers

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/techatikin/backend/dtos"
	"github.com/techatikin/backend/services"
	"github.com/techatikin/backend/utils"
)

type TBookController struct {
	service services.TBookService
}

func BookController(service services.TBookService) *TBookController {
	return &TBookController{service}
}

type BookCreateRequest struct {
	Title           string  `json:"title"`
	Description     string  `json:"description"`
	Category        string  `json:"category"`
	Image           string  `json:"image"`
	PublicationYear int     `json:"publication_year"`
	Rating          float64 `json:"rating"`
	Pages           int     `json:"pages"`
	Isbn            string  `json:"isbn"`
	AuthorName      string  `json:"author_name"`
}

type BookUpdateRequest struct {
	BookCreateRequest
	ID string `json:"id"`
}

func (c *TBookController) GetBooks(ctx *fiber.Ctx) error {
	// Get params
	params := dtos.BookQueryParams{
		Query:           ctx.Query("query"),
		Offset:          utils.ParseInt(ctx.Query("offset"), 0),
		Limit:           utils.ParseInt(ctx.Query("limit"), 10),
		Category:        strings.ToLower(ctx.Query("category")),
		PublicationYear: utils.ParseInt(ctx.Query("publication_year"), 2025),
		Rating:          utils.ParseFloat(ctx.Query("rating"), 5),
		Pages:           utils.ParseInt(ctx.Query("pages"), 0),
	}

	// Call service to get books
	books, meta, err := c.service.GetBooks(params)
	if err != nil {
		return err
	}

	// Prepare response
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"meta": meta,
		"data": books,
	})
}

func (c *TBookController) GetBookByID(ctx *fiber.Ctx) error {
	// Get book ID from params
	id := ctx.Params("id")

	// Call service to get book by ID
	book, err := c.service.GetBookByID(id)
	if err != nil {
		return err
	}

	// Prepare response
	result := dtos.BookResponse{
		ID:              book.ID,
		Title:           book.Title,
		Description:     book.Description,
		Category:        book.Category,
		Image:           book.Image,
		PublicationYear: book.PublicationYear,
		Rating:          book.Rating,
		Pages:           book.Pages,
		Isbn:            book.Isbn,
		AuthorName:      book.AuthorName,
		CreatedAt:       book.CreatedAt,
		UpdatedAt:       book.UpdatedAt,
	}

	// Return book details
	return ctx.Status(fiber.StatusOK).JSON(result)
}

func (c *TBookController) CreateBook(ctx *fiber.Ctx) error {
	// Request body
	var reqData dtos.BookCreateRequest

	// Parse request body
	if err := ctx.BodyParser(&reqData); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Create book using service
	book, err := c.service.CreateBook(&reqData)
	if err != nil {
		return err
	}

	// Prepare response
	result := dtos.BookResponse{
		ID:              book.ID,
		Title:           book.Title,
		Description:     book.Description,
		Category:        book.Category,
		Image:           book.Image,
		PublicationYear: book.PublicationYear,
		Rating:          book.Rating,
		Pages:           book.Pages,
		Isbn:            book.Isbn,
		AuthorName:      book.AuthorName,
		CreatedAt:       book.CreatedAt,
		UpdatedAt:       book.UpdatedAt,
	}

	// Return created book
	return ctx.Status(fiber.StatusCreated).JSON(result)
}

func (c *TBookController) UpdateBook(ctx *fiber.Ctx) error {
	// Get book ID from params
	id := ctx.Params("id")

	// Request body
	var requestData dtos.BookUpdateRequest

	// Parse request body
	if err := ctx.BodyParser(&requestData); err != nil {
		return err
	}

	// Update book using service
	updatedTodo, err := c.service.UpdateBook(id, &requestData)
	if err != nil {
		return err
	}

	// Prepare response
	result := dtos.BookResponse{
		ID:              updatedTodo.ID,
		Title:           updatedTodo.Title,
		Description:     updatedTodo.Description,
		Category:        updatedTodo.Category,
		Image:           updatedTodo.Image,
		PublicationYear: updatedTodo.PublicationYear,
		Rating:          updatedTodo.Rating,
		Pages:           updatedTodo.Pages,
		Isbn:            updatedTodo.Isbn,
		AuthorName:      updatedTodo.AuthorName,
		CreatedAt:       updatedTodo.CreatedAt,
		UpdatedAt:       updatedTodo.UpdatedAt,
	}

	// Return updated book
	return ctx.Status(fiber.StatusOK).JSON(result)
}

func (c *TBookController) DeleteBook(ctx *fiber.Ctx) error {
	// Get book ID from params
	id := ctx.Params("id")

	// Call service to delete book
	err := c.service.DeleteBook(id)
	if err != nil {
		return err
	}

	// Return success message
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Book deleted successfully",
	})
}
