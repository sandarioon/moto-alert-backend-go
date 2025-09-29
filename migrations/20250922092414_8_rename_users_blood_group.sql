-- +goose Up
-- +goose StatementBegin
ALTER TABLE "users" RENAME COLUMN "bloodGroup" TO "blood_group";
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "users" RENAME COLUMN "blood_group" TO "bloodGroup";
-- +goose StatementEnd
