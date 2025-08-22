package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/techatikin/backend/dto"
	"github.com/techatikin/backend/errors"
	"github.com/techatikin/backend/service"
	"github.com/techatikin/backend/utils"
)

type ReviewController interface {
	GetAllReviews(ctx *fiber.Ctx) error
	GetReviewByID(ctx *fiber.Ctx) error
	GetReviewsByBookID(ctx *fiber.Ctx) error
	CreateReview(ctx *fiber.Ctx) error
	UpdateReview(ctx *fiber.Ctx) error
	DeleteReview(ctx *fiber.Ctx) error
}

type reviewController struct {
	service service.ReviewService
}

func NewReviewController(service service.ReviewService) ReviewController {
	return &reviewController{service}
}

// GetAllReviews godoc
// @Summary Get list of all reviews
// @Description Get paginated list of reviews with optional search query
// @Tags reviews
// @Accept json
// @Produce json
// @Param query query string false "Search query"
// @Param offset query integer false "Offset for pagination" default(0)
// @Param limit query integer false "Limit for pagination" default(10)
// @Success 200 {object} dto.ReviewListResponse "Reviews fetched successfully"
// @Router /reviews [get]
func (c *reviewController) GetAllReviews(ctx *fiber.Ctx) error {
	params := dto.QueryParams{
		Query:  ctx.Query("query"),
		Offset: utils.ParseInt(ctx.Query("offset"), utils.DefaultOffset),
		Limit:  utils.ParseInt(ctx.Query("limit"), utils.DefaultLimit),
	}

	reviews, meta, err := c.service.GetAllReviews(params)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.ToReviewListResponse(reviews, *meta))
}

// GetReviewByID godoc
// @Summary Get a review by ID
// @Description Get a single review by its ID
// @Tags reviews
// @Accept json
// @Produce json
// @Param id path string true "Review ID"
// @Success 200 {object} dto.ReviewResponse "Review fetched successfully"
// @Failure 400 {object} errors.ErrorResponse "Invalid ID format"
// @Failure 404 {object} errors.ErrorResponse "Review not found"
// @Router /reviews/{id} [get]
func (c *reviewController) GetReviewByID(ctx *fiber.Ctx) error {
	id, err := utils.ParseUUIDParam(ctx, "id")
	if err != nil {
		return err
	}

	review, err := c.service.GetReviewByID(id)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.ToReviewResponse(review))
}

// GetReviewsByBookID godoc
// @Summary Get reviews for a specific book
// @Description Get paginated list of reviews for a given book ID with optional search query
// @Tags reviews
// @Accept json
// @Produce json
// @Param book_id path string true "Book ID"
// @Param query query string false "Search query"
// @Param offset query integer false "Offset for pagination" default(0)
// @Param limit query integer false "Limit for pagination" default(10)
// @Success 200 {object} dto.ReviewListResponse "Reviews fetched successfully"
// @Failure 400 {object} errors.ErrorResponse "Invalid ID format"
// @Failure 404 {object} errors.ErrorResponse "Book not found"
// @Router /books/{book_id}/reviews [get]
func (c *reviewController) GetReviewsByBookID(ctx *fiber.Ctx) error {
	bookID, err := utils.ParseUUIDParam(ctx, "book_id")
	if err != nil {
		return err
	}

	params := dto.QueryParams{
		Offset: utils.ParseInt(ctx.Query("offset"), utils.DefaultOffset),
		Limit:  utils.ParseInt(ctx.Query("limit"), utils.DefaultLimit),
		Query:  ctx.Query("query"),
	}

	reviews, meta, err := c.service.GetReviewsByBookID(bookID, params)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.ToReviewListResponse(reviews, *meta))
}

// CreateReview godoc
// @Summary Create a new review
// @Description Create a new review for a book
// @Tags reviews
// @Accept json
// @Produce json
// @Param review body dto.ReviewCreateRequest true "Review creation payload"
// @Success 201 {object} dto.ReviewResponse "Review created successfully"
// @Failure 400 {object} errors.ErrorResponse "Invalid input data"
// @Failure 404 {object} errors.ErrorResponse "Book not found"
// @Router /reviews [post]
func (c *reviewController) CreateReview(ctx *fiber.Ctx) error {
	var req dto.ReviewCreateRequest
	if err := ctx.BodyParser(&req); err != nil {
		return errors.NewBadRequestError("Invalid JSON body")
	}

	review, err := c.service.CreateReview(&req)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(dto.ToReviewResponse(review))
}

// UpdateReview godoc
// @Summary Update an existing review
// @Description Update a review by its ID
// @Tags reviews
// @Accept json
// @Produce json
// @Param id path string true "Review ID"
// @Param review body dto.ReviewUpdateRequest true "Review update payload"
// @Success 200 {object} dto.ReviewResponse "Review updated successfully"
// @Failure 400 {object} errors.ErrorResponse "Invalid input data"
// @Failure 404 {object} errors.ErrorResponse "Review not found"
// @Router /reviews/{id} [put]
func (c *reviewController) UpdateReview(ctx *fiber.Ctx) error {
	id, err := utils.ParseUUIDParam(ctx, "id")
	if err != nil {
		return err
	}

	var req dto.ReviewUpdateRequest
	if err := ctx.BodyParser(&req); err != nil {
		return errors.NewBadRequestError("Invalid JSON body")
	}

	updated, err := c.service.UpdateReview(id, &req)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.ToReviewResponse(updated))
}

// DeleteReview godoc
// @Summary Delete a review
// @Description Delete a review by its ID
// @Tags reviews
// @Accept json
// @Produce json
// @Param id path string true "Review ID"
// @Success 200 {object} map[string]string "Review deleted successfully"
// @Failure 400 {object} errors.ErrorResponse "Invalid ID format"
// @Failure 404 {object} errors.ErrorResponse "Review not found"
// @Router /reviews/{id} [delete]
func (c *reviewController) DeleteReview(ctx *fiber.Ctx) error {
	id, err := utils.ParseUUIDParam(ctx, "id")
	if err != nil {
		return err
	}

	if err := c.service.DeleteReview(id); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Review deleted successfully",
	})
}
