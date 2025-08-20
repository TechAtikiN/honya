package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/techatikin/backend/dtos"
	"github.com/techatikin/backend/services"
)

type TBookController struct {
	service services.TBookService
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
	AuthorId        string  `json:"author_id"`
}

type BookUpdateRequest struct {
	BookCreateRequest
	ID string `json:"id"`
}

type BookResponse struct {
	ID              string  `json:"id"`
	Title           string  `json:"title"`
	Description     string  `json:"description"`
	Category        string  `json:"category"`
	Image           string  `json:"image"`
	PublicationYear int     `json:"publication_year"`
	Rating          float64 `json:"rating"`
	Pages           int     `json:"pages"`
	Isbn            string  `json:"isbn"`
	AuthorId        string  `json:"author_id"`
	CreatedAt       int64   `json:"created_at"`
	UpdatedAt       int64   `json:"updated_at"`
}

func BookController(service services.TBookService) *TBookController {
	return &TBookController{service}
}

func (c *TBookController) GetBooks(ctx *fiber.Ctx) error {
	books, err := c.service.GetBooks()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(books)
}

func (c *TBookController) GetBookByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	book, err := c.service.GetBookByID(id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
	}

	result := BookResponse{
		ID:              book.ID,
		Title:           book.Title,
		Description:     book.Description,
		Category:        book.Category,
		Image:           book.Image,
		PublicationYear: book.PublicationYear,
		Rating:          book.Rating,
		Pages:           book.Pages,
		Isbn:            book.Isbn,
		AuthorId:        book.AuthorId,
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

	result := BookResponse{
		ID:              book.ID,
		Title:           book.Title,
		Description:     book.Description,
		Category:        book.Category,
		Image:           book.Image,
		PublicationYear: book.PublicationYear,
		Rating:          book.Rating,
		Pages:           book.Pages,
		Isbn:            book.Isbn,
		AuthorId:        book.AuthorId,
		CreatedAt:       book.CreatedAt,
		UpdatedAt:       book.UpdatedAt,
	}

	return ctx.Status(fiber.StatusCreated).JSON(result)
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
