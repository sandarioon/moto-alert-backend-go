-- +goose Up
-- +goose StatementBegin
ALTER TABLE "users" ADD "age" smallint;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "users" DROP COLUMN "age";
-- +goose StatementEnd
