-- +goose Up
-- +goose StatementBegin
CREATE TYPE "public"."accident_responds_state_enum" AS ENUM ( 'ACTIVE',
	'CANCELED'
);

ALTER TABLE "accident_responds"
	ADD "state" "public"."accident_responds_state_enum" NOT NULL DEFAULT 'ACTIVE';

CREATE INDEX "accident_responds_state" ON "accident_responds" ("state");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX "public"."accident_responds_state";

ALTER TABLE "accident_responds" DROP COLUMN "state";

DROP TYPE "public"."accident_responds_state_enum";
-- +goose StatementEnd
