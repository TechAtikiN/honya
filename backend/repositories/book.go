package repositories

import (
	"github.com/techatikin/backend/models"
	"gorm.io/gorm"
)

// TBookRepository defines the interface for book-related database operations.
type TBookRepository interface {
	FindAll(query string, offset, limit int, category string, publicationYear int, rating float64, pages int) ([]models.Book, int64, error)
	FindByID(id string) (models.Book, error)
	Create(book models.Book) (models.Book, error)
	Delete(id string) error
}

type bookRepository struct {
	db *gorm.DB
}

func BookRepository(db *gorm.DB) TBookRepository {
	return &bookRepository{db}
}

func (r *bookRepository) FindAll(query string, offset, limit int, category string, publicationYear int, rating float64, pages int) ([]models.Book, int64, error) {
	var books []models.Book
	var totalCount int64

	bookQuery := r.db.Model(&models.Book{})

	// Apply full-text query
	if query != "" {
		likeQuery := "%" + query + "%"
		bookQuery = bookQuery.Where("title ILIKE ? OR description ILIKE ? OR author_name ILIKE ?", likeQuery, likeQuery, likeQuery)
	}

	// Apply category filter
	if category != "" {
		bookQuery = bookQuery.Where("category = ?", category)
	}

	// Apply publication year logic
	if publicationYear > 0 {
		if publicationYear < 1950 {
			bookQuery = bookQuery.Where("publication_year <= ?", 1950)
		} else {
			bookQuery = bookQuery.Where("publication_year <= ?", publicationYear)
		}
	}

	if rating > 0 {
		bookQuery = bookQuery.Where("rating >= ?", rating)
	}

	if pages > 0 {
		bookQuery = bookQuery.Where("pages <= ?", pages)
	}

	err := bookQuery.Count(&totalCount).Error
	if err != nil {
		return nil, 0, err
	}

	// Then: apply pagination and fetch results
	err = bookQuery.Offset(offset).Limit(limit).Find(&books).Error
	if err != nil {
		return nil, 0, err
	}

	return books, totalCount, nil
}

func (r *bookRepository) FindByID(id string) (models.Book, error) {
	var book models.Book
	err := r.db.First(&book, id).Error
	return book, err
}

func (r *bookRepository) Create(book models.Book) (models.Book, error) {
	err := r.db.Create(&book).Error
	return book, err
}

func (r *bookRepository) Delete(id string) error {
	return r.db.Delete(&models.Book{}, id).Error
}
