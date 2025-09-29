-- +goose Up
-- +goose StatementBegin
ALTER TABLE "users"
	ADD "is_banned" boolean NOT NULL DEFAULT FALSE;

ALTER TABLE "users"
	ADD "is_verified" boolean NOT NULL DEFAULT FALSE;

ALTER TABLE "users"
	ADD "is_deleted" boolean NOT NULL DEFAULT FALSE;

UPDATE
	"users"
SET
	"is_banned" = "isBanned";

UPDATE
	"users"
SET
	"is_verified" = "isVerified";

UPDATE
	"users"
SET
	"is_deleted" = "isDeleted";

ALTER TABLE "users" DROP COLUMN "isBanned";

ALTER TABLE "users" DROP COLUMN "isVerified";

ALTER TABLE "users" DROP COLUMN "isDeleted";
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "users"
	ADD "isDeleted" boolean NOT NULL DEFAULT FALSE;

ALTER TABLE "users"
	ADD "isVerified" boolean NOT NULL DEFAULT FALSE;

ALTER TABLE "users"
	ADD "isBanned" boolean NOT NULL DEFAULT FALSE;

UPDATE
	"users"
SET
	"isDeleted" = "is_deleted";

UPDATE
	"users"
SET
	"isVerified" = "is_verified";

UPDATE
	"users"
SET
	"isBanned" = "is_banned";

ALTER TABLE "users" DROP COLUMN "is_deleted";

ALTER TABLE "users" DROP COLUMN "is_verified";

ALTER TABLE "users" DROP COLUMN "is_banned";
-- +goose StatementEnd
