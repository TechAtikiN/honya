package services

import (
	"github.com/techatikin/backend/dtos"
	"github.com/techatikin/backend/errors"
	"github.com/techatikin/backend/models"
	"github.com/techatikin/backend/repositories"
	"github.com/techatikin/backend/utils"
)

// TBookService defines the interface for book-related services.
type TBookService interface {
	GetBooks(params dtos.BookQueryParams) ([]models.Book, *dtos.PaginationMeta, error)
	GetBookByID(id string) (*models.Book, error)
	CreateBook(book *dtos.BookCreateRequest) (*models.Book, error)
	UpdateBook(id string, updateData *dtos.BookUpdateRequest) (*models.Book, error)
	DeleteBook(id string) error
}

type bookService struct {
	repo repositories.TBookRepository
	// s3repo repositories.TS3Repository // S3 repository for handling images
}

func BookService(repo repositories.TBookRepository) TBookService {
	return &bookService{repo}
}

func (s *bookService) GetBooks(params dtos.BookQueryParams) ([]models.Book, *dtos.PaginationMeta, error) {
	books, meta, err := s.repo.FindAll(params)
	if err != nil {
		return nil, nil, errors.NewInternalError(err)
	}

	if len(books) == 0 {
		return books, &dtos.PaginationMeta{
			TotalCount: 0,
			Offset:     params.Offset,
			Limit:      params.Limit,
		}, nil
	}

	return books, &meta, nil
}

func (s *bookService) GetBookByID(id string) (*models.Book, error) {
	if id == "" {
		return nil, errors.NewBadRequestError("ID is required")
	}

	book, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.NewInternalError(err)
	}

	if book == nil {
		return nil, errors.NewNotFoundError("Book not found")
	}

	return book, nil
}

func (s *bookService) CreateBook(book *dtos.BookCreateRequest) (*models.Book, error) {
	// Validate the book creation request
	if err := utils.ValidateBookCreateRequest(book); err != nil {
		return nil, errors.NewBadRequestError(err.Error())
	}

	// Call S3 repository to upload the image
	// imageURL, err := s.s3repo.UploadImage(book.Image)
	// if err != nil {
	// 	return nil, errors.NewInternalError(err)
	// }

	// Create a new book model
	newBook := models.Book{
		Title:           book.Title,
		Description:     book.Description,
		Category:        book.Category,
		Image:           book.Image, // imageURL,
		PublicationYear: book.PublicationYear,
		Rating:          book.Rating,
		Pages:           book.Pages,
		Isbn:            book.Isbn,
		AuthorName:      book.AuthorName,
	}

	// Call repository to create the book
	resource, err := s.repo.Create(&newBook)
	if err != nil {
		// Delete the uploaded image if book creation fails
		// s.s3repo.DeleteImage(imageURL)
		return nil, errors.NewInternalError(err)
	}

	return resource, nil
}

func (s *bookService) UpdateBook(id string, updateData *dtos.BookUpdateRequest) (*models.Book, error) {
	if id == "" {
		return nil, errors.NewBadRequestError("ID is required")
	}

	// Validate input
	if err := utils.ValidateBookUpdateRequest(updateData); err != nil {
		return nil, errors.NewBadRequestError(err.Error())
	}

	// Check if book exists
	existingBook, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.NewInternalError(err)
	}

	if existingBook == nil {
		return nil, errors.NewNotFoundError("Book not found")
	}

	// Perform update
	resource, err := s.repo.Update(id, updateData)
	if err != nil {
		return nil, errors.NewInternalError(err)
	}

	return resource, nil
}

func (s *bookService) DeleteBook(id string) error {
	// Check if book exists
	existingBook, err := s.repo.FindByID(id)
	if err != nil {
		return errors.NewInternalError(err)
	}

	if existingBook == nil {
		return errors.NewNotFoundError("Book not found")
	}

	return s.repo.Delete(id)
}
