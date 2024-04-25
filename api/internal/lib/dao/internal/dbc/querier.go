// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package dbc

import (
	"context"
	"time"

	"github.com/pubgolf/pubgolf/api/internal/lib/models"
)

type Querier interface {
	AdjustmentTemplatesByEventID(ctx context.Context, eventID models.EventID) ([]AdjustmentTemplatesByEventIDRow, error)
	AdjustmentTemplatesByStageID(ctx context.Context, stageID models.StageID) ([]AdjustmentTemplatesByStageIDRow, error)
	AdjustmentsByPlayerStage(ctx context.Context, arg AdjustmentsByPlayerStageParams) ([]AdjustmentsByPlayerStageRow, error)
	AllVenues(ctx context.Context) ([]AllVenuesRow, error)
	CreateAdjustment(ctx context.Context, arg CreateAdjustmentParams) error
	CreateAdjustmentTemplate(ctx context.Context, arg CreateAdjustmentTemplateParams) (models.AdjustmentTemplateID, error)
	CreateAdjustmentWithTemplate(ctx context.Context, arg CreateAdjustmentWithTemplateParams) error
	CreatePlayer(ctx context.Context, arg CreatePlayerParams) (models.PlayerID, error)
	DeactivateAuthTokens(ctx context.Context, phoneNumber models.PhoneNum) (bool, error)
	DeleteAdjustment(ctx context.Context, id models.AdjustmentID) error
	DeleteAdjustmentsForPlayerStage(ctx context.Context, arg DeleteAdjustmentsForPlayerStageParams) error
	DeletePlayer(ctx context.Context, id models.PlayerID) error
	DeleteScoreForPlayerStage(ctx context.Context, arg DeleteScoreForPlayerStageParams) error
	EventAdjustmentTemplates(ctx context.Context, eventID models.EventID) ([]EventAdjustmentTemplatesRow, error)
	EventAdjustments(ctx context.Context, eventID models.EventID) ([]EventAdjustmentsRow, error)
	EventCacheVersionByHash(ctx context.Context, arg EventCacheVersionByHashParams) (uint32, error)
	EventIDByKey(ctx context.Context, key string) (models.EventID, error)
	EventPlayers(ctx context.Context, eventKey string) ([]EventPlayersRow, error)
	EventSchedule(ctx context.Context, eventID models.EventID) ([]EventScheduleRow, error)
	EventScheduleWithDetails(ctx context.Context, eventID models.EventID) ([]EventScheduleWithDetailsRow, error)
	EventScores(ctx context.Context, arg EventScoresParams) ([]EventScoresRow, error)
	EventStartTime(ctx context.Context, id models.EventID) (time.Time, error)
	EventVenueKeysAreValid(ctx context.Context, eventID models.EventID) (bool, error)
	GenerateAuthToken(ctx context.Context, phoneNumber models.PhoneNum) (GenerateAuthTokenRow, error)
	PhoneNumberIsVerified(ctx context.Context, phoneNumber models.PhoneNum) (bool, error)
	PlayerAdjustments(ctx context.Context, arg PlayerAdjustmentsParams) ([]PlayerAdjustmentsRow, error)
	PlayerByID(ctx context.Context, id models.PlayerID) (PlayerByIDRow, error)
	PlayerIDByAuthToken(ctx context.Context, authToken models.AuthToken) (models.PlayerID, error)
	PlayerRegisteredForEvent(ctx context.Context, arg PlayerRegisteredForEventParams) (bool, error)
	PlayerRegistrationsByID(ctx context.Context, id models.PlayerID) ([]PlayerRegistrationsByIDRow, error)
	PlayerScores(ctx context.Context, arg PlayerScoresParams) ([]PlayerScoresRow, error)
	ScoreByPlayerStage(ctx context.Context, arg ScoreByPlayerStageParams) (ScoreByPlayerStageRow, error)
	ScoringCriteriaAllVenues(ctx context.Context, arg ScoringCriteriaAllVenuesParams) ([]ScoringCriteriaAllVenuesRow, error)
	ScoringCriteriaEveryOtherVenue(ctx context.Context, arg ScoringCriteriaEveryOtherVenueParams) ([]ScoringCriteriaEveryOtherVenueRow, error)
	SetEventCacheKeys(ctx context.Context, arg SetEventCacheKeysParams) (uint32, error)
	SetEventVenueKeys(ctx context.Context, eventID models.EventID) error
	SetNextEventVenueKey(ctx context.Context, id models.EventID) error
	StageIDByVenueKey(ctx context.Context, arg StageIDByVenueKeyParams) (models.StageID, error)
	UnverifiedScoreCountAllVenues(ctx context.Context, arg UnverifiedScoreCountAllVenuesParams) ([]UnverifiedScoreCountAllVenuesRow, error)
	UnverifiedScoreCountEveryOtherVenue(ctx context.Context, arg UnverifiedScoreCountEveryOtherVenueParams) ([]UnverifiedScoreCountEveryOtherVenueRow, error)
	UpdateAdjustmentTemplate(ctx context.Context, arg UpdateAdjustmentTemplateParams) error
	UpdatePlayer(ctx context.Context, arg UpdatePlayerParams) error
	UpdateRuleByStage(ctx context.Context, arg UpdateRuleByStageParams) error
	UpdateStage(ctx context.Context, arg UpdateStageParams) error
	UpsertRegistration(ctx context.Context, arg UpsertRegistrationParams) error
	UpsertScore(ctx context.Context, arg UpsertScoreParams) error
	VenueByKey(ctx context.Context, arg VenueByKeyParams) (VenueByKeyRow, error)
	VerifyPhoneNumber(ctx context.Context, phoneNumber models.PhoneNum) (bool, error)
}

var _ Querier = (*Queries)(nil)
