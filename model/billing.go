package model

import "api-gmr/store/repository"

type BillingFilter struct {
	Year   int
	Month  int
	UserID int
	Status string
}

func (f BillingFilter) GetYear() int {
	return f.Year
}

func (f BillingFilter) GetMonth() int {
	return f.Month
}

func (f BillingFilter) GetUserID() int {
	return f.UserID
}

func (f BillingFilter) GetStatus() string {
	return f.Status
}

type Billing struct {
	Name   string  `json:"name"`
	Status string  `json:"status"`
	Amount float64 `json:"amount"`
}

type BillingInfo struct {
	ThisMonth []Billing `json:"this_month"`
	Total     float64   `json:"total"`
}

func BillRepoToBilling(i []repository.BillingModel) []Billing {
	var b []Billing
	for _, item := range i {
		bi := Billing{
			Name:   item.GetName(),
			Status: item.GetStatus(),
			Amount: item.GetAmount(),
		}
		b = append(b, bi)
	}

	return b
}
