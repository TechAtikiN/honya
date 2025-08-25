package service_test

import (
	"mime/multipart"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/techatikin/backend/dto"
	"github.com/techatikin/backend/errors"
	"github.com/techatikin/backend/model"
	"github.com/techatikin/backend/service"
)

type MockBookRepo struct {
	mock.Mock
}

func (m *MockBookRepo) FindAll(params dto.BookQueryParams) ([]model.Book, dto.PaginationMeta, error) {
	args := m.Called(params)
	return args.Get(0).([]model.Book), args.Get(1).(dto.PaginationMeta), args.Error(2)
}

func (m *MockBookRepo) FindByID(id uuid.UUID) (*model.Book, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Book), args.Error(1)
}

func (m *MockBookRepo) Create(book *model.Book) (*model.Book, error) {
	args := m.Called(book)
	return args.Get(0).(*model.Book), args.Error(1)
}

func (m *MockBookRepo) Update(id uuid.UUID, updateData *dto.BookUpdateRequest) (*model.Book, error) {
	args := m.Called(id, updateData)
	return args.Get(0).(*model.Book), args.Error(1)
}

func (m *MockBookRepo) Delete(id uuid.UUID) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockBookRepo) CountByField(field string) (map[string]int64, error) {
	args := m.Called(field)
	return args.Get(0).(map[string]int64), args.Error(1)
}

type MockS3Repo struct {
	mock.Mock
}

// UPDATED: UploadImage now requires fileHeader and bookName only
func (m *MockS3Repo) UploadImage(fileHeader *multipart.FileHeader, bookName string) (string, error) {
	args := m.Called(fileHeader, bookName)
	return args.String(0), args.Error(1)
}

func (m *MockS3Repo) DeleteImage(key string) error {
	args := m.Called(key)
	return args.Error(0)
}

func TestBookService_GetBooks(t *testing.T) {
	mockRepo := new(MockBookRepo)
	mockS3 := new(MockS3Repo)

	svc := service.NewBookService(mockRepo, mockS3)

	params := dto.BookQueryParams{Limit: 10, Offset: 0}

	books := []model.Book{
		{Title: "Book A"},
		{Title: "Book B"},
	}
	meta := dto.PaginationMeta{TotalCount: 2, Limit: 10, Offset: 0}

	mockRepo.On("FindAll", params).Return(books, meta, nil)

	result, resultMeta, err := svc.GetBooks(params)
	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, int64(2), resultMeta.TotalCount)

	mockRepo.AssertExpectations(t)
}

func TestBookService_GetBookByID_NotFound(t *testing.T) {
	mockRepo := new(MockBookRepo)
	mockS3 := new(MockS3Repo)
	svc := service.NewBookService(mockRepo, mockS3)

	bookID := uuid.New()
	mockRepo.On("FindByID", bookID).Return((*model.Book)(nil), nil)

	book, err := svc.GetBookByID(bookID)
	assert.Nil(t, book)
	assert.IsType(t, &errors.AppError{}, err)
	assert.Equal(t, 404, err.(*errors.AppError).Code)
	assert.Equal(t, "Book not found", err.(*errors.AppError).Message)
	mockRepo.AssertExpectations(t)
}

func TestBookService_CreateBook_WithImage(t *testing.T) {
	mockRepo := new(MockBookRepo)
	mockS3 := new(MockS3Repo)
	svc := service.NewBookService(mockRepo, mockS3)

	req := &dto.BookCreateRequest{
		Title:           "Book A",
		AuthorName:      "Author",
		Category:        "fiction",
		PublicationYear: 2000,
		Rating:          4,
		Pages:           300,
		Isbn:            "12345",
	}

	fileHeader := &multipart.FileHeader{}

	// Expect UploadImage with fileHeader and book title only
	mockS3.On("UploadImage", fileHeader, req.Title).
		Return("s3://bucket/book.png", nil)

	mockRepo.On("Create", mock.AnythingOfType("*model.Book")).
		Return(&model.Book{Title: req.Title, Image: "s3://bucket/book.png"}, nil)

	book, err := svc.CreateBook(req, fileHeader)
	assert.NoError(t, err)
	assert.Equal(t, "Book A", book.Title)
	assert.Equal(t, "s3://bucket/book.png", book.Image)

	mockS3.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestBookService_DeleteBook_NotFound(t *testing.T) {
	mockRepo := new(MockBookRepo)
	mockS3 := new(MockS3Repo)
	svc := service.NewBookService(mockRepo, mockS3)

	bookID := uuid.New()
	mockRepo.On("FindByID", bookID).Return((*model.Book)(nil), nil)

	err := svc.DeleteBook(bookID)
	assert.NotNil(t, err)
	assert.Equal(t, 404, err.(*errors.AppError).Code)

	mockRepo.AssertExpectations(t)
	mockS3.AssertExpectations(t)
}
