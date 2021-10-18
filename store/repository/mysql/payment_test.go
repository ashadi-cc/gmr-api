package mysql

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func Test_paymentRepo_All(t *testing.T) {
	type fields struct {
		query string
		rows  *sqlmock.Rows
	}
	type args struct {
		ctx context.Context
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
				query: "SELECT id, name, qr_code FROM payments ORDER BY id",
				rows:  sqlmock.NewRows([]string{"id", "name", "qr_code"}).AddRow(1, "name", "image"),
			},
			args:       args{context.Background()},
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
			prep.ExpectQuery().WillReturnRows(tt.fields.rows)

			repo := paymentRepo{db: db}
			p, err := repo.All(tt.args.ctx)
			tt.assertFunc(t, err)
			tt.assertMock(t, mock.ExpectationsWereMet())
			assert.Equal(t, tt.wantRows, len(p))

		})
	}
}
