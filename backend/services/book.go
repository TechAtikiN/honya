package services

import (
	"github.com/techatikin/backend/dtos"
	"github.com/techatikin/backend/models"
	"github.com/techatikin/backend/repositories"
	"github.com/techatikin/backend/utils"
)

// TBookService defines the interface for book-related services.
type TBookService interface {
	GetBooks(query string, offset, limit int, category string, publication_year int, rating float64, pages int) ([]models.Book, int64, error)
	GetBookByID(id string) (models.Book, error)
	CreateBook(dtos.BookCreateRequest) (models.Book, error)
	DeleteBook(id string) error
}

type bookService struct {
	repo repositories.TBookRepository
}

func BookService(repo repositories.TBookRepository) TBookService {
	return &bookService{repo}
}

func (s *bookService) GetBooks(query string, offset, limit int, category string, publicationYear int, rating float64, pages int) ([]models.Book, int64, error) {
	return s.repo.FindAll(query, offset, limit, category, publicationYear, rating, pages)
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

func (s *bookService) DeleteBook(id string) error {
	return s.repo.Delete(id)
}
