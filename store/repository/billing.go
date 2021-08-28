package repository

import "context"

type BillingModel interface {
	GetID() int
	GetName() string
	GetStatus() string
	GetAmount() float64
}

type BillingFilter interface {
	GetYear() int
	GetMonth() int
	GetUserID() int
	GetStatus() string
}

type Billing interface {
	GetBillWithFilter(ctx context.Context, filter BillingFilter) ([]BillingModel, error)
	GetOtherBillWithFilter(ctx context.Context, filter BillingFilter) ([]BillingModel, error)
}
