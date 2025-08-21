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
	var books []models.Book
	var totalCount int64

	query := r.db.Model(&models.Book{})

	// Search filter
	if params.Query != "" {
		likeQuery := "%" + params.Query + "%"
		query = query.Where("title ILIKE ? OR description ILIKE ? OR author_name ILIKE ?", likeQuery, likeQuery, likeQuery)
	}

	// Category filter
	if params.Category != "" {
		query = query.Where("category = ?", params.Category)
	}

	// Publication year filter (range logic can be added if needed)
	if params.PublicationYear > 0 {
		query = query.Where("publication_year <= ?", params.PublicationYear)
	}

	// Rating filter
	if params.Rating > 0 {
		query = query.Where("rating >= ?", params.Rating)
	}

	// Pages filter
	if params.Pages > 0 {
		query = query.Where("pages <= ?", params.Pages)
	}

	// Sorting
	switch params.Sort {
	case "title_asc":
		query = query.Order("title ASC")
	case "rating_desc":
		query = query.Order("rating DESC")
	case "created_at_desc":
		query = query.Order("created_at DESC")
	case "updated_at_desc":
		query = query.Order("updated_at DESC")
	case "pages_desc":
		query = query.Order("pages DESC")
	case "publication_year_desc":
		query = query.Order("publication_year DESC")
	default:
		query = query.Order("created_at DESC")
	}

	// Count total records
	if err := query.Count(&totalCount).Error; err != nil {
		return nil, dtos.PaginationMeta{}, err
	}

	// Apply pagination
	if err := query.Offset(params.Offset).Limit(params.Limit).Find(&books).Error; err != nil {
		return nil, dtos.PaginationMeta{}, err
	}

	// Pagination metadata
	meta := dtos.PaginationMeta{
		TotalCount: totalCount,
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
