-- +goose Up
-- +goose StatementBegin
ALTER TABLE bookings
  ADD COLUMN datetime TIMESTAMP NOT NULL;

UPDATE bookings
  SET datetime = date + time;

ALTER TABLE bookings
  DROP COLUMN date,
  DROP COLUMN time;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE bookings
  ADD COLUMN date DATE NOT NULL,
  ADD COLUMN time TIME NOT NULL;

UPDATE bookings
  SET date = datetime::date,
      time = datetime::time;

ALTER TABLE bookings
  DROP COLUMN datetime;
-- +goose StatementEnd
