package service

import (
	"strings"

	"github.com/google/uuid"
	"github.com/techatikin/backend/dtos"
	"github.com/techatikin/backend/errors"
	"github.com/techatikin/backend/model"
	"github.com/techatikin/backend/repositories"
	"github.com/techatikin/backend/utils"
)

// TBookService defines the interface for book-related services.
type TBookService interface {
	GetBooks(params dtos.BookQueryParams) ([]model.Book, *dtos.PaginationMeta, error)
	GetBookByID(id uuid.UUID) (*model.Book, error)
	CreateBook(book *dtos.BookCreateRequest) (*model.Book, error)
	UpdateBook(id uuid.UUID, updateData *dtos.BookUpdateRequest) (*model.Book, error)
	DeleteBook(id uuid.UUID) error
}

type bookService struct {
	bookRepo repositories.TBookRepository
	// s3repo repositories.TS3Repository // S3 repository for handling images
}

func BookService(repo repositories.TBookRepository) TBookService {
	return &bookService{repo}
}

func (s *bookService) GetBooks(params dtos.BookQueryParams) ([]model.Book, *dtos.PaginationMeta, error) {
	books, meta, err := s.bookRepo.FindAll(params)
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

func (s *bookService) GetBookByID(id uuid.UUID) (*model.Book, error) {

	book, err := s.bookRepo.FindByID(id)
	if err != nil {
		return nil, errors.NewInternalError(err)
	}
	if book == nil {
		return nil, errors.NewNotFoundError("Book not found")
	}

	return book, nil
}

func (s *bookService) CreateBook(book *dtos.BookCreateRequest) (*model.Book, error) {
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
	newBook := model.Book{
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
	resource, err := s.bookRepo.Create(&newBook)
	if err != nil {
		if strings.Contains(err.Error(), "unique constraint") {
			return nil, errors.NewConflictError("Book with this ISBN already exists")
		}
		// Delete the uploaded image if book creation fails
		// s.s3repo.DeleteImage(imageURL)
		return nil, errors.NewInternalError(err)
	}

	return resource, nil
}

func (s *bookService) UpdateBook(id uuid.UUID, updateData *dtos.BookUpdateRequest) (*model.Book, error) {
	if updateData.Isbn != nil {
		return nil, errors.NewBadRequestError("ISBN cannot be updated once set")
	}

	if err := utils.ValidateBookUpdateRequest(updateData); err != nil {
		return nil, errors.NewBadRequestError(err.Error())
	}

	existingBook, err := s.bookRepo.FindByID(id)
	if err != nil {
		return nil, errors.NewInternalError(err)
	}
	if existingBook == nil {
		return nil, errors.NewNotFoundError("Book not found")
	}

	resource, err := s.bookRepo.Update(id, updateData)
	if err != nil {
		return nil, errors.NewInternalError(err)
	}

	return resource, nil
}

func (s *bookService) DeleteBook(id uuid.UUID) error {
	existingBook, err := s.bookRepo.FindByID(id)
	if err != nil {
		return errors.NewInternalError(err)
	}

	if existingBook == nil {
		return errors.NewNotFoundError("Book not found")
	}

	return s.bookRepo.Delete(id)
}
