package mysql

import (
	"api-gmr/store/repository"
)

type billingModel struct {
	id         int
	bilingName string
	status     string
	amount     float64
	year       int
	month      int
}

func (b billingModel) GetID() int {
	return b.id
}

func (b billingModel) GetName() string {
	return b.bilingName
}

func (b billingModel) GetStatus() string {
	return b.status
}

func (b billingModel) GetAmount() float64 {
	return b.amount
}

func (b billingModel) GetYear() int {
	return b.year
}

func (b billingModel) GetMonth() int {
	return b.month
}

type billingModels []repository.BillingModel
