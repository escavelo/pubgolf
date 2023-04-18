// Code generated by mockery v2.16.0. DO NOT EDIT.

package dbc

import (
	context "context"

	models "github.com/pubgolf/pubgolf/api/internal/lib/models"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// MockQuerier is an autogenerated mock type for the Querier type
type MockQuerier struct {
	mock.Mock
}

// CreatePlayer provides a mock function with given fields: ctx, arg
func (_m *MockQuerier) CreatePlayer(ctx context.Context, arg CreatePlayerParams) (CreatePlayerRow, error) {
	ret := _m.Called(ctx, arg)

	var r0 CreatePlayerRow
	if rf, ok := ret.Get(0).(func(context.Context, CreatePlayerParams) CreatePlayerRow); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(CreatePlayerRow)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, CreatePlayerParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EventCacheVersionByHash provides a mock function with given fields: ctx, arg
func (_m *MockQuerier) EventCacheVersionByHash(ctx context.Context, arg EventCacheVersionByHashParams) (uint32, error) {
	ret := _m.Called(ctx, arg)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(context.Context, EventCacheVersionByHashParams) uint32); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, EventCacheVersionByHashParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EventIDByKey provides a mock function with given fields: ctx, key
func (_m *MockQuerier) EventIDByKey(ctx context.Context, key string) (models.EventID, error) {
	ret := _m.Called(ctx, key)

	var r0 models.EventID
	if rf, ok := ret.Get(0).(func(context.Context, string) models.EventID); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Get(0).(models.EventID)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EventPlayers provides a mock function with given fields: ctx, eventID
func (_m *MockQuerier) EventPlayers(ctx context.Context, eventID models.EventID) ([]EventPlayersRow, error) {
	ret := _m.Called(ctx, eventID)

	var r0 []EventPlayersRow
	if rf, ok := ret.Get(0).(func(context.Context, models.EventID) []EventPlayersRow); ok {
		r0 = rf(ctx, eventID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]EventPlayersRow)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.EventID) error); ok {
		r1 = rf(ctx, eventID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EventSchedule provides a mock function with given fields: ctx, eventID
func (_m *MockQuerier) EventSchedule(ctx context.Context, eventID models.EventID) ([]EventScheduleRow, error) {
	ret := _m.Called(ctx, eventID)

	var r0 []EventScheduleRow
	if rf, ok := ret.Get(0).(func(context.Context, models.EventID) []EventScheduleRow); ok {
		r0 = rf(ctx, eventID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]EventScheduleRow)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.EventID) error); ok {
		r1 = rf(ctx, eventID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EventStartTime provides a mock function with given fields: ctx, id
func (_m *MockQuerier) EventStartTime(ctx context.Context, id models.EventID) (time.Time, error) {
	ret := _m.Called(ctx, id)

	var r0 time.Time
	if rf, ok := ret.Get(0).(func(context.Context, models.EventID) time.Time); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.EventID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EventVenueKeysAreValid provides a mock function with given fields: ctx, eventID
func (_m *MockQuerier) EventVenueKeysAreValid(ctx context.Context, eventID models.EventID) (bool, error) {
	ret := _m.Called(ctx, eventID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, models.EventID) bool); ok {
		r0 = rf(ctx, eventID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.EventID) error); ok {
		r1 = rf(ctx, eventID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetEventCacheKeys provides a mock function with given fields: ctx, arg
func (_m *MockQuerier) SetEventCacheKeys(ctx context.Context, arg SetEventCacheKeysParams) (uint32, error) {
	ret := _m.Called(ctx, arg)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(context.Context, SetEventCacheKeysParams) uint32); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, SetEventCacheKeysParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetEventVenueKeys provides a mock function with given fields: ctx, eventID
func (_m *MockQuerier) SetEventVenueKeys(ctx context.Context, eventID models.EventID) error {
	ret := _m.Called(ctx, eventID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.EventID) error); ok {
		r0 = rf(ctx, eventID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetNextEventVenueKey provides a mock function with given fields: ctx, id
func (_m *MockQuerier) SetNextEventVenueKey(ctx context.Context, id models.EventID) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.EventID) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdatePlayer provides a mock function with given fields: ctx, arg
func (_m *MockQuerier) UpdatePlayer(ctx context.Context, arg UpdatePlayerParams) (UpdatePlayerRow, error) {
	ret := _m.Called(ctx, arg)

	var r0 UpdatePlayerRow
	if rf, ok := ret.Get(0).(func(context.Context, UpdatePlayerParams) UpdatePlayerRow); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(UpdatePlayerRow)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, UpdatePlayerParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VenueByKey provides a mock function with given fields: ctx, arg
func (_m *MockQuerier) VenueByKey(ctx context.Context, arg VenueByKeyParams) (VenueByKeyRow, error) {
	ret := _m.Called(ctx, arg)

	var r0 VenueByKeyRow
	if rf, ok := ret.Get(0).(func(context.Context, VenueByKeyParams) VenueByKeyRow); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(VenueByKeyRow)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, VenueByKeyParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockQuerier interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockQuerier creates a new instance of MockQuerier. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockQuerier(t mockConstructorTestingTNewMockQuerier) *MockQuerier {
	mock := &MockQuerier{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
