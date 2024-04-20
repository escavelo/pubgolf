// Code generated by mockery v2.42.2. DO NOT EDIT.

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

// AdjustmentTemplatesByEventID provides a mock function with given fields: ctx, eventID
func (_m *MockQueryProvider) AdjustmentTemplatesByEventID(ctx context.Context, eventID models.EventID) ([]models.AdjustmentTemplate, error) {
	ret := _m.Called(ctx, eventID)

	if len(ret) == 0 {
		panic("no return value specified for AdjustmentTemplatesByEventID")
	}

	var r0 []models.AdjustmentTemplate
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.EventID) ([]models.AdjustmentTemplate, error)); ok {
		return rf(ctx, eventID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.EventID) []models.AdjustmentTemplate); ok {
		r0 = rf(ctx, eventID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.AdjustmentTemplate)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.EventID) error); ok {
		r1 = rf(ctx, eventID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AdjustmentTemplatesByStageID provides a mock function with given fields: ctx, stageID
func (_m *MockQueryProvider) AdjustmentTemplatesByStageID(ctx context.Context, stageID models.StageID) ([]models.AdjustmentTemplate, error) {
	ret := _m.Called(ctx, stageID)

	if len(ret) == 0 {
		panic("no return value specified for AdjustmentTemplatesByStageID")
	}

	var r0 []models.AdjustmentTemplate
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.StageID) ([]models.AdjustmentTemplate, error)); ok {
		return rf(ctx, stageID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.StageID) []models.AdjustmentTemplate); ok {
		r0 = rf(ctx, stageID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.AdjustmentTemplate)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.StageID) error); ok {
		r1 = rf(ctx, stageID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AdjustmentsByPlayerStage provides a mock function with given fields: ctx, playerID, stageID
func (_m *MockQueryProvider) AdjustmentsByPlayerStage(ctx context.Context, playerID models.PlayerID, stageID models.StageID) ([]models.Adjustment, error) {
	ret := _m.Called(ctx, playerID, stageID)

	if len(ret) == 0 {
		panic("no return value specified for AdjustmentsByPlayerStage")
	}

	var r0 []models.Adjustment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.PlayerID, models.StageID) ([]models.Adjustment, error)); ok {
		return rf(ctx, playerID, stageID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.PlayerID, models.StageID) []models.Adjustment); ok {
		r0 = rf(ctx, playerID, stageID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Adjustment)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.PlayerID, models.StageID) error); ok {
		r1 = rf(ctx, playerID, stageID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreatePlayer provides a mock function with given fields: ctx, name, phoneNum
func (_m *MockQueryProvider) CreatePlayer(ctx context.Context, name string, phoneNum models.PhoneNum) (models.Player, error) {
	ret := _m.Called(ctx, name, phoneNum)

	if len(ret) == 0 {
		panic("no return value specified for CreatePlayer")
	}

	var r0 models.Player
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, models.PhoneNum) (models.Player, error)); ok {
		return rf(ctx, name, phoneNum)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, models.PhoneNum) models.Player); ok {
		r0 = rf(ctx, name, phoneNum)
	} else {
		r0 = ret.Get(0).(models.Player)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, models.PhoneNum) error); ok {
		r1 = rf(ctx, name, phoneNum)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteScore provides a mock function with given fields: ctx, playerID, stageID
func (_m *MockQueryProvider) DeleteScore(ctx context.Context, playerID models.PlayerID, stageID models.StageID) error {
	ret := _m.Called(ctx, playerID, stageID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteScore")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.PlayerID, models.StageID) error); ok {
		r0 = rf(ctx, playerID, stageID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EventIDByKey provides a mock function with given fields: ctx, key
func (_m *MockQueryProvider) EventIDByKey(ctx context.Context, key string) (models.EventID, error) {
	ret := _m.Called(ctx, key)

	if len(ret) == 0 {
		panic("no return value specified for EventIDByKey")
	}

	var r0 models.EventID
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (models.EventID, error)); ok {
		return rf(ctx, key)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) models.EventID); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Get(0).(models.EventID)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EventPlayers provides a mock function with given fields: ctx, eventKey
func (_m *MockQueryProvider) EventPlayers(ctx context.Context, eventKey string) ([]models.Player, error) {
	ret := _m.Called(ctx, eventKey)

	if len(ret) == 0 {
		panic("no return value specified for EventPlayers")
	}

	var r0 []models.Player
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]models.Player, error)); ok {
		return rf(ctx, eventKey)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []models.Player); ok {
		r0 = rf(ctx, eventKey)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Player)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, eventKey)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EventSchedule provides a mock function with given fields: ctx, eventID
