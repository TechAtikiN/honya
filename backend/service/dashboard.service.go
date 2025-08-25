package service

import (
	"fmt"

	"github.com/techatikin/backend/dto"
	"github.com/techatikin/backend/errors"
	"github.com/techatikin/backend/repository"
)

type DashboardService interface {
	GetDonutChartData(filterBy string) (*dto.DonutChartData, error)
	GetBarChartData(limit int) (*dto.BarChartData, error)
}

type dashboardService struct {
	bookRepo   repository.BookRepository
	reviewRepo repository.ReviewRepository
}

func NewDashboardService(bookRepo repository.BookRepository, reviewRepo repository.ReviewRepository) DashboardService {
	return &dashboardService{
		bookRepo:   bookRepo,
		reviewRepo: reviewRepo,
	}
}

func (s *dashboardService) GetDonutChartData(filterBy string) (*dto.DonutChartData, error) {
	var field string
	switch filterBy {
	case "category":
		field = "category"
	case "rating":
		field = "rating"
	case "author":
		field = "author_name"
	default:
		return nil, errors.NewBadRequestError(fmt.Sprintf("Invalid filter_by value '%s'. Allowed values: category, rating, author", filterBy))
	}

	data, err := s.bookRepo.CountByField(field)
	if err != nil {
		return nil, errors.NewInternalError(fmt.Errorf("failed to get donut chart data: %w", err))
	}

	return &dto.DonutChartData{
		FilterBy: filterBy,
		Data:     data,
	}, nil
}

func (s *dashboardService) GetBarChartData(limit int) (*dto.BarChartData, error) {
	if limit < 0 {
		return nil, errors.NewBadRequestError("Invalid limit value")
	}

	if limit == 0 {
		limit = 10
	}

	reviewerStats, err := s.reviewRepo.GetTopReviewers(limit)
	if err != nil {
		return nil, errors.NewInternalError(fmt.Errorf("failed to get top reviewers: %w", err))
	}

	return &dto.BarChartData{
		Data: reviewerStats,
	}, nil
}
