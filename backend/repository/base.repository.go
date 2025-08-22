package repository

import (
	"github.com/google/uuid"
	"github.com/techatikin/backend/dto"
	"gorm.io/gorm"
)

type BaseRepository[T any] struct {
	db *gorm.DB
}

func NewBaseRepository[T any](db *gorm.DB) *BaseRepository[T] {
	return &BaseRepository[T]{db: db}
}

func (r *BaseRepository[T]) FindAll(params dto.QueryParams) ([]T, dto.PaginationMeta, error) {
	var results []T
	var totalCount int64

	query := r.db.Model(new(T))

	if params.Query != "" {

		query = query.Where("name ILIKE ? OR title ILIKE ?", "%"+params.Query+"%", "%"+params.Query+"%")
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

func (r *BaseRepository[T]) FindByID(id uuid.UUID) (*T, error) {
	var entity T
	if err := r.db.First(&entity, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &entity, nil
}

func (r *BaseRepository[T]) Create(entity *T) (*T, error) {
	if err := r.db.Create(entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func (r *BaseRepository[T]) Update(id uuid.UUID, updates map[string]interface{}) (*T, error) {
	var entity T
	if err := r.db.Model(&entity).Where("id = ?", id).Updates(updates).Error; err != nil {
		return nil, err
	}
	if err := r.db.First(&entity, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *BaseRepository[T]) Delete(id uuid.UUID) error {
	var entity T
	return r.db.Delete(&entity, "id = ?", id).Error
}
