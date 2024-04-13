// // Code generated by ifacemaker; DO NOT EDIT.

package dao

import (
	"context"
	"time"

	"github.com/pubgolf/pubgolf/api/internal/lib/models"
)

// QueryProvider describes all of the queries exposed by the DAO, to allow for testing mocks.
type QueryProvider interface {
	// AdjustmentsByPlayerStage returns the base score for a given player/stage combination.
	AdjustmentsByPlayerStage(ctx context.Context, playerID models.PlayerID, stageID models.StageID) ([]models.Adjustment, error)
	// CreatePlayer creates a new player.
	CreatePlayer(ctx context.Context, name string, phoneNum models.PhoneNum) (models.Player, error)
	// CreatePlayerAndRegistration creates a new player and adds them to the given event.
	CreatePlayerAndRegistration(ctx context.Context, name string, phoneNum models.PhoneNum, eventID models.EventID, cat models.ScoringCategory) (models.Player, error)
	// CreateScoreForStage creates score and adjustment records for a given stage.
	CreateScoreForStage(ctx context.Context, playerID models.PlayerID, stageID models.StageID, score uint32, adjustments []models.AdjustmentParams) error
	// DeleteScore creates score and adjustment records for a given stage.
	DeleteScore(ctx context.Context, playerID models.PlayerID, stageID models.StageID) error
	// EventIDByKey takes a human readable event key (slug) and returns the event's canonical identifier.
	EventIDByKey(ctx context.Context, key string) (models.EventID, error)
	// EventPlayers returns all players registered for a given event, in alphabetical order by name. The player's event registrations will only include the specified event.
	EventPlayers(ctx context.Context, eventKey string) ([]models.Player, error)
	// EventSchedule returns a slice of venue keys and durations for the given event ID.
	EventSchedule(ctx context.Context, eventID models.EventID) ([]VenueStop, error)
	// EventScheduleCacheVersion returns the integer version number of the latest schedule version, as well as whether or not the provided hash triggered a cache break.
	EventScheduleCacheVersion(ctx context.Context, eventID models.EventID, hash []byte) (v uint32, hashMatched bool, err error)
	// EventScheduleWithDetails returns the event schedule with venue and rule details included.
	EventScheduleWithDetails(ctx context.Context, eventID models.EventID) ([]models.Stage, error)
	// EventScores returns all the scores (and their adjustments) for an event, ordered by stage and then by creation time.
	EventScores(ctx context.Context, eventID models.EventID) ([]models.StageScore, error)
	// EventStartTime returns the start time for the given event ID.
	EventStartTime(ctx context.Context, id models.EventID) (time.Time, error)
	// GenerateAuthToken generates an auth token for the player with the given phone number, returning the auth token and the player's ID.
	GenerateAuthToken(ctx context.Context, num models.PhoneNum) (GenerateAuthTokenResult, error)
	// PhoneNumberIsVerified returns true if the phone number has been verified via an auth code.
	PhoneNumberIsVerified(ctx context.Context, num models.PhoneNum) (bool, error)
	// PlayerAdjustments returns a list of event stages where a player has an adjustment(s) and their labels/values.
	PlayerAdjustments(ctx context.Context, eventID models.EventID, playerID models.PlayerID) ([]PlayerVenueAdjustment, error)
	// PlayerByID returns a player's profile data and event registrations.
	PlayerByID(ctx context.Context, playerID models.PlayerID) (models.Player, error)
	// PlayerIDByAuthToken takes an auth token and returns the corresponding player ID (only if the token is active).
	PlayerIDByAuthToken(ctx context.Context, token models.AuthToken) (models.PlayerID, error)
	// PlayerRegisteredForEvent returns whether or not the player has a valid registration for the given event.
	PlayerRegisteredForEvent(ctx context.Context, playerID models.PlayerID, eventID models.EventID) (bool, error)
	// PlayerScores returns a list of event stages and a player's scoring info for each.
	PlayerScores(ctx context.Context, eventID models.EventID, playerID models.PlayerID) ([]PlayerVenueScore, error)
	// ScoreByPlayerStage returns the base score for a given player/stage combination.
	ScoreByPlayerStage(ctx context.Context, playerID models.PlayerID, stageID models.StageID) (models.Score, error)
	// ScoringCriteria returns a list of players competing in the given category and the data necessary to rank them.
	ScoringCriteria(ctx context.Context, eventID models.EventID, category models.ScoringCategory) ([]models.ScoringInput, error)
	// UpdatePlayer creates a new player and adds them to the given event.
	UpdatePlayer(ctx context.Context, playerID models.PlayerID, player models.PlayerParams) (models.Player, error)
	// UpdateScore creates score and adjustment records for a given stage.
	UpdateScore(ctx context.Context, playerID models.PlayerID, stageID models.StageID, score models.Score, modifyAdj []models.Adjustment, createAdj []models.AdjustmentParams) error
	// UpsertRegistration creates a new player and adds them to the given event.
	UpsertRegistration(ctx context.Context, playerID models.PlayerID, eventKey string, cat models.ScoringCategory) error
	// VenueByKey returns venue info for the venue key and event id.
	VenueByKey(ctx context.Context, eventID models.EventID, venueKey models.VenueKey) (Venue, error)
	// VerifyPhoneNumber sets the player's phone number as verified, returning a boolean to indicate whether the phone number was previously unverified (i.e. false means the DB row was not updated).
	VerifyPhoneNumber(ctx context.Context, num models.PhoneNum) (bool, error)
}
