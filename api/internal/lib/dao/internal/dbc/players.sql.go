// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: players.sql

package dbc

import (
	"context"

	"github.com/pubgolf/pubgolf/api/internal/lib/models"
)

const createPlayer = `-- name: CreatePlayer :one
INSERT INTO players(name, phone_number, updated_at)
  VALUES ($1, $2, now())
RETURNING
  id
`

type CreatePlayerParams struct {
	Name        string
	PhoneNumber models.PhoneNum
}

func (q *Queries) CreatePlayer(ctx context.Context, arg CreatePlayerParams) (models.PlayerID, error) {
	row := q.queryRow(ctx, q.createPlayerStmt, createPlayer, arg.Name, arg.PhoneNumber)
	var id models.PlayerID
	err := row.Scan(&id)
	return id, err
}

const eventPlayers = `-- name: EventPlayers :many
SELECT
  p.id,
  p.name,
  ep.scoring_category
FROM
  players p
  JOIN event_players ep ON p.id = ep.player_id
  JOIN events e ON ep.event_id = e.id
WHERE
  e.key = $1
  AND e.deleted_at IS NULL
  AND ep.deleted_at IS NULL
  AND p.deleted_at IS NULL
ORDER BY
  name ASC
`

type EventPlayersRow struct {
	ID              models.PlayerID
	Name            string
	ScoringCategory models.ScoringCategory
}

func (q *Queries) EventPlayers(ctx context.Context, eventKey string) ([]EventPlayersRow, error) {
	rows, err := q.query(ctx, q.eventPlayersStmt, eventPlayers, eventKey)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []EventPlayersRow
	for rows.Next() {
		var i EventPlayersRow
		if err := rows.Scan(&i.ID, &i.Name, &i.ScoringCategory); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const phoneNumberIsVerified = `-- name: PhoneNumberIsVerified :one
SELECT
  phone_number_verified
FROM
  players
WHERE
  deleted_at IS NULL
  AND phone_number = $1
`

func (q *Queries) PhoneNumberIsVerified(ctx context.Context, phoneNumber models.PhoneNum) (bool, error) {
	row := q.queryRow(ctx, q.phoneNumberIsVerifiedStmt, phoneNumberIsVerified, phoneNumber)
	var phone_number_verified bool
	err := row.Scan(&phone_number_verified)
	return phone_number_verified, err
}

const playerByID = `-- name: PlayerByID :one
SELECT
  id,
  name
FROM
  players
WHERE
  id = $1
  AND deleted_at IS NULL
`

type PlayerByIDRow struct {
	ID   models.PlayerID
	Name string
}

func (q *Queries) PlayerByID(ctx context.Context, id models.PlayerID) (PlayerByIDRow, error) {
	row := q.queryRow(ctx, q.playerByIDStmt, playerByID, id)
	var i PlayerByIDRow
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const playerRegistrationsByID = `-- name: PlayerRegistrationsByID :many
SELECT
  e.key AS event_key,
  ep.scoring_category
FROM
  players p
  JOIN event_players ep ON p.id = ep.player_id
  JOIN events e ON ep.event_id = e.id
WHERE
  p.deleted_at IS NULL
  AND p.id = $1
  AND ep.deleted_at IS NULL
  AND e.deleted_at IS NULL
`

type PlayerRegistrationsByIDRow struct {
	EventKey        string
	ScoringCategory models.ScoringCategory
}

func (q *Queries) PlayerRegistrationsByID(ctx context.Context, id models.PlayerID) ([]PlayerRegistrationsByIDRow, error) {
	rows, err := q.query(ctx, q.playerRegistrationsByIDStmt, playerRegistrationsByID, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []PlayerRegistrationsByIDRow
	for rows.Next() {
		var i PlayerRegistrationsByIDRow
		if err := rows.Scan(&i.EventKey, &i.ScoringCategory); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updatePlayer = `-- name: UpdatePlayer :exec
UPDATE
  players
SET
  name = $1,
  updated_at = now()
WHERE
  id = $2
`

type UpdatePlayerParams struct {
	Name string
	ID   models.PlayerID
}

func (q *Queries) UpdatePlayer(ctx context.Context, arg UpdatePlayerParams) error {
	_, err := q.exec(ctx, q.updatePlayerStmt, updatePlayer, arg.Name, arg.ID)
	return err
}

const upsertRegistration = `-- name: UpsertRegistration :exec
INSERT INTO event_players(event_id, player_id, scoring_category, deleted_at)
  VALUES ($1, $2, $3, NULL)
ON CONFLICT (event_id, player_id)
  DO UPDATE SET
    scoring_category = EXCLUDED.scoring_category, deleted_at = NULL
`

type UpsertRegistrationParams struct {
	EventID         models.EventID
	PlayerID        models.PlayerID
	ScoringCategory models.ScoringCategory
}

func (q *Queries) UpsertRegistration(ctx context.Context, arg UpsertRegistrationParams) error {
	_, err := q.exec(ctx, q.upsertRegistrationStmt, upsertRegistration, arg.EventID, arg.PlayerID, arg.ScoringCategory)
	return err
}

const verifyPhoneNumber = `-- name: VerifyPhoneNumber :one
UPDATE
  players
SET
  phone_number_verified = TRUE,
  updated_at = now()
WHERE
  deleted_at IS NULL
  AND phone_number = $1
  AND phone_number_verified = FALSE
RETURNING
  TRUE AS did_update
`

func (q *Queries) VerifyPhoneNumber(ctx context.Context, phoneNumber models.PhoneNum) (bool, error) {
	row := q.queryRow(ctx, q.verifyPhoneNumberStmt, verifyPhoneNumber, phoneNumber)
	var did_update bool
	err := row.Scan(&did_update)
	return did_update, err
}
