package controller

import (
	"honya/backend/service"
	"honya/backend/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type DashboardController interface {
	GetBooksData(ctx *fiber.Ctx) error
	GetReviewsData(ctx *fiber.Ctx) error
}

type dashboardController struct {
	service service.DashboardService
}

func NewDashboardController(service service.DashboardService) DashboardController {
	return &dashboardController{service}
}

// GetBooksData godoc
// @Summary Get books data
// @Description Get books data
// @Tags dashboard
// @Accept json
// @Produce json
// @Param filter_by query string false "Filter by"
// @Success 200 {object} model.Book
// @Router /dashboard/books [get]
func (c *dashboardController) GetBooksData(ctx *fiber.Ctx) error {
	// GetBooksData godoc
	// @Summary Get books data
	// @Description Get books data
	// @Tags dashboard
	// @Accept json
	// @Produce json
	// @Param filter_by query string false "Filter by"
	// @Success 200 {object} model.Book
	// @Router /dashboard/books [get]
	filterBy := ctx.Query("filter_by", utils.DefaultDonutChartFilterBy)

	data, err := c.service.GetBooksData(filterBy)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Failed to retrieve dashboard data",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(data)
}

func (c *dashboardController) GetReviewsData(ctx *fiber.Ctx) error {
	// GetReviewsData godoc
	// @Summary Get reviews data
	// @Description Get reviews data
	// @Tags dashboard
	// @Accept json
	// @Produce json
	// @Param limit query string false "Limit"
	// @Success 200 {object} model.Review
	limitStr := ctx.Query("limit", "10")
	limit, err := strconv.Atoi(limitStr)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid limit parameter",
			"message": "Limit must be a valid integer",
		})
	}

	data, err := c.service.GetReviewsData(limit)
	if err != nil {

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Failed to retrieve top reviewers data",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(data)
}
