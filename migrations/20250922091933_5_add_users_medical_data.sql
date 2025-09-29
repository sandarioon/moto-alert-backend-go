-- +goose Up
-- +goose StatementBegin
ALTER TABLE "users"
	ADD "height_cm" smallint;

ALTER TABLE "users"
	ADD "weight_kg" smallint;

ALTER TABLE "users"
	ADD "date_of_birth" date;

ALTER TABLE "users"
	ADD "chronic_diseases" text;

ALTER TABLE "users"
	ADD "allergies" text;

ALTER TABLE "users"
	ADD "medications" text;

ALTER TABLE "users"
	ADD "has_hypertension" boolean NOT NULL DEFAULT FALSE;

ALTER TABLE "users"
	ADD "has_hepatitis" boolean NOT NULL DEFAULT FALSE;

ALTER TABLE "users"
	ADD "has_hiv" boolean NOT NULL DEFAULT FALSE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "users" DROP COLUMN "has_hiv";

ALTER TABLE "users" DROP COLUMN "has_hepatitis";

ALTER TABLE "users" DROP COLUMN "has_hypertension";

ALTER TABLE "users" DROP COLUMN "medications";

ALTER TABLE "users" DROP COLUMN "allergies";

ALTER TABLE "users" DROP COLUMN "chronic_diseases";

ALTER TABLE "users" DROP COLUMN "date_of_birth";

ALTER TABLE "users" DROP COLUMN "weight_kg";

ALTER TABLE "users" DROP COLUMN "height_cm";
-- +goose StatementEnd
