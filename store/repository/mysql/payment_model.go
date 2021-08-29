package mysql

import "api-gmr/store/repository"

type paymentModel struct {
	id     int
	name   string
	qrCode string
}

func (p paymentModel) GetID() int {
	return p.id
}

func (p paymentModel) GetName() string {
	return p.name
}

func (p paymentModel) GetImage() string {
	return p.qrCode
}

type paymentModels []repository.PaymentModel
