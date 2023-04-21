// Code generated by mockery v2.16.0. DO NOT EDIT.

package dao

import (
	context "context"

	models "github.com/pubgolf/pubgolf/api/internal/lib/models"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// MockQueryProvider is an autogenerated mock type for the QueryProvider type
type MockQueryProvider struct {
	mock.Mock
}

// AdjustmentsByPlayerStage provides a mock function with given fields: ctx, playerID, stageID
func (_m *MockQueryProvider) AdjustmentsByPlayerStage(ctx context.Context, playerID models.PlayerID, stageID models.StageID) ([]models.Adjustment, error) {
	ret := _m.Called(ctx, playerID, stageID)

	var r0 []models.Adjustment
	if rf, ok := ret.Get(0).(func(context.Context, models.PlayerID, models.StageID) []models.Adjustment); ok {
		r0 = rf(ctx, playerID, stageID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Adjustment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.PlayerID, models.StageID) error); ok {
		r1 = rf(ctx, playerID, stageID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreatePlayer provides a mock function with given fields: ctx, eventID, player
func (_m *MockQueryProvider) CreatePlayer(ctx context.Context, eventID models.EventID, player models.PlayerParams) (models.Player, error) {
	ret := _m.Called(ctx, eventID, player)

	var r0 models.Player
	if rf, ok := ret.Get(0).(func(context.Context, models.EventID, models.PlayerParams) models.Player); ok {
		r0 = rf(ctx, eventID, player)
	} else {
		r0 = ret.Get(0).(models.Player)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.EventID, models.PlayerParams) error); ok {
		r1 = rf(ctx, eventID, player)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateScoreForStage provides a mock function with given fields: ctx, playerID, stageID, score, adjustments
func (_m *MockQueryProvider) CreateScoreForStage(ctx context.Context, playerID models.PlayerID, stageID models.StageID, score uint32, adjustments []models.AdjustmentParams) error {
	ret := _m.Called(ctx, playerID, stageID, score, adjustments)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.PlayerID, models.StageID, uint32, []models.AdjustmentParams) error); ok {
		r0 = rf(ctx, playerID, stageID, score, adjustments)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EventIDByKey provides a mock function with given fields: ctx, key
func (_m *MockQueryProvider) EventIDByKey(ctx context.Context, key string) (models.EventID, error) {
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
func (_m *MockQueryProvider) EventPlayers(ctx context.Context, eventID models.EventID) ([]models.Player, error) {
	ret := _m.Called(ctx, eventID)

	var r0 []models.Player
	if rf, ok := ret.Get(0).(func(context.Context, models.EventID) []models.Player); ok {
		r0 = rf(ctx, eventID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Player)
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
func (_m *MockQueryProvider) EventSchedule(ctx context.Context, eventID models.EventID) ([]VenueStop, error) {
	ret := _m.Called(ctx, eventID)

	var r0 []VenueStop
	if rf, ok := ret.Get(0).(func(context.Context, models.EventID) []VenueStop); ok {
		r0 = rf(ctx, eventID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]VenueStop)
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

// EventScheduleCacheVersion provides a mock function with given fields: ctx, eventID, hash
func (_m *MockQueryProvider) EventScheduleCacheVersion(ctx context.Context, eventID models.EventID, hash []byte) (uint32, bool, error) {
	ret := _m.Called(ctx, eventID, hash)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(context.Context, models.EventID, []byte) uint32); ok {
		r0 = rf(ctx, eventID, hash)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(context.Context, models.EventID, []byte) bool); ok {
		r1 = rf(ctx, eventID, hash)
	} else {
		r1 = ret.Get(1).(bool)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, models.EventID, []byte) error); ok {
		r2 = rf(ctx, eventID, hash)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// EventScheduleWithDetails provides a mock function with given fields: ctx, eventID
func (_m *MockQueryProvider) EventScheduleWithDetails(ctx context.Context, eventID models.EventID) ([]models.Stage, error) {
	ret := _m.Called(ctx, eventID)

	var r0 []models.Stage
	if rf, ok := ret.Get(0).(func(context.Context, models.EventID) []models.Stage); ok {
		r0 = rf(ctx, eventID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Stage)
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
func (_m *MockQueryProvider) EventStartTime(ctx context.Context, id models.EventID) (time.Time, error) {
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

// ScoreByPlayerStage provides a mock function with given fields: ctx, playerID, stageID
func (_m *MockQueryProvider) ScoreByPlayerStage(ctx context.Context, playerID models.PlayerID, stageID models.StageID) (models.Score, error) {
	ret := _m.Called(ctx, playerID, stageID)

	var r0 models.Score
	if rf, ok := ret.Get(0).(func(context.Context, models.PlayerID, models.StageID) models.Score); ok {
		r0 = rf(ctx, playerID, stageID)
	} else {
		r0 = ret.Get(0).(models.Score)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.PlayerID, models.StageID) error); ok {
		r1 = rf(ctx, playerID, stageID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePlayer provides a mock function with given fields: ctx, playerID, player
func (_m *MockQueryProvider) UpdatePlayer(ctx context.Context, playerID models.PlayerID, player models.PlayerParams) (models.Player, error) {
	ret := _m.Called(ctx, playerID, player)

	var r0 models.Player
	if rf, ok := ret.Get(0).(func(context.Context, models.PlayerID, models.PlayerParams) models.Player); ok {
		r0 = rf(ctx, playerID, player)
	} else {
		r0 = ret.Get(0).(models.Player)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.PlayerID, models.PlayerParams) error); ok {
		r1 = rf(ctx, playerID, player)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VenueByKey provides a mock function with given fields: ctx, eventID, venueKey
func (_m *MockQueryProvider) VenueByKey(ctx context.Context, eventID models.EventID, venueKey models.VenueKey) (Venue, error) {
	ret := _m.Called(ctx, eventID, venueKey)

	var r0 Venue
	if rf, ok := ret.Get(0).(func(context.Context, models.EventID, models.VenueKey) Venue); ok {
		r0 = rf(ctx, eventID, venueKey)
	} else {
		r0 = ret.Get(0).(Venue)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.EventID, models.VenueKey) error); ok {
		r1 = rf(ctx, eventID, venueKey)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockQueryProvider interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockQueryProvider creates a new instance of MockQueryProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockQueryProvider(t mockConstructorTestingTNewMockQueryProvider) *MockQueryProvider {
	mock := &MockQueryProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
