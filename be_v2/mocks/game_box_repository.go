// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	model "ewallet-server-v2/internal/model"
	mock "github.com/stretchr/testify/mock"
)

// GameBoxRepository is an autogenerated mock type for the GameBoxRepository type
type GameBoxRepository struct {
	mock.Mock
}

// GetAll provides a mock function with given fields: ctx, limit
func (_m *GameBoxRepository) GetAll(ctx context.Context, limit int) ([]model.GameBox, error) {
	ret := _m.Called(ctx, limit)

	var r0 []model.GameBox
	if rf, ok := ret.Get(0).(func(context.Context, int) []model.GameBox); ok {
		r0 = rf(ctx, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.GameBox)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOneById provides a mock function with given fields: ctx, boxId
func (_m *GameBoxRepository) GetOneById(ctx context.Context, boxId int64) (*model.GameBox, error) {
	ret := _m.Called(ctx, boxId)

	var r0 *model.GameBox
	if rf, ok := ret.Get(0).(func(context.Context, int64) *model.GameBox); ok {
		r0 = rf(ctx, boxId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.GameBox)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, boxId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewGameBoxRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewGameBoxRepository creates a new instance of GameBoxRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGameBoxRepository(t mockConstructorTestingTNewGameBoxRepository) *GameBoxRepository {
	mock := &GameBoxRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}