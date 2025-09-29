-- +goose Up
-- +goose StatementBegin
ALTER TABLE "users"
	ADD "is_moderation_account" boolean NOT NULL DEFAULT FALSE;

ALTER TABLE "users" ALTER COLUMN "first_name" DROP NOT NULL;

ALTER TABLE "users" ALTER COLUMN "last_name" DROP NOT NULL;

ALTER TABLE "users" ALTER COLUMN "phone" DROP NOT NULL;

ALTER TABLE "users" ALTER COLUMN "bike_model" DROP NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "users" ALTER COLUMN "bike_model" SET NOT NULL;

ALTER TABLE "users" ALTER COLUMN "phone" SET NOT NULL;

ALTER TABLE "users" ALTER COLUMN "last_name" SET NOT NULL;

ALTER TABLE "users" ALTER COLUMN "first_name" SET NOT NULL;

ALTER TABLE "users" DROP COLUMN "is_moderation_account";
-- +goose StatementEnd
