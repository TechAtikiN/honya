package controller_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/techatikin/backend/controller"
	"github.com/techatikin/backend/dto"
	"github.com/techatikin/backend/model"
)

type MockReviewService struct {
	mock.Mock
}

func (m *MockReviewService) GetAllReviews(params dto.QueryParams) ([]model.Review, *dto.PaginationMeta, error) {
	args := m.Called(params)
	return args.Get(0).([]model.Review), args.Get(1).(*dto.PaginationMeta), args.Error(2)
}

func (m *MockReviewService) GetReviewByID(id uuid.UUID) (*model.Review, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Review), args.Error(1)
}

func (m *MockReviewService) GetReviewsByBookID(bookID uuid.UUID, params dto.QueryParams) ([]model.Review, *dto.PaginationMeta, error) {
	args := m.Called(bookID, params)
	return args.Get(0).([]model.Review), args.Get(1).(*dto.PaginationMeta), args.Error(2)
}

func (m *MockReviewService) CreateReview(req *dto.ReviewCreateRequest) (*model.Review, error) {
	args := m.Called(req)
	return args.Get(0).(*model.Review), args.Error(1)
}

func (m *MockReviewService) UpdateReview(id uuid.UUID, req *dto.ReviewUpdateRequest) (*model.Review, error) {
	args := m.Called(id, req)
	return args.Get(0).(*model.Review), args.Error(1)
}

func (m *MockReviewService) DeleteReview(id uuid.UUID) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockReviewService) FindByBookID(bookID uuid.UUID, params dto.QueryParams) ([]model.Review, dto.PaginationMeta, error) {
	args := m.Called(bookID, params)
	return args.Get(0).([]model.Review), args.Get(1).(dto.PaginationMeta), args.Error(2)
}

func TestGetAllReviews(t *testing.T) {
	app := fiber.New()
	mockService := new(MockReviewService)
	ctrl := controller.NewReviewController(mockService)

	reviews := []model.Review{
		{ID: uuid.New(), Name: "Alice", Email: "a@test.com", Content: "Great book!"},
		{ID: uuid.New(), Name: "Bob", Email: "b@test.com", Content: "Not bad."},
	}
	meta := &dto.PaginationMeta{TotalCount: 2, Limit: 10, Offset: 0}

	params := dto.QueryParams{Query: "", Offset: 0, Limit: 10}
	mockService.On("GetAllReviews", params).Return(reviews, meta, nil)

	app.Get("/reviews", ctrl.GetAllReviews)
	req := httptest.NewRequest(http.MethodGet, "/reviews", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestGetReviewByID(t *testing.T) {
	app := fiber.New()
	mockService := new(MockReviewService)
	ctrl := controller.NewReviewController(mockService)

	id := uuid.New()
	review := &model.Review{ID: id, Name: "Alice", Email: "a@test.com", Content: "Great book!"}

	mockService.On("GetReviewByID", id).Return(review, nil)

	app.Get("/reviews/:id", ctrl.GetReviewByID)
	req := httptest.NewRequest(http.MethodGet, "/reviews/"+id.String(), nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestGetReviewsByBookID(t *testing.T) {
	app := fiber.New()
	mockService := new(MockReviewService)
	ctrl := controller.NewReviewController(mockService)

	bookID := uuid.New()
	reviews := []model.Review{
		{ID: uuid.New(), BookID: bookID, Name: "Alice", Email: "a@test.com", Content: "Great!"},
	}
	meta := &dto.PaginationMeta{TotalCount: 1, Limit: 10, Offset: 0}

	params := dto.QueryParams{Query: "", Offset: 0, Limit: 10}
	mockService.On("GetReviewsByBookID", bookID, params).Return(reviews, meta, nil)

	app.Get("/books/:book_id/reviews", ctrl.GetReviewsByBookID)
	req := httptest.NewRequest(http.MethodGet, "/books/"+bookID.String()+"/reviews", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestCreateReview(t *testing.T) {
	app := fiber.New()
	mockService := new(MockReviewService)
	ctrl := controller.NewReviewController(mockService)

	bookID := uuid.New()
	reqBody := dto.ReviewCreateRequest{
		BookID:  bookID,
		Name:    "Alice",
		Email:   "a@test.com",
		Content: "Excellent book",
	}
	review := &model.Review{
		ID:        uuid.New(),
		BookID:    bookID,
		Name:      "Alice",
		Email:     "a@test.com",
		Content:   "Excellent book",
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	mockService.On("CreateReview", &reqBody).Return(review, nil)

	app.Post("/reviews", ctrl.CreateReview)
	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/reviews", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}

func TestUpdateReview(t *testing.T) {
	app := fiber.New()
	mockService := new(MockReviewService)
	ctrl := controller.NewReviewController(mockService)

	id := uuid.New()
	name := "Bob"
	updateReq := dto.ReviewUpdateRequest{Name: &name}
	review := &model.Review{
		ID:      id,
		Name:    "Bob",
		Email:   "b@test.com",
		Content: "Updated content",
	}

	mockService.On("UpdateReview", id, &updateReq).Return(review, nil)

	app.Patch("/reviews/:id", ctrl.UpdateReview)
	body, _ := json.Marshal(updateReq)
	req := httptest.NewRequest(http.MethodPatch, "/reviews/"+id.String(), bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestDeleteReview(t *testing.T) {
	app := fiber.New()
	mockService := new(MockReviewService)
	ctrl := controller.NewReviewController(mockService)

	id := uuid.New()
	mockService.On("DeleteReview", id).Return(nil)

	app.Delete("/reviews/:id", ctrl.DeleteReview)
	req := httptest.NewRequest(http.MethodDelete, "/reviews/"+id.String(), nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
