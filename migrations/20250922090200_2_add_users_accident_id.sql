-- +goose Up
-- +goose StatementBegin
ALTER TABLE "users"
	ADD "accident_id" integer;

CREATE INDEX "users_accident_id" ON "users" ("accident_id");

ALTER TABLE "users"
	ADD CONSTRAINT "FK_a0c4f858d470649fd2a57fec417" FOREIGN KEY ("accident_id") REFERENCES "accidents" ("id") ON DELETE NO ACTION ON
	UPDATE
		NO ACTION;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "users" DROP CONSTRAINT "FK_a0c4f858d470649fd2a57fec417";

DROP INDEX "public"."users_accident_id";

ALTER TABLE "users" DROP COLUMN "accident_id";
-- +goose StatementEnd