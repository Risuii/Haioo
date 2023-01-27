// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	filter "github.com/Risuii/models/filter"
	mock "github.com/stretchr/testify/mock"

	product "github.com/Risuii/models/product"
)

// CartRepository is an autogenerated mock type for the CartRepository type
type CartRepository struct {
	mock.Mock
}

// Add provides a mock function with given fields: ctx, params
func (_m *CartRepository) Add(ctx context.Context, params product.Product) (int64, error) {
	ret := _m.Called(ctx, params)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, product.Product) int64); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, product.Product) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *CartRepository) Delete(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAll provides a mock function with given fields:
func (_m *CartRepository) FindAll() ([]product.Product, error) {
	ret := _m.Called()

	var r0 []product.Product
	if rf, ok := ret.Get(0).(func() []product.Product); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]product.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByFilter provides a mock function with given fields: ctx, params
func (_m *CartRepository) FindByFilter(ctx context.Context, params filter.Filter) ([]product.Product, error) {
	ret := _m.Called(ctx, params)

	var r0 []product.Product
	if rf, ok := ret.Get(0).(func(context.Context, filter.Filter) []product.Product); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]product.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, filter.Filter) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByKodeProduk provides a mock function with given fields: ctx, kodeProduk
func (_m *CartRepository) FindByKodeProduk(ctx context.Context, kodeProduk string) (product.Product, error) {
	ret := _m.Called(ctx, kodeProduk)

	var r0 product.Product
	if rf, ok := ret.Get(0).(func(context.Context, string) product.Product); ok {
		r0 = rf(ctx, kodeProduk)
	} else {
		r0 = ret.Get(0).(product.Product)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, kodeProduk)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateKuantitas provides a mock function with given fields: ctx, id, params
func (_m *CartRepository) UpdateKuantitas(ctx context.Context, id int64, params product.Product) error {
	ret := _m.Called(ctx, id, params)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, product.Product) error); ok {
		r0 = rf(ctx, id, params)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewCartRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewCartRepository creates a new instance of CartRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCartRepository(t mockConstructorTestingTNewCartRepository) *CartRepository {
	mock := &CartRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
