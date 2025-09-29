-- +goose Up
-- +goose StatementBegin
ALTER TABLE "users"
	ADD "isBanned" boolean NOT NULL DEFAULT FALSE;

ALTER TABLE "users"
	ADD "isVerified" boolean NOT NULL DEFAULT FALSE;

ALTER TABLE "users"
	ADD "isDeleted" boolean NOT NULL DEFAULT FALSE;

UPDATE
	"users"
SET
	"isDeleted" = "deleted";

UPDATE
	"users"
SET
	"isVerified" = "verified";

UPDATE
	"users"
SET
	"isBanned" = "ban";

ALTER TABLE "users" DROP COLUMN "ban";

ALTER TABLE "users" DROP COLUMN "verified";

ALTER TABLE "users" DROP COLUMN "deleted";
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "users"
	ADD "deleted" boolean NOT NULL DEFAULT FALSE;

ALTER TABLE "users"
	ADD "verified" boolean NOT NULL DEFAULT FALSE;

ALTER TABLE "users"
	ADD "ban" boolean NOT NULL DEFAULT FALSE;

UPDATE
	"users"
SET
	"ban" = "isBanned";

UPDATE
	"users"
SET
	"verified" = "isVerified";

UPDATE
	"users"
SET
	"deleted" = "isDeleted";

ALTER TABLE "users" DROP COLUMN "isDeleted";

ALTER TABLE "users" DROP COLUMN "isVerified";

ALTER TABLE "users" DROP COLUMN "isBanned";
-- +goose StatementEnd
