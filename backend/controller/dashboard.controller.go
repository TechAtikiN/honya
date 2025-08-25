package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/techatikin/backend/service"
	"github.com/techatikin/backend/utils"
)

type DashboardController interface {
	GetDonutChart(ctx *fiber.Ctx) error
	GetBarChart(ctx *fiber.Ctx) error
}

type dashboardController struct {
	service service.DashboardService
}

func NewDashboardController(service service.DashboardService) DashboardController {
	return &dashboardController{service}
}

func (c *dashboardController) GetDonutChart(ctx *fiber.Ctx) error {
	filterBy := ctx.Query("filter_by", utils.DefaultDonutChartFilterBy)

	data, err := c.service.GetDonutChartData(filterBy)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Failed to retrieve dashboard data",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(data)
}

func (c *dashboardController) GetBarChart(ctx *fiber.Ctx) error {
	limitStr := ctx.Query("limit", "10")
	limit, err := strconv.Atoi(limitStr)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid limit parameter",
			"message": "Limit must be a valid integer",
		})
	}

	data, err := c.service.GetBarChartData(limit)
	if err != nil {

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Failed to retrieve top reviewers data",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(data)
}
