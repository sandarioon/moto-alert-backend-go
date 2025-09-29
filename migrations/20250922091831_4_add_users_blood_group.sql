-- +goose Up
-- +goose StatementBegin
CREATE TYPE "public"."users_blood_group" AS ENUM ( 'NOT_SELECTED',
	'I_0_POSITIVE',
	'I_0_NEGATIVE',
	'II_A_POSITIVE',
	'II_A_NEGATIVE',
	'III_B_POSITIVE',
	'III_B_NEGATIVE',
	'IV_AB_POSITIVE',
	'IV_AB_NEGATIVE'
);

ALTER TABLE "users"
	ADD "bloodGroup" "public"."users_blood_group" NOT NULL DEFAULT 'NOT_SELECTED';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "users" DROP COLUMN "bloodGroup";

DROP TYPE "public"."users_blood_group";
-- +goose StatementEnd
