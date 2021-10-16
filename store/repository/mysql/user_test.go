package mysql

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func newMockDB() (*sql.DB, sqlmock.Sqlmock, error) {
	return sqlmock.New()
}

func TestUserFindByUsername(t *testing.T) {
	db, mock, err := newMockDB()
	if err != nil {
		t.Skip(err.Error())
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "username", "email", "password", "blok", "name"}).AddRow(
		1, "username", "email", "password", "blok", "name",
	)

	sqlQuery := "SELECT id,username,email,password,blok,name FROM users WHERE username = ?"
	mock.ExpectPrepare(sqlQuery)
	mock.ExpectQuery(sqlQuery).WithArgs("username").WillReturnRows(rows)

	user := &userRepo{db: db}

	u, err := user.FindByUsername(context.Background(), "username")
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
	assert.Equal(t, 1, u.GetUserID())
	assert.Equal(t, "username", u.GetUsername())
	assert.Equal(t, "email", u.GetEmail())
	assert.Equal(t, "password", u.GetPasswordHash())
	assert.Equal(t, "blok", u.GetBlok())
	assert.Equal(t, "name", u.GetName())
}
