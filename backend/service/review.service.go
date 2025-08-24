package service

import (
	"github.com/google/uuid"
	"github.com/techatikin/backend/dto"
	"github.com/techatikin/backend/errors"
	"github.com/techatikin/backend/model"
	"github.com/techatikin/backend/repository"
	"github.com/techatikin/backend/utils"
)

// ReviewService defines service-level operations for reviews
type ReviewService interface {
	GetAllReviews(params dto.QueryParams) ([]model.Review, *dto.PaginationMeta, error)
	GetReviewByID(id uuid.UUID) (*model.Review, error)
	CreateReview(req *dto.ReviewCreateRequest) (*model.Review, error)
	UpdateReview(id uuid.UUID, req *dto.ReviewUpdateRequest) (*model.Review, error)
	DeleteReview(id uuid.UUID) error
	FindByBookID(bookID uuid.UUID, params dto.QueryParams) ([]model.Review, dto.PaginationMeta, error)
	GetReviewsByBookID(bookID uuid.UUID, params dto.QueryParams) ([]model.Review, *dto.PaginationMeta, error)
}

type reviewService struct {
	repo repository.ReviewRepository
}

func NewReviewService(repo repository.ReviewRepository) ReviewService {
	return &reviewService{repo: repo}
}

func (s *reviewService) FindByBookID(bookID uuid.UUID, params dto.QueryParams) ([]model.Review, dto.PaginationMeta, error) {
	reviews, meta, err := s.repo.FindByBookID(bookID, params)
	if err != nil {
		return nil, dto.PaginationMeta{}, errors.NewInternalError(err)
	}
	return reviews, meta, nil
}

func (s *reviewService) GetAllReviews(params dto.QueryParams) ([]model.Review, *dto.PaginationMeta, error) {
	reviews, meta, err := s.repo.FindAll(params)
	if err != nil {
		return nil, nil, errors.NewInternalError(err)
	}
	return reviews, &meta, nil
}

func (s *reviewService) GetReviewByID(id uuid.UUID) (*model.Review, error) {
	review, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.NewInternalError(err)
	}
	if review == nil {
		return nil, errors.NewNotFoundError("Review not found")
	}
	return review, nil
}

func (s *reviewService) CreateReview(req *dto.ReviewCreateRequest) (*model.Review, error) {
	if err := utils.ValidateReviewCreateRequest(req); err != nil {
		return nil, errors.NewBadRequestError(err.Error())
	}

	review := model.Review{
		BookID:  req.BookID,
		Name:    req.Name,
		Email:   req.Email,
		Content: req.Content,
	}

	resource, err := s.repo.Create(&review)
	if err != nil {
		return nil, errors.NewInternalError(err)
	}
	return resource, nil
}

func (s *reviewService) UpdateReview(id uuid.UUID, req *dto.ReviewUpdateRequest) (*model.Review, error) {
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.NewInternalError(err)
	}
	if existing == nil {
		return nil, errors.NewNotFoundError("Review not found")
	}

	updates := map[string]interface{}{}
	if req.Content != nil {
		updates["content"] = *req.Content
	}
	if req.Name != nil {
		updates["name"] = *req.Name
	}
	if req.Email != nil {
		updates["email"] = *req.Email
	}

	updated, err := s.repo.Update(id, updates)
	if err != nil {
		return nil, errors.NewInternalError(err)
	}

	return updated, nil
}

func (s *reviewService) DeleteReview(id uuid.UUID) error {
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return errors.NewInternalError(err)
	}
	if existing == nil {
		return errors.NewNotFoundError("Review not found")
	}

	return s.repo.Delete(id)
}

func (s *reviewService) GetReviewsByBookID(bookID uuid.UUID, params dto.QueryParams) ([]model.Review, *dto.PaginationMeta, error) {
	if bookID == uuid.Nil {
		return nil, nil, errors.NewBadRequestError("Invalid book ID")
	}

	reviews, meta, err := s.repo.FindByBookID(bookID, params)
	if err != nil {
		return nil, nil, errors.NewInternalError(err)
	}

	return reviews, &meta, nil
}
