-- +goose Up
-- +goose StatementBegin
ALTER TABLE "users" DROP COLUMN "has_hypertension";

CREATE TYPE "public"."selection_type" AS ENUM ( 'TRUE',
	'FALSE',
	'NOT_SELECTED'
);

ALTER TABLE "users"
	ADD "has_hypertension" "public"."selection_type" NOT NULL DEFAULT 'NOT_SELECTED';

ALTER TABLE "users" DROP COLUMN "has_hepatitis";

ALTER TABLE "users"
	ADD "has_hepatitis" "public"."selection_type" NOT NULL DEFAULT 'NOT_SELECTED';

ALTER TABLE "users" DROP COLUMN "has_hiv";

ALTER TABLE "users"
	ADD "has_hiv" "public"."selection_type" NOT NULL DEFAULT 'NOT_SELECTED';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "users" DROP COLUMN "has_hiv";

ALTER TABLE "users"
	ADD "has_hiv" boolean NOT NULL DEFAULT FALSE;

ALTER TABLE "users" DROP COLUMN "has_hepatitis";

ALTER TABLE "users"
	ADD "has_hepatitis" boolean NOT NULL DEFAULT FALSE;

ALTER TABLE "users" DROP COLUMN "has_hypertension";

DROP TYPE "public"."selection_type";

ALTER TABLE "users"
	ADD "has_hypertension" boolean NOT NULL DEFAULT FALSE;
-- +goose StatementEnd
