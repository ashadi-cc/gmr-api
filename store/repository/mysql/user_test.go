package mysql

import (
	"api-gmr/store/repository"
	"context"
	"database/sql"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func newMockDB() (*sql.DB, sqlmock.Sqlmock, error) {
	return sqlmock.New()
}

func Test_userRepo_FindByUsername(t *testing.T) {
	type args struct {
		username string
	}

	type fields struct {
		query string
		rows  *sqlmock.Rows
	}

	type wants struct {
		userId   int
		username string
		email    string
		password string
		blok     string
		name     string
	}

	tests := []struct {
		name       string
		args       args
		fields     fields
		wants      wants
		assertUser assert.ErrorAssertionFunc
		assertMock assert.ErrorAssertionFunc
	}{
		{
			name: "valid query",
			args: args{
				username: "usenamae",
			},
			fields: fields{
				query: "SELECT id,username,email,password,blok,name FROM users WHERE username = ?",
				rows: sqlmock.NewRows([]string{"id", "username", "email", "password", "blok", "name"}).AddRow(
					1, "username", "email", "password", "blok", "name",
				),
			},
			wants: wants{
				userId:   1,
				username: "username",
				email:    "email",
				password: "password",
				blok:     "blok",
				name:     "name",
			},
			assertUser: assert.NoError,
			assertMock: assert.NoError,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := newMockDB()
			if err != nil {
				t.Skip(err.Error())
			}
			defer db.Close()

			mock.ExpectPrepare(tt.fields.query)
			mock.ExpectQuery(tt.fields.query).WithArgs(tt.args.username).WillReturnRows(tt.fields.rows)

			user := userRepo{db: db}
			u, err := user.FindByUsername(context.Background(), tt.args.username)

			tt.assertMock(t, mock.ExpectationsWereMet())

			if tt.assertUser(t, err) {
				assert.Equal(t, tt.wants.userId, u.GetUserID())
				assert.Equal(t, tt.wants.username, u.GetUsername())
				assert.Equal(t, tt.wants.email, u.GetEmail())
				assert.Equal(t, tt.wants.password, u.GetPasswordHash())
				assert.Equal(t, tt.wants.blok, u.GetBlok())
				assert.Equal(t, tt.wants.name, u.GetName())
			}
		})
	}

}

func Test_userRepo_FindByUserID(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx    context.Context
		userID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    repository.UserModel
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := userRepo{
				db: tt.fields.db,
			}
			got, err := repo.FindByUserID(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("userRepo.FindByUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userRepo.FindByUserID() = %v, want %v", got, tt.want)
			}
		})
	}
}
