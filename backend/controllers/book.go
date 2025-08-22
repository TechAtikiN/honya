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

	result := dtos.ToBookListResponse(books, *meta)
	return ctx.Status(fiber.StatusOK).JSON(result)
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

	result := dtos.ToBookResponse(book)
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

	result := dtos.ToBookResponse(book)
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

	result := dtos.ToBookResponse(updatedBook)
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
