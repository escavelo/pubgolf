// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package dbc

import (
	"context"
	"time"

	"github.com/pubgolf/pubgolf/api/internal/lib/models"
)

type Querier interface {
	CreatePlayer(ctx context.Context, arg CreatePlayerParams) (CreatePlayerRow, error)
	EventIDByKey(ctx context.Context, key string) (models.EventID, error)
	EventPlayers(ctx context.Context, eventID models.EventID) ([]EventPlayersRow, error)
	EventSchedule(ctx context.Context, eventID models.EventID) ([]EventScheduleRow, error)
	EventStartTime(ctx context.Context, id models.EventID) (time.Time, error)
	EventVenueKeysAreValid(ctx context.Context, eventID models.EventID) (bool, error)
	SetEventVenueKeys(ctx context.Context, eventID models.EventID) error
	SetNextEventVenueKey(ctx context.Context, id models.EventID) error
	UpdatePlayer(ctx context.Context, arg UpdatePlayerParams) (UpdatePlayerRow, error)
	VenueByKey(ctx context.Context, arg VenueByKeyParams) (VenueByKeyRow, error)
}

var _ Querier = (*Queries)(nil)
