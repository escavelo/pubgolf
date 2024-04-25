// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: leaderboard.sql

package dbc

import (
	"context"

	"github.com/pubgolf/pubgolf/api/internal/lib/models"
)

const scoringCriteriaAllVenues = `-- name: ScoringCriteriaAllVenues :many
WITH separated AS ((
    -- Get score contributions.
    SELECT
      p.id AS player_id,
      p.name,
      coalesce(count(DISTINCT (s.id)), 0)::bigint AS num_scores,
      coalesce(sum(s.value), 0)::bigint AS total_points,
      0 AS points_from_penalties,
      0 AS points_from_bonuses
    FROM
      players p
    LEFT JOIN event_players ep ON p.id = ep.player_id
    LEFT JOIN scores s ON s.player_id = p.id
      AND s.is_verified IS TRUE
    LEFT JOIN stages st ON s.stage_id = st.id
      AND st.event_id = ep.event_id
  WHERE
    p.deleted_at IS NULL
    AND ep.deleted_at IS NULL
    AND ep.event_id = $1
    AND ep.scoring_category = $2
    AND s.deleted_at IS NULL
    AND st.deleted_at IS NULL
  GROUP BY
    p.id,
    p.name)
UNION (
  -- Get adjustment contributions.
  SELECT
    p.id AS player_id,
    p.name,
    coalesce(count(DISTINCT (s.id)), 0)::bigint AS num_scores,
    coalesce(sum(a.value), 0)::bigint AS total_points,
    sum(
      CASE WHEN a.value > 0 THEN
        a.value
      ELSE
        0
      END) AS points_from_penalties,
    sum(
      CASE WHEN a.value < 0 THEN
        a.value
      ELSE
        0
      END) AS points_from_bonuses
  FROM
    players p
    LEFT JOIN event_players ep ON p.id = ep.player_id
    LEFT JOIN scores s ON s.player_id = p.id
      AND s.is_verified IS TRUE
    LEFT JOIN stages st ON s.stage_id = st.id
      AND st.event_id = ep.event_id
    LEFT JOIN adjustments a ON a.stage_id = s.stage_id
      AND a.player_id = s.player_id
  WHERE
    p.deleted_at IS NULL
    AND ep.deleted_at IS NULL
    AND ep.event_id = $1
    AND ep.scoring_category = $2
    AND s.deleted_at IS NULL
    AND st.deleted_at IS NULL
    AND a.deleted_at IS NULL
  GROUP BY
    p.id,
    p.name))
SELECT
  player_id,
  name,
  num_scores,
  SUM(total_points) AS total_points,
  SUM(points_from_penalties) AS points_from_penalties,
  SUM(points_from_bonuses) AS points_from_bonuses
FROM
  separated
GROUP BY
  player_id,
  name,
  num_scores
ORDER BY
  num_scores DESC,
  total_points ASC,
  points_from_penalties ASC,
  points_from_bonuses DESC
`

type ScoringCriteriaAllVenuesParams struct {
	EventID         models.EventID
	ScoringCategory models.ScoringCategory
}

type ScoringCriteriaAllVenuesRow struct {
	PlayerID            models.DatabaseULID
	Name                string
	NumScores           int64
	TotalPoints         int64
	PointsFromPenalties int64
	PointsFromBonuses   int64
}

