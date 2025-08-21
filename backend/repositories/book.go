package repositories

import (
	"github.com/techatikin/backend/dtos"
	"github.com/techatikin/backend/models"
	"gorm.io/gorm"
)

// TBookRepository defines the interface for book-related database operations.
type TBookRepository interface {
	FindAll(params dtos.BookQueryParams) ([]models.Book, dtos.PaginationMeta, error)
	FindByID(id string) (*models.Book, error)
	Create(book *models.Book) (*models.Book, error)
	Update(id string, updateData *dtos.BookUpdateRequest) (*models.Book, error)
	Delete(id string) error
}

type bookRepository struct {
	db *gorm.DB
}

func BookRepository(db *gorm.DB) TBookRepository {
	return &bookRepository{db}
}

func (r *bookRepository) FindAll(params dtos.BookQueryParams) ([]models.Book, dtos.PaginationMeta, error) {
	// Variales to hold results and metadata
	var books []models.Book
	var totalCount int64

	// Build the query
	query := r.db.Model(&models.Book{})

	// Apply filters based on search query
	if params.Query != "" {
		likeQuery := "%" + params.Query + "%"
		query = query.Where("title ILIKE ? OR description ILIKE ? OR author_name ILIKE ?", likeQuery, likeQuery, likeQuery)
	}

	// Apply category filter
	if params.Category != "" {
		query = query.Where("category = ?", params.Category)
	}

	// Apply publication year logic
	if params.PublicationYear > 0 {
		if params.PublicationYear < 1950 {
			query = query.Where("publication_year <= ?", 1950)
		} else {
			query = query.Where("publication_year <= ?", params.PublicationYear)
		}
	}

	// Apply rating and pages filters
	if params.Rating > 0 {
		query = query.Where("rating >= ?", params.Rating)
	}

	// Apply pages filter
	if params.Pages > 0 {
		query = query.Where("pages <= ?", params.Pages)
	}

	// Count total records for pagination
	err := query.Count(&totalCount).Error
	if err != nil {
		return nil, dtos.PaginationMeta{}, err
	}

	// Apply pagination
	err = query.Offset(params.Offset).Limit(params.Limit).Find(&books).Error
	if err != nil {
		return nil, dtos.PaginationMeta{}, err
	}

	// Prepare pagination metadata
	meta := dtos.PaginationMeta{
		TotalCount: int64(totalCount),
		Offset:     params.Offset,
		Limit:      params.Limit,
	}

	return books, meta, nil
}

func (r *bookRepository) FindByID(id string) (*models.Book, error) {
	var book models.Book

	// Find book by ID
	err := r.db.First(&book, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &book, nil
}

func (r *bookRepository) Create(book *models.Book) (*models.Book, error) {
	// Create a new book record
	err := r.db.Create(book).Error
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (r *bookRepository) Update(id string, updateData *dtos.BookUpdateRequest) (*models.Book, error) {
	book, err := r.FindByID(id)
	if err != nil {
		return nil, err
	}
	if book == nil {
		return nil, gorm.ErrRecordNotFound
	}

	// Update only non-nil fields
	if updateData.Title != nil {
		book.Title = *updateData.Title
	}
	if updateData.Description != nil {
		book.Description = *updateData.Description
	}
	if updateData.Category != nil {
		book.Category = *updateData.Category
	}
	if updateData.Image != nil {
		book.Image = *updateData.Image
	}
	if updateData.PublicationYear != nil {
		book.PublicationYear = *updateData.PublicationYear
	}
	if updateData.Rating != nil {
		book.Rating = *updateData.Rating
	}
	if updateData.Pages != nil {
		book.Pages = *updateData.Pages
	}
	if updateData.Isbn != nil {
		book.Isbn = *updateData.Isbn
	}
	if updateData.AuthorName != nil {
		book.AuthorName = *updateData.AuthorName
	}

	// Save the updated book
	if err := r.db.Save(book).Error; err != nil {
		return nil, err
	}

	return book, nil
}

func (r *bookRepository) Delete(id string) error {
	return r.db.Delete(&models.Book{}, "id = ?", id).Error
}
