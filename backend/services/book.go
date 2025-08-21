package services

import (
	"github.com/techatikin/backend/dtos"
	"github.com/techatikin/backend/models"
	"github.com/techatikin/backend/repositories"
	"github.com/techatikin/backend/utils"
)

// TBookService defines the interface for book-related services.
type TBookService interface {
	GetBooks(params dtos.BookQueryParams) ([]models.Book, int64, error)
	GetBookByID(id string) (models.Book, error)
	CreateBook(dtos.BookCreateRequest) (models.Book, error)
	UpdateBook(id string, updates dtos.BookUpdateRequest) (models.Book, error) // <-- added
	DeleteBook(id string) error
}

type bookService struct {
	repo repositories.TBookRepository
}

func BookService(repo repositories.TBookRepository) TBookService {
	return &bookService{repo}
}

func (s *bookService) GetBooks(params dtos.BookQueryParams) ([]models.Book, int64, error) {
	return s.repo.FindAll(params)
}

func (s *bookService) GetBookByID(id string) (models.Book, error) {
	return s.repo.FindByID(id)
}

func (s *bookService) CreateBook(book dtos.BookCreateRequest) (models.Book, error) {
	// Validate the book creation request
	if err := utils.ValidateBookCreateRequest(book); err != nil {
		return models.Book{}, err
	}

	resource := models.Book{
		Title:           book.Title,
		Description:     book.Description,
		Category:        book.Category,
		Image:           book.Image,
		PublicationYear: book.PublicationYear,
		Rating:          book.Rating,
		Pages:           book.Pages,
		Isbn:            book.Isbn,
		AuthorName:      book.AuthorName,
	}

	return s.repo.Create(resource)
}

func (s *bookService) UpdateBook(id string, updates dtos.BookUpdateRequest) (models.Book, error) {
	return s.repo.Update(id, updates)
}

func (s *bookService) DeleteBook(id string) error {
	return s.repo.Delete(id)
}
