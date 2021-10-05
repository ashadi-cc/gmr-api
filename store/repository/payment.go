//go:generate mockery --name=PaymentModel
//go:generate mockery --name=Payment
package repository

import (
	"context"
)

//PaymentModel base payment method collections interface
type PaymentModel interface {
	//GetID returns id
	GetID() int
	//GetName returns payment name
	GetName() string
	//GetImage returns payment image path
	GetImage() string
}

//Payment payment repository methods collection
type Payment interface {
	//All returns all payments
	All(ctx context.Context) ([]PaymentModel, error)
}
