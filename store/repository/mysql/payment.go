package mysql

import (
	"api-gmr/store/repository"
	"context"
	"database/sql"

	"github.com/pkg/errors"
)

type paymentRepo struct {
	db *sql.DB
}

//NewUserRepo returns a new instance of repository.Payment
func NewPaymentRepo() repository.Payment {
	return &paymentRepo{
		db: getDB(),
	}
}

//All returns all payment
func (repo paymentRepo) All(ctx context.Context) ([]repository.PaymentModel, error) {
	query := "SELECT id, name, qr_code FROM payments ORDER BY id"
	statement, err := repo.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, errors.Wrapf(err, "failed preparing query: %s", query)
	}
	defer statement.Close()

	rows, err := statement.QueryContext(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "failed execute query: %s", query)
	}

	var payments paymentModels
	for rows.Next() {
		var p paymentModel
		if err := rows.Scan(&p.id, &p.name, &p.qrCode); err != nil {
			return nil, errors.Wrap(err, "failed scanning rows")
		}
		payments = append(payments, p)
	}

	return payments, nil
}
