package repositories

import (
	"github.com/techatikin/backend/models"
	"gorm.io/gorm"
)

// TBookRepository defines the interface for book-related database operations.
type TBookRepository interface {
	FindAll() ([]models.Book, error)
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

func (r *bookRepository) FindAll() ([]models.Book, error) {
	var books []models.Book
	err := r.db.Find(&books).Error
	return books, err
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
