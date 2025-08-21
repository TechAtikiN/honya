package controllers

import (
	"fmt"
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
	books, count, err := c.service.GetBooks(params)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Prepare response
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"count": count,
		"data":  books,
	})
}

func (c *TBookController) GetBookByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	book, err := c.service.GetBookByID(id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
	}

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

	return ctx.Status(fiber.StatusOK).JSON(result)
}

func (c *TBookController) CreateBook(ctx *fiber.Ctx) error {
	var req dtos.BookCreateRequest

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	book, err := c.service.CreateBook(req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

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

	return ctx.Status(fiber.StatusCreated).JSON(result)
}

func (c *TBookController) UpdateBook(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	fmt.Println("Updating book with ID:", id)

	var req dtos.BookUpdateRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	book, err := c.service.UpdateBook(id, req)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

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

	return ctx.Status(fiber.StatusOK).JSON(result)
}

func (c *TBookController) DeleteBook(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	err := c.service.DeleteBook(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Book deleted successfully",
	})
}
