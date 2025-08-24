package repository_test

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/techatikin/backend/dto"
	"github.com/techatikin/backend/model"
	"github.com/techatikin/backend/repository"
	"github.com/techatikin/backend/utils"
)

// Function to create a BookRepository with mock DB
func NewMockBookRepository(t *testing.T) (*repository.BookRepositoryImpl, sqlmock.Sqlmock, func()) {
	db, mock, cleanup := utils.NewMockDB(t)
	repo := &repository.BookRepositoryImpl{
		BaseRepository: repository.NewBaseRepository[model.Book](db),
	}
	return repo, mock, cleanup
}

func TestBookRepository_FindAll_WithFilters(t *testing.T) {
	repo, mock, cleanup := NewMockBookRepository(t)
	defer cleanup()

	rows := mock.NewRows([]string{"id", "title", "description", "author_name"}).
		AddRow(uuid.New(), "Book A", "Description A", "Author A").
		AddRow(uuid.New(), "Book B", "Description B", "Author B")

	mock.ExpectQuery(`SELECT count\(\*\) FROM "books"`).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(2))

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "books"`)).
		WillReturnRows(rows)

	params := dto.BookQueryParams{
		Query:           "Book",
		Category:        "",
		PublicationYear: 0,
		Rating:          0,
		Pages:           0,
		Sort:            "title",
		Limit:           10,
		Offset:          0,
	}

	books, meta, err := repo.FindAll(params)
	assert.NoError(t, err)
	assert.Len(t, books, 2)
	assert.Equal(t, int64(2), meta.TotalCount)
	assert.Equal(t, "Book A", books[0].Title)

	assert.NoError(t, mock.ExpectationsWereMet())
}
