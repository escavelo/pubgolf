package dbc_test

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/pubgolf/pubgolf/api/internal/lib/models"
)

func setupEvent(ctx context.Context, t *testing.T, tx *sql.Tx) models.EventID {
	t.Helper()

	row := tx.QueryRowContext(ctx, `
		INSERT INTO events 
			(key, starts_at) 
		VALUES 
			($1, now())
		RETURNING id;
	`, faker.Word())
	require.NoError(t, row.Err(), "create fixture data of event")

	var e models.EventID
	require.NoError(t, row.Scan(&e), "scan returned event ID")

	return e
}

func setupVenue(ctx context.Context, t *testing.T, tx *sql.Tx) models.Venue {
	t.Helper()

	name := faker.Word()
	address := faker.Sentence()

	row := tx.QueryRowContext(ctx, `
		INSERT INTO venues 
			(name, address) 
		VALUES 
			($1, $2)
		RETURNING 
			id;
	`, name, address)
	require.NoError(t, row.Err(), "create fixture data of venue")

	var v models.VenueID
	require.NoError(t, row.Scan(&v), "scan returned venue ID")

	return models.Venue{
		ID:       v,
		Name:     name,
		Address:  address,
		ImageURL: "",
	}
}

func TestEventIDByKey(t *testing.T) {
	t.Parallel()

	t.Run("Returns ID for matching event key", func(t *testing.T) {
		t.Parallel()

		ctx, tx, cleanup := initDB(t)
		defer cleanup()

		// Insert fixture data.
		expectedID := models.EventIDFromULID(ulid.Make())
		slug := "my-test-slug"
		_, err := tx.ExecContext(ctx, `
			INSERT INTO events 
				(id, key, starts_at) 
			VALUES 
				($1, $2, now());
		`, expectedID, slug)
		require.NoError(t, err)

		// Run query and assert results.
		gotID, err := _sharedDBC.WithTx(tx).EventIDByKey(ctx, slug)
		require.NoError(t, err)
		assert.Equal(t, expectedID, gotID)
	})

	t.Run("Returns sql.ErrNoRows when no matching event key", func(t *testing.T) {
		t.Parallel()

		ctx, tx, cleanup := initDB(t)
		defer cleanup()

		// Insert fixture data.
		expectedID := models.EventIDFromULID(ulid.Make())
		slug := "my-test-slug"
		_, err := tx.ExecContext(ctx, `
			INSERT INTO events 
				(id, key, starts_at) 
			VALUES 
				($1, $2, now());
		`, expectedID, slug)
		require.NoError(t, err)

		// Run query and assert results.
		_, err = _sharedDBC.WithTx(tx).EventIDByKey(ctx, slug+"-does-not-match")
		assert.ErrorIs(t, err, sql.ErrNoRows)
	})

	t.Run("Does not return event when deleted_at is set", func(t *testing.T) {
		t.Parallel()

		ctx, tx, cleanup := initDB(t)
		defer cleanup()

		// Insert fixture data.
		expectedID := models.EventIDFromULID(ulid.Make())
		slug := "my-test-slug"
		_, err := tx.ExecContext(ctx, `
			INSERT INTO events 
				(id, key, starts_at, deleted_at) 
			VALUES 
				($1, $2, now(), now() - INTERVAL '1 hour');
		`, expectedID, slug)
		require.NoError(t, err)

		// Run query and assert results.
		_, err = _sharedDBC.WithTx(tx).EventIDByKey(ctx, slug)
		assert.ErrorIs(t, err, sql.ErrNoRows)
	})
}

