package repository

import (
	"github.com/google/uuid"
	"github.com/techatikin/backend/config"
	"github.com/techatikin/backend/dto"
	"github.com/techatikin/backend/model"
)

type ReviewRepository interface {
	FindAll(params dto.QueryParams) ([]model.Review, dto.PaginationMeta, error)
	FindByID(id uuid.UUID) (*model.Review, error)
	Create(review *model.Review) (*model.Review, error)
	Update(id uuid.UUID, updates map[string]interface{}) (*model.Review, error)
	Delete(id uuid.UUID) error
	FindByBookID(bookID uuid.UUID, params dto.QueryParams) ([]model.Review, dto.PaginationMeta, error)
}

type reviewRepository struct {
	*BaseRepository[model.Review]
}

func NewReviewRepository() ReviewRepository {
	return &reviewRepository{
		BaseRepository: NewBaseRepository[model.Review](config.DB.Db),
	}
}

func (r *reviewRepository) FindByBookID(bookID uuid.UUID, params dto.QueryParams) ([]model.Review, dto.PaginationMeta, error) {
	var results []model.Review
	var totalCount int64

	query := r.db.Model(&model.Review{}).Where("book_id = ?", bookID)

	if params.Query != "" {
		query = query.Where("content ILIKE ? OR name ILIKE ?", "%"+params.Query+"%", "%"+params.Query+"%")
	}

	if err := query.Count(&totalCount).Error; err != nil {
		return nil, dto.PaginationMeta{}, err
	}

	if params.Limit > 0 {
		query = query.Limit(params.Limit)
	}
	if params.Offset >= 0 {
		query = query.Offset(params.Offset)
	}

	if err := query.Find(&results).Error; err != nil {
		return nil, dto.PaginationMeta{}, err
	}

	meta := dto.PaginationMeta{
		TotalCount: totalCount,
		Limit:      params.Limit,
		Offset:     params.Offset,
	}

	return results, meta, nil
}
