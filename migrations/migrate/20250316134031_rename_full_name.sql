-- +goose Up
-- +goose StatementBegin
ALTER TABLE bookings RENAME COLUMN full_name TO fullname;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE bookings RENAME COLUMN fullname TO full_name;
-- +goose StatementEnd
