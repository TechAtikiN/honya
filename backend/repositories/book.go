package repositories

import (
	"github.com/techatikin/backend/dtos"
	"github.com/techatikin/backend/models"
	"gorm.io/gorm"
)

// TBookRepository defines the interface for book-related database operations.
type TBookRepository interface {
	FindAll(params dtos.BookQueryParams) ([]models.Book, int64, error)
	FindByID(id string) (models.Book, error)
	Create(book models.Book) (models.Book, error)
	Update(id string, updates dtos.BookUpdateRequest) (models.Book, error)
	Delete(id string) error
}

type bookRepository struct {
	db *gorm.DB
}

func BookRepository(db *gorm.DB) TBookRepository {
	return &bookRepository{db}
}

func (r *bookRepository) FindAll(params dtos.BookQueryParams) ([]models.Book, int64, error) {
	var books []models.Book
	var totalCount int64

	bookQuery := r.db.Model(&models.Book{})

	// Apply full-text query
	if params.Query != "" {
		likeQuery := "%" + params.Query + "%"
		bookQuery = bookQuery.Where("title ILIKE ? OR description ILIKE ? OR author_name ILIKE ?", likeQuery, likeQuery, likeQuery)
	}

	// Apply category filter
	if params.Category != "" {
		bookQuery = bookQuery.Where("category = ?", params.Category)
	}

	// Apply publication year logic
	if params.PublicationYear > 0 {
		if params.PublicationYear < 1950 {
			bookQuery = bookQuery.Where("publication_year <= ?", 1950)
		} else {
			bookQuery = bookQuery.Where("publication_year <= ?", params.PublicationYear)
		}
	}

	if params.Rating > 0 {
		bookQuery = bookQuery.Where("rating >= ?", params.Rating)
	}

	if params.Pages > 0 {
		bookQuery = bookQuery.Where("pages <= ?", params.Pages)
	}

	err := bookQuery.Count(&totalCount).Error
	if err != nil {
		return nil, 0, err
	}

	// Then: apply pagination and fetch results
	err = bookQuery.Offset(params.Offset).Limit(params.Limit).Find(&books).Error
	if err != nil {
		return nil, 0, err
	}

	return books, totalCount, nil
}

func (r *bookRepository) FindByID(id string) (models.Book, error) {
	var book models.Book
	err := r.db.First(&book, "id = ?", id).Error
	return book, err
}

func (r *bookRepository) Create(book models.Book) (models.Book, error) {
	err := r.db.Create(&book).Error
	return book, err
}

func (r *bookRepository) Update(id string, updates dtos.BookUpdateRequest) (models.Book, error) {
	var book models.Book

	// Find existing book
	if err := r.db.First(&book, "id = ?", id).Error; err != nil {
		return models.Book{}, err
	}

	// Apply updates only for non-nil fields
	if updates.Title != nil {
		book.Title = *updates.Title
	}
	if updates.Description != nil {
		book.Description = *updates.Description
	}
	if updates.Category != nil {
		book.Category = *updates.Category
	}
	if updates.Image != nil {
		book.Image = *updates.Image
	}
	if updates.PublicationYear != nil {
		book.PublicationYear = *updates.PublicationYear
	}
	if updates.Rating != nil {
		book.Rating = *updates.Rating
	}
	if updates.Pages != nil {
		book.Pages = *updates.Pages
	}
	if updates.Isbn != nil {
		book.Isbn = *updates.Isbn
	}
	if updates.AuthorName != nil {
		book.AuthorName = *updates.AuthorName
	}

	// Save changes
	if err := r.db.Save(&book).Error; err != nil {
		return models.Book{}, err
	}

	return book, nil
}

func (r *bookRepository) Delete(id string) error {
	return r.db.Delete(&models.Book{}, "id = ?", id).Error
}
