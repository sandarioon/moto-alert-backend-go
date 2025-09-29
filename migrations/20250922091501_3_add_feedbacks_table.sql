-- +goose Up
-- +goose StatementBegin
CREATE TABLE "feedbacks" (
	"id" SERIAL NOT NULL,
	"title" text,
	"body" text NOT NULL,
	"user_id" integer NOT NULL,
	"created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(
),
	CONSTRAINT "feedbacks_id" PRIMARY KEY (
		"id"
)
);

CREATE INDEX "feedbacks_user_id" ON "feedbacks" ("user_id");

ALTER TABLE "feedbacks"
	ADD CONSTRAINT "FK_4334f6be2d7d841a9d5205a100e" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE NO ACTION ON
	UPDATE
		NO ACTION;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "feedbacks" DROP CONSTRAINT "FK_4334f6be2d7d841a9d5205a100e";

DROP INDEX "public"."feedbacks_user_id";

DROP TABLE "feedbacks";
-- +goose StatementEnd