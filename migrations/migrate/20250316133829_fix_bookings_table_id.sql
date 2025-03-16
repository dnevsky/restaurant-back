-- +goose Up
-- +goose StatementBegin
ALTER TABLE bookings RENAME COLUMN id_table TO table_id;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE bookings RENAME COLUMN table_id TO id_table;
-- +goose StatementEnd
