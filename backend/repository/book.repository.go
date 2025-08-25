package repository

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/techatikin/backend/config"
	"github.com/techatikin/backend/dto"
	"github.com/techatikin/backend/model"
	"gorm.io/gorm"
)

// BookRepository defines the interface for book data operations
type BookRepository interface {
	FindAll(params dto.BookQueryParams) ([]model.Book, dto.PaginationMeta, error)
	FindByID(id uuid.UUID) (*model.Book, error)
	Create(book *model.Book) (*model.Book, error)
	Update(id uuid.UUID, updateData *dto.BookUpdateRequest) (*model.Book, error)
	Delete(id uuid.UUID) error
	CountByField(field string) (map[string]int64, error)
}

type BookRepositoryImpl struct {
	*BaseRepository[model.Book]
}

func NewBookRepository() BookRepository {
	return &BookRepositoryImpl{
		BaseRepository: NewBaseRepository[model.Book](config.DB.Db),
	}
}

func (r *BookRepositoryImpl) FindAll(params dto.BookQueryParams) ([]model.Book, dto.PaginationMeta, error) {
	var books []model.Book
	var totalCount int64

	query := r.db.Model(&model.Book{})

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
	case "title":
		query = query.Order("title ASC")
	case "rating":
		query = query.Order("rating DESC")
	case "recently_added":
		query = query.Order("created_at DESC")
	case "recently_updated":
		query = query.Order("updated_at DESC")
	case "pages":
		query = query.Order("pages DESC")
	case "publication_year":
		query = query.Order("publication_year DESC")
	default:
		query = query.Order("created_at DESC")
	}

	if err := query.Count(&totalCount).Error; err != nil {
		return nil, dto.PaginationMeta{}, err
	}

	if err := query.Offset(params.Offset).Limit(params.Limit).Find(&books).Error; err != nil {
		return nil, dto.PaginationMeta{}, err
	}

	meta := dto.PaginationMeta{
		TotalCount: totalCount,
		Offset:     params.Offset,
		Limit:      params.Limit,
	}

	return books, meta, nil
}

func (r *BookRepositoryImpl) Update(id uuid.UUID, updateData *dto.BookUpdateRequest) (*model.Book, error) {
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

func (r *BookRepositoryImpl) CountByField(field string) (map[string]int64, error) {
	var results []struct {
		Key   *string `gorm:"column:key"`
		Count int64   `gorm:"column:count"`
	}

	allowedFields := map[string]bool{
		"category":    true,
		"rating":      true,
		"author_name": true,
	}

	if !allowedFields[field] {
		return nil, errors.New("invalid field for aggregation")
	}

	query := r.db.Model(&model.Book{}).
		Select(fmt.Sprintf("%s as key, COUNT(*) as count", field)).
		Group(field)

	if field == "rating" {
		query = query.Where(fmt.Sprintf("%s IS NOT NULL", field))
	}

	if err := query.Scan(&results).Error; err != nil {
		return nil, errors.New("failed to execute aggregation query")
	}

	counts := make(map[string]int64)
	for _, result := range results {
		var keyStr string

		if result.Key == nil {
			keyStr = "Unknown"
		} else {
			keyStr = *result.Key

			if keyStr == "" {
				keyStr = "Unknown"
			}
		}

		counts[keyStr] = result.Count
	}

	return counts, nil
}
