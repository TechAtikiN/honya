package services

import (
	"github.com/techatikin/backend/dtos"
	"github.com/techatikin/backend/models"
	"github.com/techatikin/backend/repositories"
	"github.com/techatikin/backend/utils"
)

// TBookService defines the interface for book-related services.
type TBookService interface {
	GetBooks() ([]models.Book, error)
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

func (s *bookService) GetBooks() ([]models.Book, error) {
	return s.repo.FindAll()
}

func (s *bookService) GetBookByID(id string) (models.Book, error) {
	return s.repo.FindByID(id)
}

func (s *bookService) CreateBook(book dtos.BookCreateRequest) (models.Book, error) {
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
		AuthorId:        book.AuthorId,
	}

	return s.repo.Create(resource)
}

func (s *bookService) DeleteBook(id string) error {
	return s.repo.Delete(id)
}
