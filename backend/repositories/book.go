package repositories

import (
	"github.com/google/uuid"
	"github.com/techatikin/backend/dtos"
	"github.com/techatikin/backend/models"
	"gorm.io/gorm"
)

// TBookRepository defines the interface for book-related database operations.
type TBookRepository interface {
	FindAll(params dtos.BookQueryParams) ([]models.Book, dtos.PaginationMeta, error)
	FindByID(id uuid.UUID) (*models.Book, error)
	Create(book *models.Book) (*models.Book, error)
	Update(id uuid.UUID, updateData *dtos.BookUpdateRequest) (*models.Book, error)
	Delete(id uuid.UUID) error
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

	if params.Query != "" {
		likeQuery := "%" + params.Query + "%"
		query = query.Where("title ILIKE ? OR description ILIKE ? OR author_name ILIKE ?", likeQuery, likeQuery, likeQuery)
	}

	if params.Category != "" {
		query = query.Where("category = ?", params.Category)
	}

	if params.PublicationYear > 0 {
		query = query.Where("publication_year <= ?", params.PublicationYear)
	}

	if params.Rating > 0 {
		query = query.Where("rating >= ?", params.Rating)
	}

	if params.Pages > 0 {
		query = query.Where("pages <= ?", params.Pages)
	}

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

	if err := query.Count(&totalCount).Error; err != nil {
		return nil, dtos.PaginationMeta{}, err
	}

	if err := query.Offset(params.Offset).Limit(params.Limit).Find(&books).Error; err != nil {
		return nil, dtos.PaginationMeta{}, err
	}

	meta := dtos.PaginationMeta{
		TotalCount: totalCount,
		Offset:     params.Offset,
		Limit:      params.Limit,
	}

	return books, meta, nil
}

func (r *bookRepository) FindByID(id uuid.UUID) (*models.Book, error) {
	var book models.Book

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
	err := r.db.Create(book).Error
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (r *bookRepository) Update(id uuid.UUID, updateData *dtos.BookUpdateRequest) (*models.Book, error) {
	book, err := r.FindByID(id)
	if err != nil {
		return nil, err
	}
	if book == nil {
		return nil, gorm.ErrRecordNotFound
	}

	updates := map[string]interface{}{}

	if updateData.Title != nil {
		updates["title"] = *updateData.Title
	}
	if updateData.Description != nil {
		updates["description"] = *updateData.Description
	}
	if updateData.Category != nil {
		updates["category"] = *updateData.Category
	}
	if updateData.Image != nil {
		updates["image"] = *updateData.Image
	}
	if updateData.PublicationYear != nil {
		updates["publication_year"] = *updateData.PublicationYear
	}
	if updateData.Rating != nil {
		updates["rating"] = *updateData.Rating
	}
	if updateData.Pages != nil {
		updates["pages"] = *updateData.Pages
	}
	if updateData.AuthorName != nil {
		updates["author_name"] = *updateData.AuthorName
	}

	if len(updates) == 0 {
		return book, nil
	}

	if err := r.db.Model(book).Where("id = ?", id).Updates(updates).Error; err != nil {
		return nil, err
	}

	if err := r.db.First(book, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return book, nil
}

func (r *bookRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Book{}, "id = ?", id).Error
}
