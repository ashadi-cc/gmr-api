// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	repository "api-gmr/store/repository"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// Billing is an autogenerated mock type for the Billing type
type Billing struct {
	mock.Mock
}

// GetBillWithFilter provides a mock function with given fields: ctx, filter
func (_m *Billing) GetBillWithFilter(ctx context.Context, filter repository.BillingFilter) ([]repository.BillingModel, error) {
	ret := _m.Called(ctx, filter)

	var r0 []repository.BillingModel
	if rf, ok := ret.Get(0).(func(context.Context, repository.BillingFilter) []repository.BillingModel); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]repository.BillingModel)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, repository.BillingFilter) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOtherBillWithFilter provides a mock function with given fields: ctx, userId, year, month
func (_m *Billing) GetOtherBillWithFilter(ctx context.Context, userId int, year int, month int) ([]repository.BillingModel, error) {
	ret := _m.Called(ctx, userId, year, month)

	var r0 []repository.BillingModel
	if rf, ok := ret.Get(0).(func(context.Context, int, int, int) []repository.BillingModel); ok {
		r0 = rf(ctx, userId, year, month)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]repository.BillingModel)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, int, int) error); ok {
		r1 = rf(ctx, userId, year, month)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StoreBillingFile provides a mock function with given fields: ctx, userId, driver, fileURL, description
func (_m *Billing) StoreBillingFile(ctx context.Context, userId int, driver string, fileURL string, description string) error {
	ret := _m.Called(ctx, userId, driver, fileURL, description)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, string, string, string) error); ok {
		r0 = rf(ctx, userId, driver, fileURL, description)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
