-- +goose Up
-- +goose StatementBegin
ALTER TABLE "users"
	ADD "deleted" boolean NOT NULL DEFAULT FALSE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "users" DROP COLUMN "deleted";
-- +goose StatementEnd
