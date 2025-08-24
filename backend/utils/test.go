package utils

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewMockDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock, func()) {
	sqlDB, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
	require.NoError(t, err)

	dialector := postgres.New(postgres.Config{
		Conn: sqlDB,
	})

	gdb, err := gorm.Open(dialector, &gorm.Config{})
	require.NoError(t, err)

	cleanup := func() { sqlDB.Close() }
	return gdb, mock, cleanup
}
