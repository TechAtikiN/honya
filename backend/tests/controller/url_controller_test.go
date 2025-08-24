package controller_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/techatikin/backend/controller"
	"github.com/techatikin/backend/dto"
	"github.com/techatikin/backend/middleware"
)

type MockUrlService struct {
	mock.Mock
}

func (m *MockUrlService) GetRedirectionUrl(url string) (string, error) {
	args := m.Called(url)
	return args.String(0), args.Error(1)
}

func (m *MockUrlService) GetCanonicalUrl(url string) (string, error) {
	args := m.Called(url)
	return args.String(0), args.Error(1)
}

func TestProcessUrl_Redirection(t *testing.T) {
	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
	})
	mockService := new(MockUrlService)
	ctrl := controller.NewUrlController(mockService)

	mockService.On("GetRedirectionUrl", "https://example.com/test").Return("https://example.com/test-redirect", nil)

	app.Post("/process-url", ctrl.ProcessUrl)

	reqBody := dto.ProcessUrlRequest{
		Url:       "https://example.com/test",
		Operation: "redirection",
	}

	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/process-url", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestProcessUrl_Canonical(t *testing.T) {
	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
	})
	mockService := new(MockUrlService)
	ctrl := controller.NewUrlController(mockService)

	mockService.On("GetCanonicalUrl", "https://example.com/test").Return("https://example.com/test-canonical", nil)

	app.Post("/process-url", ctrl.ProcessUrl)

	reqBody := dto.ProcessUrlRequest{
		Url:       "https://example.com/test",
		Operation: "canonical",
	}

	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/process-url", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestProcessUrl_All(t *testing.T) {
	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
	})
	mockService := new(MockUrlService)
	ctrl := controller.NewUrlController(mockService)

	mockService.On("GetRedirectionUrl", "https://example.com/test").Return("https://example.com/redirected", nil)
	mockService.On("GetCanonicalUrl", "https://example.com/redirected").Return("https://example.com/canonical", nil)

	app.Post("/process-url", ctrl.ProcessUrl)

	reqBody := dto.ProcessUrlRequest{
		Url:       "https://example.com/test",
		Operation: "all",
	}

	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/process-url", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestProcessUrl_InvalidURL(t *testing.T) {
	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
	})
	mockService := new(MockUrlService)
	ctrl := controller.NewUrlController(mockService)

	app.Post("/process-url", ctrl.ProcessUrl)

	reqBody := dto.ProcessUrlRequest{
		Url:       "invalid-url",
		Operation: "redirection",
	}

	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/process-url", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestProcessUrl_InvalidOperation(t *testing.T) {
	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
	})
	mockService := new(MockUrlService)
	ctrl := controller.NewUrlController(mockService)

	app.Post("/process-url", ctrl.ProcessUrl)

	reqBody := dto.ProcessUrlRequest{
		Url:       "https://example.com",
		Operation: "invalid-op",
	}

	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/process-url", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}
