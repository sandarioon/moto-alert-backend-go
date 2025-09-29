-- +goose Up
-- +goose StatementBegin
DROP INDEX "public"."users_phone";

ALTER TABLE "users" ALTER COLUMN "phone" SET DATA TYPE character varying (50);

CREATE UNIQUE INDEX "users_phone" ON "users" ("phone");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX "public"."users_phone";

ALTER TABLE "users" ALTER COLUMN "phone" SET DATA TYPE character varying (20);

CREATE UNIQUE INDEX "users_phone" ON "users" ("phone");
-- +goose StatementEnd
