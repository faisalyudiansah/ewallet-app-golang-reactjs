// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	dtos "git.garena.com/sea-labs-id/bootcamp/batch-04/faisal.yudiansah/assignment-e-wallet-rest-api/dtos"
	mock "github.com/stretchr/testify/mock"

	models "git.garena.com/sea-labs-id/bootcamp/batch-04/faisal.yudiansah/assignment-e-wallet-rest-api/models"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// GetUserByEmail provides a mock function with given fields: _a0, _a1
func (_m *UserRepository) GetUserByEmail(_a0 context.Context, _a1 string) (*models.User, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.User); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserById provides a mock function with given fields: _a0, _a1
func (_m *UserRepository) GetUserById(_a0 context.Context, _a1 int64) (*models.User, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(context.Context, int64) *models.User); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsEmailAlreadyRegistered provides a mock function with given fields: _a0, _a1
func (_m *UserRepository) IsEmailAlreadyRegistered(_a0 context.Context, _a1 string) bool {
	ret := _m.Called(_a0, _a1)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// PostUser provides a mock function with given fields: _a0, _a1, _a2
func (_m *UserRepository) PostUser(_a0 context.Context, _a1 dtos.RequestRegisterUser, _a2 string) (*models.User, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(context.Context, dtos.RequestRegisterUser, string) *models.User); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, dtos.RequestRegisterUser, string) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutAttemptGame provides a mock function with given fields: _a0, _a1, _a2
func (_m *UserRepository) PutAttemptGame(_a0 context.Context, _a1 int, _a2 int64) (*models.User, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(context.Context, int, int64) *models.User); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, int64) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutResetPassword provides a mock function with given fields: _a0, _a1, _a2
func (_m *UserRepository) PutResetPassword(_a0 context.Context, _a1 string, _a2 int64) (*models.User, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) *models.User); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, int64) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserRepository creates a new instance of UserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserRepository(t mockConstructorTestingTNewUserRepository) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
