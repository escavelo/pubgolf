BEGIN
;


DROP INDEX
  IF EXISTS event_venues_by_venue_id;


DROP TABLE
  IF EXISTS event_venues;


DROP TABLE
  IF EXISTS venues;


DROP TABLE
  IF EXISTS events;


DROP FUNCTION
  IF EXISTS generate_ulid;


DROP EXTENSION
  IF EXISTS pgcrypto;


COMMIT
;