func TestEventVenueKeysAreValid(t *testing.T) {
	t.Parallel()

	venueCounts := []int{1, 3, 9}

	for _, numVenues := range venueCounts {
		t.Run(fmt.Sprintf("Returns true if all event venues have a lookup key (%d venues)", numVenues), func(t *testing.T) {
			t.Parallel()

			ctx, tx, cleanup := initDB(t)
			defer cleanup()

			// Insert fixture data.
			eventID := setupEvent(ctx, t, tx)

			var venues []models.Venue
			for range numVenues {
				venues = append(venues, setupVenue(ctx, t, tx))
			}

			for i, v := range venues {
				_, err := tx.ExecContext(ctx, `
				INSERT INTO stages 
					(event_id, venue_id, duration_minutes, rank, venue_key) 
				VALUES 
					($1, $2, 30, $3, $3);
			`, eventID, v.ID, i)
				require.NoError(t, err)
			}

			// Run query and assert results.
			gotValid, err := _sharedDBC.WithTx(tx).EventVenueKeysAreValid(ctx, eventID)
			require.NoError(t, err)
			assert.True(t, gotValid)
		})

		t.Run(fmt.Sprintf("Returns false if first event venue has a NULL lookup key (%d venues)", numVenues), func(t *testing.T) {
			t.Parallel()

			ctx, tx, cleanup := initDB(t)
			defer cleanup()

			// Insert fixture data.
			eventID := setupEvent(ctx, t, tx)

			var venues []models.Venue
			for range numVenues {
				venues = append(venues, setupVenue(ctx, t, tx))
			}

			for i, v := range venues {
				if i == 0 {
					_, err := tx.ExecContext(ctx, `
					INSERT INTO stages 
						(event_id, venue_id, duration_minutes, rank, venue_key) 
					VALUES 
						($1, $2, 30, $3, NULL);
				`, eventID, v.ID, i)
					require.NoError(t, err)

					continue
				}

				_, err := tx.ExecContext(ctx, `
				INSERT INTO stages 
					(event_id, venue_id, duration_minutes, rank, venue_key) 
				VALUES 
					($1, $2, 30, $3, $3);
			`, eventID, v.ID, i)
				require.NoError(t, err)
			}

			// Run query and assert results.
			gotValid, err := _sharedDBC.WithTx(tx).EventVenueKeysAreValid(ctx, eventID)
			require.NoError(t, err)
			assert.False(t, gotValid)
		})

		t.Run(fmt.Sprintf("Returns false if last event venue has a NULL lookup key (%d venues)", numVenues), func(t *testing.T) {
			ctx, tx, cleanup := initDB(t)
			defer cleanup()

			// Insert fixture data.
			eventID := setupEvent(ctx, t, tx)

			var venues []models.Venue
			for range numVenues {
				venues = append(venues, setupVenue(ctx, t, tx))
			}

			for i, v := range venues {
				if i == len(venues)-1 {
					_, err := tx.ExecContext(ctx, `
					INSERT INTO stages 
						(event_id, venue_id, duration_minutes, rank, venue_key) 
					VALUES 
						($1, $2, 30, $3, NULL);
				`, eventID, v.ID, i)
					require.NoError(t, err)

					continue
				}

				_, err := tx.ExecContext(ctx, `
				INSERT INTO stages 
					(event_id, venue_id, duration_minutes, rank, venue_key) 
				VALUES 
					($1, $2, 30, $3, $3);
			`, eventID, v.ID, i)

				require.NoError(t, err)
			}

			// Run query and assert results.
			gotValid, err := _sharedDBC.WithTx(tx).EventVenueKeysAreValid(ctx, eventID)
			require.NoError(t, err)
			assert.False(t, gotValid)
		})

		t.Run(fmt.Sprintf("Returns false if all event venues have a NULL lookup key (%d venues)", numVenues), func(t *testing.T) {
			ctx, tx, cleanup := initDB(t)
			defer cleanup()

			// Insert fixture data.
			eventID := setupEvent(ctx, t, tx)

			var venues []models.Venue
			for range numVenues {
				venues = append(venues, setupVenue(ctx, t, tx))
			}

			for i, v := range venues {
				_, err := tx.ExecContext(ctx, `
				INSERT INTO stages 
					(event_id, venue_id, duration_minutes, rank, venue_key) 
				VALUES 
					($1, $2, 30, $3, NULL);
			`, eventID, v.ID, i)
				require.NoError(t, err)
			}

			// Run query and assert results.
			gotValid, err := _sharedDBC.WithTx(tx).EventVenueKeysAreValid(ctx, eventID)
			require.NoError(t, err)
			assert.False(t, gotValid)
		})
	}
}
