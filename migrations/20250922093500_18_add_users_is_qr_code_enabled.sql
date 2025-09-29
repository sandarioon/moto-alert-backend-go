-- +goose Up
-- +goose StatementBegin
ALTER TABLE "users"
	ADD "is_qr_code_enabled" boolean NOT NULL DEFAULT FALSE;

ALTER TABLE "users" ALTER COLUMN "uuid" SET NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "users" ALTER COLUMN "uuid" DROP NOT NULL;

ALTER TABLE "users" DROP COLUMN "is_qr_code_enabled";
-- +goose StatementEnd
