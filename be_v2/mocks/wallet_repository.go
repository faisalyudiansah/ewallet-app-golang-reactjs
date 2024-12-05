// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	model "ewallet-server-v2/internal/model"
	mock "github.com/stretchr/testify/mock"
)

// WalletRepository is an autogenerated mock type for the WalletRepository type
type WalletRepository struct {
	mock.Mock
}

// CreateOne provides a mock function with given fields: ctx, wallet
func (_m *WalletRepository) CreateOne(ctx context.Context, wallet model.Wallet) (*model.Wallet, error) {
	ret := _m.Called(ctx, wallet)

	var r0 *model.Wallet
	if rf, ok := ret.Get(0).(func(context.Context, model.Wallet) *model.Wallet); ok {
		r0 = rf(ctx, wallet)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Wallet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.Wallet) error); ok {
		r1 = rf(ctx, wallet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOneByIdWithLock provides a mock function with given fields: ctx, walletId
func (_m *WalletRepository) GetOneByIdWithLock(ctx context.Context, walletId int64) (*model.Wallet, error) {
	ret := _m.Called(ctx, walletId)

	var r0 *model.Wallet
	if rf, ok := ret.Get(0).(func(context.Context, int64) *model.Wallet); ok {
		r0 = rf(ctx, walletId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Wallet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, walletId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOneByNumber provides a mock function with given fields: ctx, walletNumber
func (_m *WalletRepository) GetOneByNumber(ctx context.Context, walletNumber string) (*model.Wallet, error) {
	ret := _m.Called(ctx, walletNumber)

	var r0 *model.Wallet
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.Wallet); ok {
		r0 = rf(ctx, walletNumber)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Wallet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, walletNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOneByNumberWithLock provides a mock function with given fields: ctx, walletNumber
func (_m *WalletRepository) GetOneByNumberWithLock(ctx context.Context, walletNumber string) (*model.Wallet, error) {
	ret := _m.Called(ctx, walletNumber)

	var r0 *model.Wallet
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.Wallet); ok {
		r0 = rf(ctx, walletNumber)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Wallet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, walletNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOneByUserId provides a mock function with given fields: ctx, userId
func (_m *WalletRepository) GetOneByUserId(ctx context.Context, userId int64) (*model.Wallet, error) {
	ret := _m.Called(ctx, userId)

	var r0 *model.Wallet
	if rf, ok := ret.Get(0).(func(context.Context, int64) *model.Wallet); ok {
		r0 = rf(ctx, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Wallet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveOne provides a mock function with given fields: ctx, wallet
func (_m *WalletRepository) SaveOne(ctx context.Context, wallet model.Wallet) (*model.Wallet, error) {
	ret := _m.Called(ctx, wallet)

	var r0 *model.Wallet
	if rf, ok := ret.Get(0).(func(context.Context, model.Wallet) *model.Wallet); ok {
		r0 = rf(ctx, wallet)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Wallet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.Wallet) error); ok {
		r1 = rf(ctx, wallet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewWalletRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewWalletRepository creates a new instance of WalletRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewWalletRepository(t mockConstructorTestingTNewWalletRepository) *WalletRepository {
	mock := &WalletRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}