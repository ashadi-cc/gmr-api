package mysql

import (
	"api-gmr/store/repository"
	"context"
	"database/sql"
	"database/sql/driver"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func newMockDB() (*sql.DB, sqlmock.Sqlmock, error) {
	return sqlmock.New()
}

type userModelTest struct {
	userId   int
	userName string
	email    string
	password string
	group    string
	blok     string
	name     string
}

func (u userModelTest) GetUserID() int {
	return u.userId
}

func (u userModelTest) GetUsername() string {
	return u.userName
}

func (u userModelTest) GetEmail() string {
	return u.email
}

func (u userModelTest) GetGroup() string {
	return u.group
}

func (u userModelTest) GetPasswordHash() string {
	return u.password
}

func (u userModelTest) GetBlok() string {
	return u.blok
}

func (u userModelTest) GetName() string {
	return u.name
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
			name: "returns user",
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
		}, {
			name: "returns empty user",
			args: args{
				username: "usenamae",
			},
			fields: fields{
				query: "SELECT id,username,email,password,blok,name FROM users WHERE username = ?",
				rows:  sqlmock.NewRows([]string{"id", "username", "email", "password", "blok", "name"}),
			},
			assertUser: assert.Error,
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

			prep := mock.ExpectPrepare(tt.fields.query)
			prep.ExpectQuery().WithArgs(tt.args.username).WillReturnRows(tt.fields.rows)

			user := userRepo{db: db}
			u, err := user.FindByUsername(context.Background(), tt.args.username)
			tt.assertUser(t, err)

			tt.assertMock(t, mock.ExpectationsWereMet())

			if err == nil {
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
	type args struct {
		userid int
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
				userid: 1,
			},
			fields: fields{
				query: "SELECT id,username,email,password,blok,name FROM users WHERE id = ?",
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

			prep := mock.ExpectPrepare(tt.fields.query)
			prep.ExpectQuery().WithArgs(tt.args.userid).WillReturnRows(tt.fields.rows)

			user := userRepo{db: db}
			u, err := user.FindByUserID(context.Background(), tt.args.userid)

			tt.assertUser(t, err)
			tt.assertMock(t, mock.ExpectationsWereMet())

			if err == nil {
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

func Test_userRepo_UpdateEmailandPassword(t *testing.T) {
	type fields struct {
		query string
	}
	type args struct {
		user      repository.UserModel
		queryArgs []driver.Value
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		assertUser assert.ErrorAssertionFunc
		assertMock assert.ErrorAssertionFunc
	}{
		{
			name:   "update only email",
			fields: fields{"UPDATE users SET email = \\? WHERE id = \\?"},
			args: args{
				user:      userModelTest{userId: 1, email: "aa@gmail.com"},
				queryArgs: []driver.Value{"aa@gmail.com", 1},
			},
			assertUser: assert.NoError,
			assertMock: assert.NoError,
		},
		{
			name:   "update only email + password",
			fields: fields{"UPDATE users SET email = \\?,password = \\? WHERE id = \\?"},
			args: args{
				user:      userModelTest{userId: 1, email: "aa@gmail.com", password: "test123"},
				queryArgs: []driver.Value{"aa@gmail.com", "test123", 1},
			},
			assertUser: assert.NoError,
			assertMock: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := newMockDB()
			if err != nil {
				t.Skip(err.Error())
			}
			defer db.Close()

			prep := mock.ExpectPrepare(tt.fields.query)
			prep.ExpectExec().WithArgs(tt.args.queryArgs...).WillReturnResult(sqlmock.NewResult(0, 1))

			user := userRepo{db: db}
			err = user.UpdateEmailandPassword(context.Background(), tt.args.user)
			tt.assertUser(t, err)
			tt.assertMock(t, mock.ExpectationsWereMet())
		})
	}
}
