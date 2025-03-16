-- +goose Up
-- +goose StatementBegin
ALTER TABLE tables DROP COLUMN status;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE tables ADD COLUMN status varchar(255) NOT NULL DEFAULT '';
-- +goose StatementEnd
