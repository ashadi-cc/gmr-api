package mysql

import (
	"api-gmr/store/repository"
	"context"
	"database/sql/driver"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func Test_billingRepo_GetBillWithFilter(t *testing.T) {
	type fields struct {
		query     string
		rows      *sqlmock.Rows
		queryArgs []driver.Value
	}
	type args struct {
		ctx    context.Context
		filter repository.BillingFilter
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantRows   int
		assertFunc assert.ErrorAssertionFunc
		assertMock assert.ErrorAssertionFunc
	}{
		{
			name: "valid query",
			fields: fields{
				query: "SELECT id,billing_name,status,amount,year,month FROM billing_users WHERE year = \\? AND user_id = \\? ORDER BY billing_name",
				rows: sqlmock.NewRows([]string{"id", "billing_name", "status", "amount", "year", "month"}).AddRow(
					1, "name", "status", 200, 2021, 2,
				),
				queryArgs: []driver.Value{"2021", "1"},
			},
			args: args{
				ctx:    context.Background(),
				filter: filterTest{userID: 1, year: 2021},
			},
			wantRows:   1,
			assertFunc: assert.NoError,
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
			prep.ExpectQuery().WithArgs(tt.fields.queryArgs...).WillReturnRows(tt.fields.rows)

			repo := billingRepo{db: db}
			p, err := repo.GetBillWithFilter(tt.args.ctx, tt.args.filter)
			tt.assertFunc(t, err)
			tt.assertMock(t, mock.ExpectationsWereMet())
			assert.Equal(t, tt.wantRows, len(p))
		})
	}
}

func Test_billingRepo_StoreBillingFile(t *testing.T) {
	//skip now
	t.Skip()
	type fields struct {
		query string
	}
	type args struct {
		ctx         context.Context
		userId      int
		driver      string
		fileURL     string
		description string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		assertFunc assert.ErrorAssertionFunc
		assertMock assert.ErrorAssertionFunc
	}{
		{
			name: "valid insert query",
			fields: fields{
				query: "INSERT INTO billing_files \\(user_id, driver, file_url, description, created_at\\) VALUES(\\?, \\?, \\?, \\?\\, NOW())",
			},
			args: args{
				ctx:         context.Background(),
				userId:      1,
				driver:      "file",
				fileURL:     "/media/test.jpg",
				description: "test upload",
			},
			assertFunc: assert.NoError,
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
			prep.ExpectExec().WithArgs(
				tt.args.userId,
				tt.args.driver,
				tt.args.fileURL,
				tt.args.description,
			).WillReturnResult(sqlmock.NewResult(1, 0))

			repo := billingRepo{db: db}
			err = repo.StoreBillingFile(tt.args.ctx, tt.args.userId, tt.args.driver, tt.args.fileURL, tt.args.description)

			tt.assertMock(t, mock.ExpectationsWereMet())
			tt.assertFunc(t, err)
		})
	}
}
