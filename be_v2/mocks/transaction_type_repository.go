// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	model "ewallet-server-v2/internal/model"
	mock "github.com/stretchr/testify/mock"
)

// TransactionTypeRepository is an autogenerated mock type for the TransactionTypeRepository type
type TransactionTypeRepository struct {
	mock.Mock
}

// GetAll provides a mock function with given fields: ctx
func (_m *TransactionTypeRepository) GetAll(ctx context.Context) ([]model.TransactionType, error) {
	ret := _m.Called(ctx)

	var r0 []model.TransactionType
	if rf, ok := ret.Get(0).(func(context.Context) []model.TransactionType); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.TransactionType)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewTransactionTypeRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewTransactionTypeRepository creates a new instance of TransactionTypeRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTransactionTypeRepository(t mockConstructorTestingTNewTransactionTypeRepository) *TransactionTypeRepository {
	mock := &TransactionTypeRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}