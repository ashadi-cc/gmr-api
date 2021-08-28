package mysql

import (
	"api-gmr/store/repository"
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

type billingRepo struct {
	db *sql.DB
}

//NewBillingRepo returns a new mysql repository.Billing instance
func NewBillingRepo() repository.Billing {
	return &billingRepo{
		db: getDB(),
	}
}

func (repo billingRepo) selectFields() []string {
	fields := []string{
		"id",
		"billing_name",
		"status",
		"amount",
		"year",
		"month",
	}

	return fields
}

//GetBillWithFilter implementing repository.Billing.GetBillWithFilter
func (repo billingRepo) GetBillWithFilter(ctx context.Context, filter repository.BillingFilter) ([]repository.BillingModel, error) {
	f, args := buildBillingFilter(filter)
	whereClause := ""
	if len(f) > 0 {
		whereClause = fmt.Sprintf("WHERE %s", strings.Join(f, " AND "))
	}

	query := fmt.Sprintf("SELECT %s FROM billing_users %s ORDER BY billing_name",
		strings.Join(repo.selectFields(), ","),
		whereClause,
	)

	statement, err := repo.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, errors.Wrapf(err, "failed preparing query: %s", query)
	}
	defer statement.Close()

	rows, err := statement.QueryContext(ctx, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "failed execute query: %s", query)
	}

	var bs billingModels
	for rows.Next() {
		var b billingModel
		if err := rows.Scan(&b.id, &b.bilingName, &b.status, &b.amount, &b.year, &b.month); err != nil {
			return nil, errors.Wrap(err, "failed scanning rows")
		}
		bs = append(bs, b)
	}

	return bs, nil
}

//GetOtherBillWithFilter implementing repository.Billing.GetOtherBillWithFilter
func (repo billingRepo) GetOtherBillWithFilter(ctx context.Context, filter repository.BillingFilter) ([]repository.BillingModel, error) {
	return nil, nil
}
