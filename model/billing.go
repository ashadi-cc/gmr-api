package model

import (
	"api-gmr/config"
	"api-gmr/store/repository"
	"fmt"
)

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
	Status string  `json:"status,omitempty"`
	Amount float64 `json:"amount"`
	Year   int     `json:"year,omitempty"`
	Month  int     `json:"month,omitempty"`
}

func (b Billing) Display() Billing {
	b.Status = ""
	b.Year = 0
	b.Month = 0
	return b
}

type Billings []Billing

func (b Billings) Display() Billings {
	for idx, i := range b {
		b[idx] = i.Display()
	}
	return b
}

func (b Billings) TotalAmount() float64 {
	var total float64
	for _, i := range b {
		total = total + i.Amount
	}
	return total
}

type ItemBilling struct {
	Data  Billings `json:"items"`
	Total float64  `json:"total"`
}

type BillingPayment struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type BillingPayments []BillingPayment

type BillingInfo struct {
	ThisMonth     ItemBilling     `json:"this_month"`
	OtherBill     ItemBilling     `json:"other_month"`
	PaymentMethod BillingPayments `json:"payment_method"`
}

func BillRepoToBilling(i []repository.BillingModel) Billings {
	var b Billings
	for _, item := range i {
		bi := Billing{
			Name:   item.GetName(),
			Status: item.GetStatus(),
			Amount: item.GetAmount(),
			Year:   item.GetYear(),
			Month:  item.GetMonth(),
		}
		b = append(b, bi)
	}

	return b
}

func PaymentRepoToPayments(p []repository.PaymentModel) BillingPayments {
	var ps BillingPayments
	for _, item := range p {
		pi := BillingPayment{
			Name:     item.GetName(),
			ImageURL: paymentImageUrl(item.GetImage()),
		}
		ps = append(ps, pi)
	}
	return ps
}

func paymentImageUrl(img string) string {
	if len(img) == 0 {
		return ""
	}
	return fmt.Sprintf("%s/%s", config.GetBaseQrCodeURL(), img)
}
