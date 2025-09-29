-- +goose Up
-- +goose StatementBegin
ALTER TABLE "users"
	ADD "uuid" character varying (36);

UPDATE
	"users"
SET
	"uuid" = "id";

CREATE UNIQUE INDEX "users_uuid" ON "users" ("uuid");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX "public"."users_uuid";

ALTER TABLE "users" DROP COLUMN "uuid";
-- +goose StatementEnd
