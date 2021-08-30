package repository

import "context"

//BillingModel represents Billing model interface
type BillingModel interface {
	//GetID returns billing id
	GetID() int
	//GetName returns billing name
	GetName() string
	//GetStatus returns billing status
	GetStatus() string
	//GetAmount returns billing amount
	GetAmount() float64
	//GetYear returns billing year
	GetYear() int
	//GetMonth returns billing month
	GetMonth() int
}

//BillingFilter represents billing filter interface
type BillingFilter interface {
	//GetYear return year filter
	GetYear() int
	//GetMonth return month filter
	GetMonth() int
	//GetUserID returns user id filter
	GetUserID() int
	//GetStatus returns status filter
	GetStatus() string
}

//Billing represents billing methods interface
type Billing interface {
	//GetBillWithFilter returns billings by given filter payload
	GetBillWithFilter(ctx context.Context, filter BillingFilter) ([]BillingModel, error)
	//GetOtherBillWithFilter returns billings by given filter payload
	GetOtherBillWithFilter(ctx context.Context, userId, year, month int) ([]BillingModel, error)
	//StoreBillingFile stored billing image by given userid
	StoreBillingFile(ctx context.Context, userId int, driver, fileURL, description string) error
}
