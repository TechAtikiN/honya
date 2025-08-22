package controllers

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/techatikin/backend/dtos"
	"github.com/techatikin/backend/errors"
	"github.com/techatikin/backend/services"
	"github.com/techatikin/backend/utils"
)

type TBookController struct {
	service services.TBookService
}

func BookController(service services.TBookService) *TBookController {
	return &TBookController{service}
}

func (c *TBookController) GetBooks(ctx *fiber.Ctx) error {
	params := dtos.BookQueryParams{
		Query:           ctx.Query("query"),
		Offset:          utils.ParseInt(ctx.Query("offset"), 0),
		Limit:           utils.ParseInt(ctx.Query("limit"), 10),
		Category:        strings.ToLower(ctx.Query("category")),
		PublicationYear: utils.ParseInt(ctx.Query("publication_year"), 2025),
		Rating:          utils.ParseFloat(ctx.Query("rating"), 0),
		Pages:           utils.ParseInt(ctx.Query("pages"), 0),
		Sort:            strings.ToLower(ctx.Query("sort")),
	}

	books, meta, err := c.service.GetBooks(params)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"meta": meta,
		"data": books,
	})
}

func (c *TBookController) GetBookByID(ctx *fiber.Ctx) error {
	id, err := utils.ParseUUIDParam(ctx, "id")
	if err != nil {
		return err
	}

	book, err := c.service.GetBookByID(id)
	if err != nil {
		return err
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
	var reqData dtos.BookCreateRequest
	if err := ctx.BodyParser(&reqData); err != nil {
		return errors.NewBadRequestError("Invalid JSON body")
	}

	book, err := c.service.CreateBook(&reqData)
	if err != nil {
		return err
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

	updatedBook, err := c.service.UpdateBook(id, &requestData)
	if err != nil {
		return err
	}

	result := dtos.BookResponse{
		ID:              updatedBook.ID,
		Title:           updatedBook.Title,
		Description:     updatedBook.Description,
		Category:        updatedBook.Category,
		Image:           updatedBook.Image,
		PublicationYear: updatedBook.PublicationYear,
		Rating:          updatedBook.Rating,
		Pages:           updatedBook.Pages,
		Isbn:            updatedBook.Isbn,
		AuthorName:      updatedBook.AuthorName,
		CreatedAt:       updatedBook.CreatedAt,
		UpdatedAt:       updatedBook.UpdatedAt,
	}

	return ctx.Status(fiber.StatusOK).JSON(result)
}

func (c *TBookController) DeleteBook(ctx *fiber.Ctx) error {
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
