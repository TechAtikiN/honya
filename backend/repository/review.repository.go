package repository

import (
	"errors"

	"github.com/google/uuid"
	"github.com/techatikin/backend/config"
	"github.com/techatikin/backend/dto"
	"github.com/techatikin/backend/model"
)

// ReviewRepository defines methods for interacting with the reviews in the database.
type ReviewRepository interface {
	FindAll(params dto.QueryParams) ([]model.Review, dto.PaginationMeta, error)
	FindByID(id uuid.UUID) (*model.Review, error)
	Create(review *model.Review) (*model.Review, error)
	Update(id uuid.UUID, updates map[string]interface{}) (*model.Review, error)
	Delete(id uuid.UUID) error
	FindByBookID(bookID uuid.UUID, params dto.QueryParams) ([]model.Review, dto.PaginationMeta, error)
	GetTopReviewers(limit int) ([]dto.ReviewerStats, error)
}

type ReviewRepositoryImpl struct {
	*BaseRepository[model.Review]
}

func NewReviewRepository() ReviewRepository {
	return &ReviewRepositoryImpl{
		BaseRepository: NewBaseRepository[model.Review](config.DB.Db),
	}
}

func (r *ReviewRepositoryImpl) FindByBookID(bookID uuid.UUID, params dto.QueryParams) ([]model.Review, dto.PaginationMeta, error) {
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

func (r *ReviewRepositoryImpl) GetTopReviewers(limit int) ([]dto.ReviewerStats, error) {
	var results []struct {
		Name  *string `gorm:"column:name"`
		Count int64   `gorm:"column:count"`
	}

	if limit <= 0 {
		limit = 10
	}

	query := r.db.Model(&model.Review{}).
		Select("name, COUNT(*) as count").
		Where("name IS NOT NULL AND name != ''").
		Group("name").
		Order("count DESC").
		Limit(limit)

	if err := query.Scan(&results).Error; err != nil {
		return nil, errors.New("failed to fetch top reviewers: " + err.Error())
	}

	reviewerStats := make([]dto.ReviewerStats, 0, len(results))
	for _, result := range results {
		var name string
		if result.Name == nil || *result.Name == "" {
			name = "Anonymous"
		} else {
			name = *result.Name
		}

		reviewerStats = append(reviewerStats, dto.ReviewerStats{
			Name:  name,
			Count: result.Count,
		})
	}

	return reviewerStats, nil
}
