package controller_test

import (
	"bytes"
	"encoding/json"
	"honya/backend/controller"
	"honya/backend/dto"
	"honya/backend/model"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockBookService struct {
	mock.Mock
}

func (m *MockBookService) GetBooks(params dto.BookQueryParams) ([]model.Book, *dto.PaginationMeta, error) {
	args := m.Called(params)
	return args.Get(0).([]model.Book), args.Get(1).(*dto.PaginationMeta), args.Error(2)
}

func (m *MockBookService) GetBookByID(id uuid.UUID) (*model.Book, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Book), args.Error(1)
}

func (m *MockBookService) CreateBook(req *dto.BookCreateRequest, file *multipart.FileHeader) (*model.Book, error) {
	args := m.Called(req, file)
	return args.Get(0).(*model.Book), args.Error(1)
}

func (m *MockBookService) UpdateBook(id uuid.UUID, req *dto.BookUpdateRequest, fileHeader *multipart.FileHeader) (*model.Book, error) {
	args := m.Called(id, req, fileHeader)
	return args.Get(0).(*model.Book), args.Error(1)
}

func (m *MockBookService) DeleteBook(id uuid.UUID) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestGetBooks_WithQueryParams(t *testing.T) {
	app := fiber.New()
	mockService := new(MockBookService)
	ctrl := controller.NewBookController(mockService)

	books := []model.Book{
		{
			ID:    uuid.New(),
			Title: "Book 1",
		},
		{
			ID:    uuid.New(),
			Title: "Book 2",
		},
	}

	meta := &dto.PaginationMeta{
		TotalCount: 2,
		Limit:      10,
		Offset:     0,
	}

	params := dto.BookQueryParams{
		Query:           "book",
		Offset:          0,
		Limit:           10,
		Category:        "fiction",
		PublicationYear: 2025,
		Rating:          0,
		Pages:           0,
		Sort:            "title",
	}

	mockService.On("GetBooks", params).Return(books, meta, nil)

	app.Get("/api/books", ctrl.GetBooks)

	req := httptest.NewRequest(http.MethodGet, "/api/books?query=book&offset=0&limit=10&category=fiction&publication_year=2025&sort=title", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestGetBookByID(t *testing.T) {
	app := fiber.New()
	mockService := new(MockBookService)
	ctrl := controller.NewBookController(mockService)

	bookID := uuid.New()
	book := &model.Book{
		ID:    bookID,
		Title: "Test Book",
	}

	mockService.On("GetBookByID", bookID).Return(book, nil)

	app.Get("/api/books/:id", ctrl.GetBookByID)

	req := httptest.NewRequest(http.MethodGet, "/api/books/"+bookID.String(), nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestCreateBook_MultipartFormData(t *testing.T) {
	app := fiber.New()
	mockService := new(MockBookService)
	ctrl := controller.NewBookController(mockService)

	bookID := uuid.New()
	dummyBook := &model.Book{
		ID:              bookID,
		Title:           "Test Book",
		Description:     "A test book",
		Category:        "fiction",
		PublicationYear: 2025,
		Rating:          4.5,
		Pages:           300,
		Isbn:            "1234567890",
		AuthorName:      "John Doe",
	}

	mockService.On("CreateBook", mock.AnythingOfType("*dto.BookCreateRequest"), mock.AnythingOfType("*multipart.FileHeader")).Return(dummyBook, nil)

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	_ = writer.WriteField("title", "Test Book")
	_ = writer.WriteField("description", "A test book")
	_ = writer.WriteField("category", "fiction")
	_ = writer.WriteField("publication_year", "2025")
	_ = writer.WriteField("rating", "4.5")
	_ = writer.WriteField("pages", "300")
	_ = writer.WriteField("isbn", "1234567890")
	_ = writer.WriteField("author_name", "John Doe")

	fileWriter, _ := writer.CreateFormFile("image", "dummy.jpg")
	if _, err := fileWriter.Write([]byte("dummy file content")); err != nil {
		t.Fatalf("failed to write file content: %v", err)
	}

	writer.Close()

	app.Post("/api/books", ctrl.CreateBook)

	req := httptest.NewRequest(http.MethodPost, "/api/books", &body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}

func TestUpdateBook_JSON(t *testing.T) {
	app := fiber.New()
	mockService := new(MockBookService)
	ctrl := controller.NewBookController(mockService)

	bookID := uuid.New()
	updatedTitle := "Updated Book"
	updateReq := dto.BookUpdateRequest{
		Title: &updatedTitle,
	}

	updatedBook := &model.Book{
		ID:         bookID,
		Title:      updatedTitle,
		AuthorName: "John Doe",
		Category:   "fiction",
	}

	mockService.On("UpdateBook", bookID, &updateReq, (*multipart.FileHeader)(nil)).Return(updatedBook, nil)

	app.Patch("/api/books/:id", ctrl.UpdateBook)

	body, _ := json.Marshal(updateReq)
	req := httptest.NewRequest(http.MethodPatch, "/api/books/"+bookID.String(), bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestUpdateBook_MultipartFormData(t *testing.T) {
	app := fiber.New()
	mockService := new(MockBookService)
	ctrl := controller.NewBookController(mockService)

	bookID := uuid.New()
	updatedTitle := "Updated Book"
	updatedPages := 350

	updatedBook := &model.Book{
		ID:         bookID,
		Title:      updatedTitle,
		Pages:      updatedPages,
		AuthorName: "John Doe",
		Category:   "fiction",
	}

	mockService.On("UpdateBook", bookID, mock.AnythingOfType("*dto.BookUpdateRequest"), mock.AnythingOfType("*multipart.FileHeader")).Return(updatedBook, nil)

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	_ = writer.WriteField("title", updatedTitle)
	_ = writer.WriteField("pages", strconv.Itoa(updatedPages))

	fileWriter, _ := writer.CreateFormFile("image", "dummy.jpg")
	if _, err := fileWriter.Write([]byte("dummy content")); err != nil {
		t.Fatalf("failed to write file content: %v", err)
	}

	writer.Close()

	app.Patch("/api/books/:id", ctrl.UpdateBook)

	req := httptest.NewRequest(http.MethodPatch, "/api/books/"+bookID.String(), &body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestDeleteBook(t *testing.T) {
	app := fiber.New()
	mockService := new(MockBookService)
	ctrl := controller.NewBookController(mockService)

	bookID := uuid.New()

	mockService.On("DeleteBook", bookID).Return(nil)

	app.Delete("/api/books/:id", ctrl.DeleteBook)

	req := httptest.NewRequest(http.MethodDelete, "/api/books/"+bookID.String(), nil)

	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