func (q *Queries) ScoringCriteriaAllVenues(ctx context.Context, arg ScoringCriteriaAllVenuesParams) ([]ScoringCriteriaAllVenuesRow, error) {
	rows, err := q.query(ctx, q.scoringCriteriaAllVenuesStmt, scoringCriteriaAllVenues, arg.EventID, arg.ScoringCategory)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ScoringCriteriaAllVenuesRow
	for rows.Next() {
		var i ScoringCriteriaAllVenuesRow
		if err := rows.Scan(
			&i.PlayerID,
			&i.Name,
			&i.NumScores,
			&i.TotalPoints,
			&i.PointsFromPenalties,
			&i.PointsFromBonuses,
		); err != nil {
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

const scoringCriteriaEveryOtherVenue = `-- name: ScoringCriteriaEveryOtherVenue :many
WITH st AS (
  -- Replaces the stages table in the later section with only the odd numbered stages.
  SELECT
    st.event_id, st.venue_id, st.venue_key, st.rank, st.duration_minutes, st.created_at, st.updated_at, st.deleted_at, st.rule_id, st.id,
    mod(row_number() OVER (ORDER BY st.rank ASC), 2) = 1 AS is_odd
  FROM
    stages st
  WHERE
    st.deleted_at IS NULL
    AND st.event_id = $1
),
separated AS ((
    -- Get score contributions.
    SELECT
      p.id AS player_id,
      p.name,
      coalesce(count(DISTINCT (s.id)), 0)::bigint AS num_scores,
      coalesce(sum(s.value), 0)::bigint AS total_points,
      0 AS points_from_penalties,
      0 AS points_from_bonuses
    FROM
      players p
    LEFT JOIN event_players ep ON p.id = ep.player_id
    LEFT JOIN scores s ON s.player_id = p.id
      AND s.is_verified IS TRUE
    LEFT JOIN st ON s.stage_id = st.id
      AND st.event_id = ep.event_id
  WHERE
    p.deleted_at IS NULL
    AND ep.deleted_at IS NULL
    AND ep.event_id = $1
    AND ep.scoring_category = $2
    AND s.deleted_at IS NULL
    AND st.deleted_at IS NULL
    AND st.is_odd
  GROUP BY
    p.id,
    p.name)
UNION (
  -- Get adjustment contributions.
  SELECT
    p.id AS player_id,
    p.name,
    coalesce(count(DISTINCT (s.id)), 0)::bigint AS num_scores,
    coalesce(sum(a.value), 0)::bigint AS total_points,
    sum(
      CASE WHEN a.value > 0 THEN
        a.value
      ELSE
        0
      END) AS points_from_penalties,
    sum(
      CASE WHEN a.value < 0 THEN
        a.value
      ELSE
        0
      END) AS points_from_bonuses
  FROM
    players p
    LEFT JOIN event_players ep ON p.id = ep.player_id
    LEFT JOIN scores s ON s.player_id = p.id
      AND s.is_verified IS TRUE
    LEFT JOIN st ON s.stage_id = st.id
      AND st.event_id = ep.event_id
    LEFT JOIN adjustments a ON a.stage_id = s.stage_id
      AND a.player_id = s.player_id
  WHERE
    p.deleted_at IS NULL
    AND ep.deleted_at IS NULL
    AND ep.event_id = $1
    AND ep.scoring_category = $2
    AND s.deleted_at IS NULL
    AND st.deleted_at IS NULL
    AND st.is_odd
    AND a.deleted_at IS NULL
  GROUP BY
    p.id,
    p.name))
SELECT
  player_id,
  name,
  num_scores,
  SUM(total_points) AS total_points,
  SUM(points_from_penalties) AS points_from_penalties,
  SUM(points_from_bonuses) AS points_from_bonuses
FROM
  separated
GROUP BY
  player_id,
  name,
  num_scores
ORDER BY
  num_scores DESC,
  total_points ASC,
  points_from_penalties ASC,
  points_from_bonuses DESC
`

type ScoringCriteriaEveryOtherVenueParams struct {
	EventID         models.EventID
	ScoringCategory models.ScoringCategory
}

type ScoringCriteriaEveryOtherVenueRow struct {
	PlayerID            models.DatabaseULID
	Name                string
	NumScores           int64
	TotalPoints         int64
	PointsFromPenalties int64
	PointsFromBonuses   int64
}

func (q *Queries) ScoringCriteriaEveryOtherVenue(ctx context.Context, arg ScoringCriteriaEveryOtherVenueParams) ([]ScoringCriteriaEveryOtherVenueRow, error) {
	rows, err := q.query(ctx, q.scoringCriteriaEveryOtherVenueStmt, scoringCriteriaEveryOtherVenue, arg.EventID, arg.ScoringCategory)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ScoringCriteriaEveryOtherVenueRow
	for rows.Next() {
		var i ScoringCriteriaEveryOtherVenueRow
		if err := rows.Scan(
			&i.PlayerID,
			&i.Name,
			&i.NumScores,
			&i.TotalPoints,
			&i.PointsFromPenalties,
			&i.PointsFromBonuses,
		); err != nil {
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

const unverifiedScoreCountAllVenues = `-- name: UnverifiedScoreCountAllVenues :many
SELECT
  p.id AS player_id,
  SUM(
    CASE WHEN s.is_verified IS FALSE THEN
      1
    ELSE
      0
    END) AS count
FROM
  players p
  JOIN event_players ep ON p.id = ep.player_id
  JOIN stages st ON st.event_id = ep.event_id
  LEFT JOIN scores s ON s.player_id = p.id
    AND s.stage_id = st.id
WHERE
  p.deleted_at IS NULL
  AND ep.deleted_at IS NULL
  AND ep.event_id = $1
  AND ep.scoring_category = $2
  AND s.deleted_at IS NULL
  AND st.deleted_at IS NULL
GROUP BY
  p.id
`

type UnverifiedScoreCountAllVenuesParams struct {
	EventID         models.EventID
	ScoringCategory models.ScoringCategory
}

type UnverifiedScoreCountAllVenuesRow struct {
	PlayerID models.PlayerID
	Count    int64
}

func (q *Queries) UnverifiedScoreCountAllVenues(ctx context.Context, arg UnverifiedScoreCountAllVenuesParams) ([]UnverifiedScoreCountAllVenuesRow, error) {
	rows, err := q.query(ctx, q.unverifiedScoreCountAllVenuesStmt, unverifiedScoreCountAllVenues, arg.EventID, arg.ScoringCategory)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UnverifiedScoreCountAllVenuesRow
	for rows.Next() {
		var i UnverifiedScoreCountAllVenuesRow
		if err := rows.Scan(&i.PlayerID, &i.Count); err != nil {
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

const unverifiedScoreCountEveryOtherVenue = `-- name: UnverifiedScoreCountEveryOtherVenue :many
WITH st AS (
  -- Replaces the stages table in the later section with only the odd numbered stages.
  SELECT
    st.event_id, st.venue_id, st.venue_key, st.rank, st.duration_minutes, st.created_at, st.updated_at, st.deleted_at, st.rule_id, st.id,
    mod(row_number() OVER (ORDER BY st.rank ASC), 2) = 1 AS is_odd
  FROM
    stages st
  WHERE
    st.deleted_at IS NULL
    AND st.event_id = $1
)
SELECT
  p.id AS player_id,
  SUM(
    CASE WHEN s.is_verified IS FALSE THEN
      1
    ELSE
      0
    END) AS count
FROM
  players p
  JOIN event_players ep ON p.id = ep.player_id
  JOIN st ON st.event_id = ep.event_id
  LEFT JOIN scores s ON s.player_id = p.id
    AND s.stage_id = st.id
WHERE
  p.deleted_at IS NULL
  AND ep.deleted_at IS NULL
  AND ep.event_id = $1
  AND ep.scoring_category = $2
  AND s.deleted_at IS NULL
  AND st.deleted_at IS NULL
  AND st.is_odd
GROUP BY
  p.id
`

type UnverifiedScoreCountEveryOtherVenueParams struct {
	EventID         models.EventID
	ScoringCategory models.ScoringCategory
}

type UnverifiedScoreCountEveryOtherVenueRow struct {
	PlayerID models.PlayerID
	Count    int64
}

func (q *Queries) UnverifiedScoreCountEveryOtherVenue(ctx context.Context, arg UnverifiedScoreCountEveryOtherVenueParams) ([]UnverifiedScoreCountEveryOtherVenueRow, error) {
	rows, err := q.query(ctx, q.unverifiedScoreCountEveryOtherVenueStmt, unverifiedScoreCountEveryOtherVenue, arg.EventID, arg.ScoringCategory)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UnverifiedScoreCountEveryOtherVenueRow
	for rows.Next() {
		var i UnverifiedScoreCountEveryOtherVenueRow
		if err := rows.Scan(&i.PlayerID, &i.Count); err != nil {
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
