-- +goose Up
-- +goose StatementBegin
ALTER TABLE "users" DROP COLUMN "is_moderation_account";
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "users" ADD "is_moderation_account" boolean NOT NULL DEFAULT false;
-- +goose StatementEnd
