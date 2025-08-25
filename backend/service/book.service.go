package service

import (
	"mime/multipart"

	"github.com/google/uuid"
	"github.com/techatikin/backend/config"
	"github.com/techatikin/backend/dto"
	"github.com/techatikin/backend/errors"
	"github.com/techatikin/backend/model"
	"github.com/techatikin/backend/repository"
	"github.com/techatikin/backend/utils"
)

var (
	env, _     = config.GetEnvConfig()
	AWS_BUCKET = env.AWSBucket
	AWS_REGION = env.AWSRegion
)

// BookService defines the interface for book-related services.
type BookService interface {
	GetBooks(params dto.BookQueryParams) ([]model.Book, *dto.PaginationMeta, error)
	GetBookByID(id uuid.UUID) (*model.Book, error)
	CreateBook(book *dto.BookCreateRequest, fileHeader *multipart.FileHeader) (*model.Book, error)
	UpdateBook(id uuid.UUID, updateData *dto.BookUpdateRequest, fileHeader *multipart.FileHeader) (*model.Book, error)
	DeleteBook(id uuid.UUID) error
}

type bookService struct {
	repo   repository.BookRepository
	s3repo repository.S3Repository
}

func NewBookService(repo repository.BookRepository, s3repo repository.S3Repository) BookService {
	return &bookService{repo, s3repo}
}

func (s *bookService) GetBooks(params dto.BookQueryParams) ([]model.Book, *dto.PaginationMeta, error) {
	books, meta, err := s.repo.FindAll(params)
	if err != nil {
		return nil, nil, errors.NewInternalError(err)
	}

	if len(books) == 0 {
		return books, &dto.PaginationMeta{
			TotalCount: 0,
			Offset:     params.Offset,
			Limit:      params.Limit,
		}, nil
	}

	return books, &meta, nil
}

func (s *bookService) GetBookByID(id uuid.UUID) (*model.Book, error) {

	book, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.NewInternalError(err)
	}
	if book == nil {
		return nil, errors.NewNotFoundError("Book not found")
	}

	return book, nil
}

func (s *bookService) CreateBook(book *dto.BookCreateRequest, fileHeader *multipart.FileHeader) (*model.Book, error) {
	if err := utils.ValidateBookCreateRequest(book); err != nil {
		return nil, errors.NewBadRequestError(err.Error())
	}

	var imageURL string
	if fileHeader != nil {
		url, err := s.s3repo.UploadImage(fileHeader, book.Title)
		if err != nil {
			return nil, errors.NewInternalError(err)
		}
		imageURL = url
	}

	newBook := model.Book{
		Title:           book.Title,
		Description:     book.Description,
		Category:        book.Category,
		Image:           imageURL,
		PublicationYear: book.PublicationYear,
		Rating:          book.Rating,
		Pages:           book.Pages,
		Isbn:            book.Isbn,
		AuthorName:      book.AuthorName,
	}

	resource, err := s.repo.Create(&newBook)
	if err != nil {
		if imageURL != "" {
			key := utils.ExtractS3Key(imageURL, AWS_BUCKET, AWS_REGION)
			_ = s.s3repo.DeleteImage(key)
		}
		return nil, errors.NewInternalError(err)
	}

	return resource, nil
}

func (s *bookService) UpdateBook(id uuid.UUID, updateData *dto.BookUpdateRequest, fileHeader *multipart.FileHeader) (*model.Book, error) {
	existingBook, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.NewInternalError(err)
	}
	if existingBook == nil {
		return nil, errors.NewNotFoundError("Book not found")
	}

	// If new image uploaded, replace old one
	if fileHeader != nil {
		url, err := s.s3repo.UploadImage(fileHeader, existingBook.Title)
		if err != nil {
			return nil, errors.NewInternalError(err)
		}
		updateData.Image = &url

		if existingBook.Image != "" {
			key := utils.ExtractS3Key(existingBook.Image, AWS_BUCKET, AWS_REGION)
			_ = s.s3repo.DeleteImage(key)
		}
	}

	resource, err := s.repo.Update(id, updateData)
	if err != nil {
		if updateData.Image != nil && *updateData.Image != "" {
			key := utils.ExtractS3Key(*updateData.Image, AWS_BUCKET, AWS_REGION)
			_ = s.s3repo.DeleteImage(key)
		}
		return nil, errors.NewInternalError(err)
	}

	return resource, nil
}

func (s *bookService) DeleteBook(id uuid.UUID) error {
	existingBook, err := s.repo.FindByID(id)
	if err != nil {
		return errors.NewInternalError(err)
	}

	if existingBook == nil {
		return errors.NewNotFoundError("Book not found")
	}

	// Delete image from S3 if exists
	if existingBook.Image != "" {
		key := utils.ExtractS3Key(existingBook.Image, AWS_BUCKET, AWS_REGION)
		_ = s.s3repo.DeleteImage(key)
	}

	// Delete book from DB
	return s.repo.Delete(id)
}
