package repository_test

import (
	"honya/backend/dto"
	"honya/backend/model"
	"honya/backend/repository"
	"honya/backend/utils"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func NewMockReviewRepository(t *testing.T) (*repository.ReviewRepositoryImpl, sqlmock.Sqlmock, func()) {
	db, mock, cleanup := utils.NewMockDB(t)
	repo := &repository.ReviewRepositoryImpl{
		BaseRepository: repository.NewBaseRepository[model.Review](db),
	}
	return repo, mock, cleanup
}

func TestReviewRepository_FindByBookID(t *testing.T) {
	repo, mock, cleanup := NewMockReviewRepository(t)
	defer cleanup()

	bookID := uuid.New()

	rows := mock.NewRows([]string{"id", "book_id", "name", "email", "content", "created_at", "updated_at"}).
		AddRow(uuid.New(), bookID, "Reviewer A", "a@example.com", "Great book!", int64(1640995200), int64(1640995200)).
		AddRow(uuid.New(), bookID, "Reviewer B", "b@example.com", "Loved it!", int64(1640995200), int64(1640995200))

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT count(*) FROM "reviews" WHERE book_id = $1`)).
		WithArgs(bookID).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(2))

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "reviews" WHERE book_id = $1 LIMIT $2`)).
		WithArgs(bookID, 10).
		WillReturnRows(rows)

	params := dto.QueryParams{
		Limit:  10,
		Offset: 0,
		Query:  "",
	}

	reviews, meta, err := repo.FindByBookID(bookID, params)
	assert.NoError(t, err)
	assert.Len(t, reviews, 2)
	assert.Equal(t, int64(2), meta.TotalCount)
	assert.Equal(t, "Reviewer A", reviews[0].Name)

	assert.NoError(t, mock.ExpectationsWereMet())
}
