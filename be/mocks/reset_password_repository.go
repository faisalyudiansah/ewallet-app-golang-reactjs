// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	models "git.garena.com/sea-labs-id/bootcamp/batch-04/faisal.yudiansah/assignment-e-wallet-rest-api/models"
	mock "github.com/stretchr/testify/mock"
)

// ResetPasswordRepository is an autogenerated mock type for the ResetPasswordRepository type
type ResetPasswordRepository struct {
	mock.Mock
}

// DeleteToken provides a mock function with given fields: _a0, _a1
func (_m *ResetPasswordRepository) DeleteToken(_a0 context.Context, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IsTokenResetValid provides a mock function with given fields: _a0, _a1
func (_m *ResetPasswordRepository) IsTokenResetValid(_a0 context.Context, _a1 string) bool {
	ret := _m.Called(_a0, _a1)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// PostNewDataResetPassword provides a mock function with given fields: _a0, _a1, _a2
func (_m *ResetPasswordRepository) PostNewDataResetPassword(_a0 context.Context, _a1 int64, _a2 string) (*models.ResetPassword, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *models.ResetPassword
	if rf, ok := ret.Get(0).(func(context.Context, int64, string) *models.ResetPassword); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ResetPassword)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64, string) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewResetPasswordRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewResetPasswordRepository creates a new instance of ResetPasswordRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewResetPasswordRepository(t mockConstructorTestingTNewResetPasswordRepository) *ResetPasswordRepository {
	mock := &ResetPasswordRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}