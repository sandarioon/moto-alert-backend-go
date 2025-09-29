-- +goose Up
-- +goose StatementBegin
ALTER TABLE "accidents"
	ADD "geom" geometry (Point, 4326);

ALTER TABLE "users"
	ADD "geom" geometry (Point, 4326);

CREATE INDEX "accidents_geom" ON "accidents"
USING GiST ("geom");

CREATE INDEX "users_geom" ON "users"
USING GiST ("geom");

UPDATE
	"users"
SET
	"geom" = ST_SetSRID (ST_MakePoint ("longitude",
			"latitude"),
		4326);

UPDATE
	"accidents"
SET
	"geom" = ST_SetSRID (ST_MakePoint ("longitude",
			"latitude"),
		4326);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX "public"."accidents_geom";

DROP INDEX "public"."users_geom";

ALTER TABLE "users" DROP COLUMN "geom";

ALTER TABLE "accidents" DROP COLUMN "geom";
-- +goose StatementEnd
