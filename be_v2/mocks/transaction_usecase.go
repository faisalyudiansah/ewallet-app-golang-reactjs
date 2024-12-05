// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	appdto "ewallet-server-v2/internal/dto/appdto"

	decimal "github.com/shopspring/decimal"

	mock "github.com/stretchr/testify/mock"

	model "ewallet-server-v2/internal/model"

	pagedto "ewallet-server-v2/internal/dto/pagedto"

	time "time"
)

// TransactionUsecase is an autogenerated mock type for the TransactionUsecase type
type TransactionUsecase struct {
	mock.Mock
}

// GetExpenseSumByMonth provides a mock function with given fields: ctx, walletId, month
func (_m *TransactionUsecase) GetExpenseSumByMonth(ctx context.Context, walletId int64, month time.Month) (*appdto.TransactionSum, error) {
	ret := _m.Called(ctx, walletId, month)

	var r0 *appdto.TransactionSum
	if rf, ok := ret.Get(0).(func(context.Context, int64, time.Month) *appdto.TransactionSum); ok {
		r0 = rf(ctx, walletId, month)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appdto.TransactionSum)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64, time.Month) error); ok {
		r1 = rf(ctx, walletId, month)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetListByWalletId provides a mock function with given fields: ctx, walletId, pageDto
func (_m *TransactionUsecase) GetListByWalletId(ctx context.Context, walletId int64, pageDto pagedto.PageSortDto) (*appdto.TransactionListDto, error) {
	ret := _m.Called(ctx, walletId, pageDto)

	var r0 *appdto.TransactionListDto
	if rf, ok := ret.Get(0).(func(context.Context, int64, pagedto.PageSortDto) *appdto.TransactionListDto); ok {
		r0 = rf(ctx, walletId, pageDto)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appdto.TransactionListDto)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64, pagedto.PageSortDto) error); ok {
		r1 = rf(ctx, walletId, pageDto)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetThisMonthExpenseSum provides a mock function with given fields: ctx, walletId, date
func (_m *TransactionUsecase) GetThisMonthExpenseSum(ctx context.Context, walletId int64, date time.Time) (*appdto.TransactionSum, error) {
	ret := _m.Called(ctx, walletId, date)

	var r0 *appdto.TransactionSum
	if rf, ok := ret.Get(0).(func(context.Context, int64, time.Time) *appdto.TransactionSum); ok {
		r0 = rf(ctx, walletId, date)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appdto.TransactionSum)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64, time.Time) error); ok {
		r1 = rf(ctx, walletId, date)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransactionType provides a mock function with given fields: ctx
func (_m *TransactionUsecase) GetTransactionType(ctx context.Context) ([]model.TransactionType, error) {
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

// GetTransactionTypeSumMultiple provides a mock function with given fields: ctx, walletId, transactionTypeId, transactionAdditionalDetailId, minAmount
func (_m *TransactionUsecase) GetTransactionTypeSumMultiple(ctx context.Context, walletId int64, transactionTypeId int64, transactionAdditionalDetailId []int64, minAmount decimal.Decimal) (*decimal.Decimal, error) {
	ret := _m.Called(ctx, walletId, transactionTypeId, transactionAdditionalDetailId, minAmount)

	var r0 *decimal.Decimal
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64, []int64, decimal.Decimal) *decimal.Decimal); ok {
		r0 = rf(ctx, walletId, transactionTypeId, transactionAdditionalDetailId, minAmount)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*decimal.Decimal)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64, int64, []int64, decimal.Decimal) error); ok {
		r1 = rf(ctx, walletId, transactionTypeId, transactionAdditionalDetailId, minAmount)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Topup provides a mock function with given fields: ctx, walletId, amount, sourceOfFundId
func (_m *TransactionUsecase) Topup(ctx context.Context, walletId int64, amount decimal.Decimal, sourceOfFundId int64) (*model.Transaction, error) {
	ret := _m.Called(ctx, walletId, amount, sourceOfFundId)

	var r0 *model.Transaction
	if rf, ok := ret.Get(0).(func(context.Context, int64, decimal.Decimal, int64) *model.Transaction); ok {
		r0 = rf(ctx, walletId, amount, sourceOfFundId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64, decimal.Decimal, int64) error); ok {
		r1 = rf(ctx, walletId, amount, sourceOfFundId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Transfer provides a mock function with given fields: ctx, fromId, to, amount, description
func (_m *TransactionUsecase) Transfer(ctx context.Context, fromId int64, to string, amount decimal.Decimal, description string) (*model.Transaction, error) {
	ret := _m.Called(ctx, fromId, to, amount, description)

	var r0 *model.Transaction
	if rf, ok := ret.Get(0).(func(context.Context, int64, string, decimal.Decimal, string) *model.Transaction); ok {
		r0 = rf(ctx, fromId, to, amount, description)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64, string, decimal.Decimal, string) error); ok {
		r1 = rf(ctx, fromId, to, amount, description)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewTransactionUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewTransactionUsecase creates a new instance of TransactionUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTransactionUsecase(t mockConstructorTestingTNewTransactionUsecase) *TransactionUsecase {
	mock := &TransactionUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}