func (_m *MockQueryProvider) EventSchedule(ctx context.Context, eventID models.EventID) ([]VenueStop, error) {
	ret := _m.Called(ctx, eventID)

	if len(ret) == 0 {
		panic("no return value specified for EventSchedule")
	}

	var r0 []VenueStop
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.EventID) ([]VenueStop, error)); ok {
		return rf(ctx, eventID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.EventID) []VenueStop); ok {
		r0 = rf(ctx, eventID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]VenueStop)
		}
	}

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

	if len(ret) == 0 {
		panic("no return value specified for EventScheduleCacheVersion")
	}

	var r0 uint32
	var r1 bool
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, models.EventID, []byte) (uint32, bool, error)); ok {
		return rf(ctx, eventID, hash)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.EventID, []byte) uint32); ok {
		r0 = rf(ctx, eventID, hash)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.EventID, []byte) bool); ok {
		r1 = rf(ctx, eventID, hash)
	} else {
		r1 = ret.Get(1).(bool)
	}

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

	if len(ret) == 0 {
		panic("no return value specified for EventScheduleWithDetails")
	}

	var r0 []models.Stage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.EventID) ([]models.Stage, error)); ok {
		return rf(ctx, eventID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.EventID) []models.Stage); ok {
		r0 = rf(ctx, eventID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Stage)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.EventID) error); ok {
		r1 = rf(ctx, eventID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EventScores provides a mock function with given fields: ctx, eventID
func (_m *MockQueryProvider) EventScores(ctx context.Context, eventID models.EventID) ([]models.StageScore, error) {
	ret := _m.Called(ctx, eventID)

	if len(ret) == 0 {
		panic("no return value specified for EventScores")
	}

	var r0 []models.StageScore
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.EventID) ([]models.StageScore, error)); ok {
		return rf(ctx, eventID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.EventID) []models.StageScore); ok {
		r0 = rf(ctx, eventID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.StageScore)
		}
	}

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

	if len(ret) == 0 {
		panic("no return value specified for EventStartTime")
	}

	var r0 time.Time
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.EventID) (time.Time, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.EventID) time.Time); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.EventID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GenerateAuthToken provides a mock function with given fields: ctx, num
func (_m *MockQueryProvider) GenerateAuthToken(ctx context.Context, num models.PhoneNum) (GenerateAuthTokenResult, error) {
	ret := _m.Called(ctx, num)

	if len(ret) == 0 {
		panic("no return value specified for GenerateAuthToken")
	}

	var r0 GenerateAuthTokenResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.PhoneNum) (GenerateAuthTokenResult, error)); ok {
		return rf(ctx, num)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.PhoneNum) GenerateAuthTokenResult); ok {
		r0 = rf(ctx, num)
	} else {
		r0 = ret.Get(0).(GenerateAuthTokenResult)
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.PhoneNum) error); ok {
		r1 = rf(ctx, num)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PhoneNumberIsVerified provides a mock function with given fields: ctx, num
func (_m *MockQueryProvider) PhoneNumberIsVerified(ctx context.Context, num models.PhoneNum) (bool, error) {
	ret := _m.Called(ctx, num)

	if len(ret) == 0 {
		panic("no return value specified for PhoneNumberIsVerified")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.PhoneNum) (bool, error)); ok {
		return rf(ctx, num)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.PhoneNum) bool); ok {
		r0 = rf(ctx, num)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.PhoneNum) error); ok {
		r1 = rf(ctx, num)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PlayerAdjustments provides a mock function with given fields: ctx, eventID, playerID
func (_m *MockQueryProvider) PlayerAdjustments(ctx context.Context, eventID models.EventID, playerID models.PlayerID) ([]PlayerVenueAdjustment, error) {
	ret := _m.Called(ctx, eventID, playerID)

	if len(ret) == 0 {
		panic("no return value specified for PlayerAdjustments")
	}

	var r0 []PlayerVenueAdjustment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.EventID, models.PlayerID) ([]PlayerVenueAdjustment, error)); ok {
		return rf(ctx, eventID, playerID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.EventID, models.PlayerID) []PlayerVenueAdjustment); ok {
		r0 = rf(ctx, eventID, playerID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]PlayerVenueAdjustment)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.EventID, models.PlayerID) error); ok {
		r1 = rf(ctx, eventID, playerID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PlayerByID provides a mock function with given fields: ctx, playerID
func (_m *MockQueryProvider) PlayerByID(ctx context.Context, playerID models.PlayerID) (models.Player, error) {
	ret := _m.Called(ctx, playerID)

	if len(ret) == 0 {
		panic("no return value specified for PlayerByID")
	}

	var r0 models.Player
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.PlayerID) (models.Player, error)); ok {
		return rf(ctx, playerID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.PlayerID) models.Player); ok {
		r0 = rf(ctx, playerID)
	} else {
		r0 = ret.Get(0).(models.Player)
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.PlayerID) error); ok {
		r1 = rf(ctx, playerID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PlayerIDByAuthToken provides a mock function with given fields: ctx, token
func (_m *MockQueryProvider) PlayerIDByAuthToken(ctx context.Context, token models.AuthToken) (models.PlayerID, error) {
	ret := _m.Called(ctx, token)

	if len(ret) == 0 {
		panic("no return value specified for PlayerIDByAuthToken")
	}

	var r0 models.PlayerID
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.AuthToken) (models.PlayerID, error)); ok {
		return rf(ctx, token)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.AuthToken) models.PlayerID); ok {
		r0 = rf(ctx, token)
	} else {
		r0 = ret.Get(0).(models.PlayerID)
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.AuthToken) error); ok {
		r1 = rf(ctx, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PlayerRegisteredForEvent provides a mock function with given fields: ctx, playerID, eventID
func (_m *MockQueryProvider) PlayerRegisteredForEvent(ctx context.Context, playerID models.PlayerID, eventID models.EventID) (bool, error) {
	ret := _m.Called(ctx, playerID, eventID)

	if len(ret) == 0 {
		panic("no return value specified for PlayerRegisteredForEvent")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.PlayerID, models.EventID) (bool, error)); ok {
		return rf(ctx, playerID, eventID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.PlayerID, models.EventID) bool); ok {
		r0 = rf(ctx, playerID, eventID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.PlayerID, models.EventID) error); ok {
		r1 = rf(ctx, playerID, eventID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PlayerScores provides a mock function with given fields: ctx, eventID, playerID, includeUnverified
func (_m *MockQueryProvider) PlayerScores(ctx context.Context, eventID models.EventID, playerID models.PlayerID, includeUnverified bool) ([]PlayerVenueScore, error) {
	ret := _m.Called(ctx, eventID, playerID, includeUnverified)

	if len(ret) == 0 {
		panic("no return value specified for PlayerScores")
	}

	var r0 []PlayerVenueScore
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.EventID, models.PlayerID, bool) ([]PlayerVenueScore, error)); ok {
		return rf(ctx, eventID, playerID, includeUnverified)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.EventID, models.PlayerID, bool) []PlayerVenueScore); ok {
		r0 = rf(ctx, eventID, playerID, includeUnverified)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]PlayerVenueScore)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.EventID, models.PlayerID, bool) error); ok {
		r1 = rf(ctx, eventID, playerID, includeUnverified)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ScoreByPlayerStage provides a mock function with given fields: ctx, playerID, stageID
func (_m *MockQueryProvider) ScoreByPlayerStage(ctx context.Context, playerID models.PlayerID, stageID models.StageID) (models.Score, error) {
	ret := _m.Called(ctx, playerID, stageID)

	if len(ret) == 0 {
		panic("no return value specified for ScoreByPlayerStage")
	}

	var r0 models.Score
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.PlayerID, models.StageID) (models.Score, error)); ok {
		return rf(ctx, playerID, stageID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.PlayerID, models.StageID) models.Score); ok {
		r0 = rf(ctx, playerID, stageID)
	} else {
		r0 = ret.Get(0).(models.Score)
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.PlayerID, models.StageID) error); ok {
		r1 = rf(ctx, playerID, stageID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ScoringCriteria provides a mock function with given fields: ctx, eventID, category
func (_m *MockQueryProvider) ScoringCriteria(ctx context.Context, eventID models.EventID, category models.ScoringCategory) ([]models.ScoringInput, error) {
	ret := _m.Called(ctx, eventID, category)

	if len(ret) == 0 {
		panic("no return value specified for ScoringCriteria")
	}

	var r0 []models.ScoringInput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.EventID, models.ScoringCategory) ([]models.ScoringInput, error)); ok {
		return rf(ctx, eventID, category)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.EventID, models.ScoringCategory) []models.ScoringInput); ok {
		r0 = rf(ctx, eventID, category)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.ScoringInput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.EventID, models.ScoringCategory) error); ok {
		r1 = rf(ctx, eventID, category)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StageIDByVenueKey provides a mock function with given fields: ctx, eventID, venueKey
func (_m *MockQueryProvider) StageIDByVenueKey(ctx context.Context, eventID models.EventID, venueKey models.VenueKey) (models.StageID, error) {
	ret := _m.Called(ctx, eventID, venueKey)

	if len(ret) == 0 {
		panic("no return value specified for StageIDByVenueKey")
	}

	var r0 models.StageID
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.EventID, models.VenueKey) (models.StageID, error)); ok {
		return rf(ctx, eventID, venueKey)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.EventID, models.VenueKey) models.StageID); ok {
		r0 = rf(ctx, eventID, venueKey)
	} else {
		r0 = ret.Get(0).(models.StageID)
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.EventID, models.VenueKey) error); ok {
		r1 = rf(ctx, eventID, venueKey)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePlayer provides a mock function with given fields: ctx, playerID, player
func (_m *MockQueryProvider) UpdatePlayer(ctx context.Context, playerID models.PlayerID, player models.PlayerParams) (models.Player, error) {
	ret := _m.Called(ctx, playerID, player)

	if len(ret) == 0 {
		panic("no return value specified for UpdatePlayer")
	}

	var r0 models.Player
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.PlayerID, models.PlayerParams) (models.Player, error)); ok {
		return rf(ctx, playerID, player)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.PlayerID, models.PlayerParams) models.Player); ok {
		r0 = rf(ctx, playerID, player)
	} else {
		r0 = ret.Get(0).(models.Player)
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.PlayerID, models.PlayerParams) error); ok {
		r1 = rf(ctx, playerID, player)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpsertRegistration provides a mock function with given fields: ctx, playerID, eventKey, cat
func (_m *MockQueryProvider) UpsertRegistration(ctx context.Context, playerID models.PlayerID, eventKey string, cat models.ScoringCategory) error {
	ret := _m.Called(ctx, playerID, eventKey, cat)

	if len(ret) == 0 {
		panic("no return value specified for UpsertRegistration")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.PlayerID, string, models.ScoringCategory) error); ok {
		r0 = rf(ctx, playerID, eventKey, cat)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpsertScore provides a mock function with given fields: ctx, playerID, stageID, score, adjustments, isVerified
func (_m *MockQueryProvider) UpsertScore(ctx context.Context, playerID models.PlayerID, stageID models.StageID, score uint32, adjustments []AdjustmentParams, isVerified bool) error {
	ret := _m.Called(ctx, playerID, stageID, score, adjustments, isVerified)

	if len(ret) == 0 {
		panic("no return value specified for UpsertScore")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.PlayerID, models.StageID, uint32, []AdjustmentParams, bool) error); ok {
		r0 = rf(ctx, playerID, stageID, score, adjustments, isVerified)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// VenueByKey provides a mock function with given fields: ctx, eventID, venueKey
func (_m *MockQueryProvider) VenueByKey(ctx context.Context, eventID models.EventID, venueKey models.VenueKey) (Venue, error) {
	ret := _m.Called(ctx, eventID, venueKey)

	if len(ret) == 0 {
		panic("no return value specified for VenueByKey")
	}

	var r0 Venue
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.EventID, models.VenueKey) (Venue, error)); ok {
		return rf(ctx, eventID, venueKey)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.EventID, models.VenueKey) Venue); ok {
		r0 = rf(ctx, eventID, venueKey)
	} else {
		r0 = ret.Get(0).(Venue)
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.EventID, models.VenueKey) error); ok {
		r1 = rf(ctx, eventID, venueKey)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VerifyPhoneNumber provides a mock function with given fields: ctx, num
func (_m *MockQueryProvider) VerifyPhoneNumber(ctx context.Context, num models.PhoneNum) (bool, error) {
	ret := _m.Called(ctx, num)

	if len(ret) == 0 {
		panic("no return value specified for VerifyPhoneNumber")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.PhoneNum) (bool, error)); ok {
		return rf(ctx, num)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.PhoneNum) bool); ok {
		r0 = rf(ctx, num)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.PhoneNum) error); ok {
		r1 = rf(ctx, num)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockQueryProvider creates a new instance of MockQueryProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockQueryProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockQueryProvider {
	mock := &MockQueryProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
