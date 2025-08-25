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

func TestBaseRepository_FindAll(t *testing.T) {
	db, mock, cleanup := utils.NewMockDB(t)
	defer cleanup()

	repo := repository.NewBaseRepository[model.Book](db)

	rows := mock.NewRows([]string{"id", "title"}).
		AddRow(uuid.New(), "Book A").
		AddRow(uuid.New(), "Book B")

	mock.ExpectQuery(`SELECT count\(\*\) FROM "books"`).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(2))

	mock.ExpectQuery(`SELECT \* FROM "books"`).
		WillReturnRows(rows)

	books, meta, err := repo.FindAll(dto.QueryParams{Limit: 10, Offset: 0})
	assert.NoError(t, err)
	assert.Len(t, books, 2)
	assert.Equal(t, int64(2), meta.TotalCount)

	assert.NoError(t, mock.ExpectationsWereMet())
}
func TestBaseRepository_FindByID(t *testing.T) {
	db, mock, cleanup := utils.NewMockDB(t)
	defer cleanup()

	repo := repository.NewBaseRepository[model.Book](db)

	id := uuid.New()
	rows := mock.NewRows([]string{"id", "title"}).
		AddRow(id, "Test Book")

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "books" WHERE id = $1 ORDER BY "books"."id" LIMIT $2`,
	)).
		WithArgs(id, 1).
		WillReturnRows(rows)

	book, err := repo.FindByID(id)
	assert.NoError(t, err)
	assert.NotNil(t, book)
	assert.Equal(t, "Test Book", book.Title)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestBaseRepository_Create(t *testing.T) {
	db, mock, cleanup := utils.NewMockDB(t)
	defer cleanup()

	repo := repository.NewBaseRepository[model.Book](db)

	book := &model.Book{ID: uuid.New(), Title: "New Book"}

	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO "books"`).
		WithArgs(
			book.ID,
			book.Title,
			sqlmock.AnyArg(), // Description
			sqlmock.AnyArg(), // Category
			sqlmock.AnyArg(), // Image
			sqlmock.AnyArg(), // PublicationYear
			sqlmock.AnyArg(), // Rating
			sqlmock.AnyArg(), // Pages
			sqlmock.AnyArg(), // ISBN
			sqlmock.AnyArg(), // CreatedAt
			sqlmock.AnyArg(), // UpdatedAt
			sqlmock.AnyArg(), // AuthorName
		).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	created, err := repo.Create(book)
	assert.NoError(t, err)
	assert.Equal(t, "New Book", created.Title)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestBaseRepository_Update(t *testing.T) {
	db, mock, cleanup := utils.NewMockDB(t)
	defer cleanup()

	repo := repository.NewBaseRepository[model.Book](db)
	id := uuid.New()

	updateData := map[string]interface{}{
		"title": "New Title",
	}

	mock.ExpectBegin()

	mock.ExpectExec(`UPDATE "books" SET`).
		WithArgs("New Title", sqlmock.AnyArg(), id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	rows := mock.NewRows([]string{"id", "title"}).
		AddRow(id, "New Title")

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "books" WHERE id = $1 ORDER BY "books"."id" LIMIT $2`,
	)).
		WithArgs(id, 1).
		WillReturnRows(rows)

	updatedBook, err := repo.Update(id, updateData)
	assert.NoError(t, err)
	assert.NotNil(t, updatedBook)
	assert.Equal(t, "New Title", updatedBook.Title)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestBaseRepository_Delete(t *testing.T) {
	db, mock, cleanup := utils.NewMockDB(t)
	defer cleanup()

	repo := repository.NewBaseRepository[model.Book](db)
	id := uuid.New()

	mock.ExpectBegin()

	mock.ExpectExec(`DELETE FROM "books"`).
		WithArgs(id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	err := repo.Delete(id)
	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}
