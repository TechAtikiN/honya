package service_test

import (
	"honya/backend/dto"
	"honya/backend/model"
	"honya/backend/service"
	"testing"

	"honya/backend/errors"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockReviewRepo struct {
	mock.Mock
}

func (m *MockReviewRepo) FindAll(params dto.QueryParams) ([]model.Review, dto.PaginationMeta, error) {
	args := m.Called(params)
	return args.Get(0).([]model.Review), args.Get(1).(dto.PaginationMeta), args.Error(2)
}

func (m *MockReviewRepo) FindByID(id uuid.UUID) (*model.Review, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Review), args.Error(1)
}

func (m *MockReviewRepo) Create(review *model.Review) (*model.Review, error) {
	args := m.Called(review)
	return args.Get(0).(*model.Review), args.Error(1)
}

func (m *MockReviewRepo) Update(id uuid.UUID, updates map[string]interface{}) (*model.Review, error) {
	args := m.Called(id, updates)
	return args.Get(0).(*model.Review), args.Error(1)
}

func (m *MockReviewRepo) Delete(id uuid.UUID) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockReviewRepo) FindByBookID(bookID uuid.UUID, params dto.QueryParams) ([]model.Review, dto.PaginationMeta, error) {
	args := m.Called(bookID, params)
	return args.Get(0).([]model.Review), args.Get(1).(dto.PaginationMeta), args.Error(2)
}

func (m *MockReviewRepo) GetTopReviewers(limit int) ([]dto.ReviewerStats, error) {
	args := m.Called(limit)
	return args.Get(0).([]dto.ReviewerStats), args.Error(1)
}

func TestReviewService_CreateReview(t *testing.T) {
	mockRepo := new(MockReviewRepo)
	svc := service.NewReviewService(mockRepo)

	bookID := uuid.New()
	req := &dto.ReviewCreateRequest{
		BookID:  bookID,
		Name:    "John",
		Email:   "john@example.com",
		Content: "Great book!",
	}

	reviewModel := &model.Review{
		ID:      uuid.New(),
		BookID:  bookID,
		Name:    "John",
		Email:   "john@example.com",
		Content: "Great book!",
	}

	mockRepo.On("Create", mock.AnythingOfType("*model.Review")).Return(reviewModel, nil)

	result, err := svc.CreateReview(req)
	assert.NoError(t, err)
	assert.Equal(t, req.Name, result.Name)
	assert.Equal(t, req.Content, result.Content)

	mockRepo.AssertExpectations(t)
}

func TestReviewService_GetReviewByID_NotFound(t *testing.T) {
	mockRepo := new(MockReviewRepo)
	svc := service.NewReviewService(mockRepo)

	id := uuid.New()
	mockRepo.On("FindByID", id).Return((*model.Review)(nil), nil)

	review, err := svc.GetReviewByID(id)
	assert.Nil(t, review)
	assert.IsType(t, &errors.AppError{}, err)
	assert.Equal(t, 404, err.(*errors.AppError).Code)
	assert.Equal(t, "Review not found", err.(*errors.AppError).Message)

	mockRepo.AssertExpectations(t)
}

func TestReviewService_FindByBookID(t *testing.T) {
	mockRepo := new(MockReviewRepo)
	svc := service.NewReviewService(mockRepo)

	bookID := uuid.New()
	params := dto.QueryParams{Limit: 10, Offset: 0}

	reviews := []model.Review{
		{ID: uuid.New(), BookID: bookID, Name: "John", Email: "john@example.com", Content: "Great!"},
		{ID: uuid.New(), BookID: bookID, Name: "Bob", Email: "bob@example.com", Content: "Loved it!"},
	}
	meta := dto.PaginationMeta{TotalCount: 2, Limit: 10, Offset: 0}

	mockRepo.On("FindByBookID", bookID, params).Return(reviews, meta, nil)

	result, resultMeta, err := svc.FindByBookID(bookID, params)
	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, int64(2), resultMeta.TotalCount)

	mockRepo.AssertExpectations(t)
}

func TestReviewService_DeleteReview_NotFound(t *testing.T) {
	mockRepo := new(MockReviewRepo)
	svc := service.NewReviewService(mockRepo)

	id := uuid.New()
	mockRepo.On("FindByID", id).Return((*model.Review)(nil), nil)

	err := svc.DeleteReview(id)
	assert.NotNil(t, err)
	assert.Equal(t, 404, err.(*errors.AppError).Code)

	mockRepo.AssertExpectations(t)
}
