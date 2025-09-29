-- +goose Up
-- +goose StatementBegin
CREATE TYPE "users_gender_enum" AS ENUM ( 'MALE',
	'FEMALE'
);

CREATE TABLE "users" (
	"id" SERIAL NOT NULL,
	"ban" boolean NOT NULL DEFAULT FALSE,
	"verified" boolean NOT NULL DEFAULT FALSE,
	"code" character varying (6) NOT NULL,
	"email" character varying (100) NOT NULL,
	"hashed_password" text NOT NULL,
	"first_name" character varying (50) NOT NULL,
	"last_name" character varying (50) NOT NULL,
	"username" character varying (50),
	"expo_push_token" character varying (50),
	"gender" "users_gender_enum" NOT NULL DEFAULT 'MALE',
	"phone" character varying (20) NOT NULL,
	"longitude" double precision,
	"latitude" double precision,
	"bike_model" character varying (50) NOT NULL,
	"comment" text,
	"last_auth" TIMESTAMP WITH TIME ZONE,
	"geo_updated_at" TIMESTAMP WITH TIME ZONE,
	"created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(
),
	CONSTRAINT "users_id" PRIMARY KEY (
		"id"
)
);

CREATE UNIQUE INDEX "users_username" ON "users" ("username");

CREATE UNIQUE INDEX "users_phone" ON "users" ("phone");

CREATE UNIQUE INDEX "users_email" ON "users" ("email");

CREATE TABLE "push_notification_templates" (
	"id" SERIAL NOT NULL,
	"type" character varying NOT NULL,
	"title" text NOT NULL,
	"body" text NOT NULL,
	CONSTRAINT "push_notification_templates_id" PRIMARY KEY ("id")
);

CREATE INDEX "push_notification_templates_type" ON "push_notification_templates" ("type");

CREATE TABLE "email_templates" (
	"id" SERIAL NOT NULL,
	"type" character varying NOT NULL,
	"title" text NOT NULL,
	"body" text NOT NULL,
	CONSTRAINT "email_templates_id" PRIMARY KEY ("id")
);

CREATE INDEX "email_templates_type" ON "email_templates" ("type");

CREATE TABLE "emergency_contacts" (
	"id" SERIAL NOT NULL,
	"first_name" character varying (50) NOT NULL,
	"last_name" character varying (50) NOT NULL,
	"relation" character varying (100) NOT NULL,
	"user_id" integer NOT NULL,
	"phone" character varying (20) NOT NULL,
	"created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(
),
	CONSTRAINT "emergency_contacts_id" PRIMARY KEY (
		"id"
)
);

CREATE INDEX "emergency_contacts_user_id" ON "emergency_contacts" ("user_id");

CREATE TYPE "accidents_state_enum" AS ENUM ( 'ACTIVE',
	'EXPIRED',
	'CANCELED',
	'COMPLETED'
);

CREATE TABLE "accidents" (
	"id" SERIAL NOT NULL,
	"state" "accidents_state_enum" NOT NULL DEFAULT 'ACTIVE',
	"victim_id" integer NOT NULL,
	"title" text,
	"description" text,
	"image_name" text,
	"longitude" double precision NOT NULL,
	"latitude" double precision NOT NULL,
	"created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(
),
	CONSTRAINT "accidents_id" PRIMARY KEY (
		"id"
)
);

CREATE INDEX "accidents_state" ON "accidents" ("state");

CREATE INDEX "accidents_victim_id" ON "accidents" ("victim_id");

CREATE INDEX "accidents_victim_id_state" ON "accidents" ("victim_id", "state");

CREATE TABLE "chats" (
	"id" SERIAL NOT NULL,
	"accident_id" integer NOT NULL,
	CONSTRAINT "chats_id" PRIMARY KEY ("id")
);

CREATE INDEX "chats_accident_id" ON "chats" ("accident_id");

CREATE TABLE "chat_messages" (
	"id" SERIAL NOT NULL,
	"chat_id" integer NOT NULL,
	"user_id" integer NOT NULL,
	"message" text NOT NULL,
	"timestamp" bigint NOT NULL,
	CONSTRAINT "chat_messages_id" PRIMARY KEY ("id")
);

CREATE INDEX "chat_messages_chat_id" ON "chat_messages" ("chat_id");

CREATE INDEX "chat_messages_user_id" ON "chat_messages" ("user_id");

CREATE INDEX "chat_messages_timestamp" ON "chat_messages" ("timestamp");

CREATE TABLE "accident_responds" (
	"id" SERIAL NOT NULL,
	"accident_id" integer NOT NULL,
	"user_id" integer NOT NULL,
	"created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(
),
	CONSTRAINT "accident_responds_id" PRIMARY KEY (
		"id"
)
);

CREATE INDEX "accident_responds_accident_id" ON "accident_responds" ("accident_id");

CREATE TABLE "accidents_push_recipients" (
	"accident_id" integer NOT NULL,
	"user_id" integer NOT NULL,
	CONSTRAINT "PK_37f3b4ace5429c7f44f2a7b8b45" PRIMARY KEY ("accident_id", "user_id")
);

CREATE INDEX "IDX_1474c9e2179850ccbf4eeb5172" ON "accidents_push_recipients" ("accident_id");

CREATE INDEX "IDX_5ee99733b0df79de442bd2c069" ON "accidents_push_recipients" ("user_id");

CREATE TABLE "chats_users" (
	"chat_id" integer NOT NULL,
	"user_id" integer NOT NULL,
	CONSTRAINT "PK_c17eee035ab608024c453194450" PRIMARY KEY ("chat_id", "user_id")
);

CREATE INDEX "IDX_ad093e1e4f96074d43e91e8501" ON "chats_users" ("chat_id");

CREATE INDEX "IDX_ecc5e7195e36df05a27debc649" ON "chats_users" ("user_id");

ALTER TABLE "emergency_contacts"
	ADD CONSTRAINT "FK_1cf39ea46db44d95b34d58d3605" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE NO ACTION ON
	UPDATE
		NO ACTION;

ALTER TABLE "accidents"
	ADD CONSTRAINT "FK_29600b229e43dd2a803ee87f188" FOREIGN KEY ("victim_id") REFERENCES "users" ("id") ON DELETE NO ACTION ON
	UPDATE
		NO ACTION;

ALTER TABLE "chats"
	ADD CONSTRAINT "FK_1ccb1c641181320e64e2648fad8" FOREIGN KEY ("accident_id") REFERENCES "accidents" ("id") ON DELETE NO ACTION ON
	UPDATE
		NO ACTION;

ALTER TABLE "chat_messages"
	ADD CONSTRAINT "FK_9f5c0b96255734666b7b4bc98c3" FOREIGN KEY ("chat_id") REFERENCES "chats" ("id") ON DELETE NO ACTION ON
	UPDATE
		NO ACTION;

ALTER TABLE "chat_messages"
	ADD CONSTRAINT "FK_5588b6cea298cedec7063c0d33e" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE NO ACTION ON
	UPDATE
		NO ACTION;

ALTER TABLE "accident_responds"
	ADD CONSTRAINT "FK_9c92e58d22ade05ea03eecd1be5" FOREIGN KEY ("accident_id") REFERENCES "accidents" ("id") ON DELETE NO ACTION ON
	UPDATE
		NO ACTION;

ALTER TABLE "accident_responds"
	ADD CONSTRAINT "FK_e6f83bb9eb4c39b52ca6d58fc7d" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE NO ACTION ON
	UPDATE
		NO ACTION;

ALTER TABLE "accidents_push_recipients"
	ADD CONSTRAINT "FK_1474c9e2179850ccbf4eeb51723" FOREIGN KEY ("accident_id") REFERENCES "accidents" ("id") ON DELETE CASCADE ON
	UPDATE
		CASCADE;

ALTER TABLE "accidents_push_recipients"
	ADD CONSTRAINT "FK_5ee99733b0df79de442bd2c0691" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE ON
	UPDATE
		CASCADE;

ALTER TABLE "chats_users"
	ADD CONSTRAINT "FK_ad093e1e4f96074d43e91e85016" FOREIGN KEY ("chat_id") REFERENCES "chats" ("id") ON DELETE CASCADE ON
	UPDATE
		CASCADE;

ALTER TABLE "chats_users"
	ADD CONSTRAINT "FK_ecc5e7195e36df05a27debc649d" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE ON
	UPDATE
		CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "chats_users" DROP CONSTRAINT "FK_ecc5e7195e36df05a27debc649d";

ALTER TABLE "chats_users" DROP CONSTRAINT "FK_ad093e1e4f96074d43e91e85016";

ALTER TABLE "accidents_push_recipients" DROP CONSTRAINT "FK_5ee99733b0df79de442bd2c0691";

ALTER TABLE "accidents_push_recipients" DROP CONSTRAINT "FK_1474c9e2179850ccbf4eeb51723";

ALTER TABLE "accident_responds" DROP CONSTRAINT "FK_e6f83bb9eb4c39b52ca6d58fc7d";

ALTER TABLE "accident_responds" DROP CONSTRAINT "FK_9c92e58d22ade05ea03eecd1be5";

ALTER TABLE "chat_messages" DROP CONSTRAINT "FK_5588b6cea298cedec7063c0d33e";

ALTER TABLE "chat_messages" DROP CONSTRAINT "FK_9f5c0b96255734666b7b4bc98c3";

ALTER TABLE "chats" DROP CONSTRAINT "FK_1ccb1c641181320e64e2648fad8";

ALTER TABLE "accidents" DROP CONSTRAINT "FK_29600b229e43dd2a803ee87f188";

ALTER TABLE "emergency_contacts" DROP CONSTRAINT "FK_1cf39ea46db44d95b34d58d3605";

DROP INDEX "IDX_ecc5e7195e36df05a27debc649";

DROP INDEX "IDX_ad093e1e4f96074d43e91e8501";

DROP TABLE "chats_users";

DROP INDEX "IDX_5ee99733b0df79de442bd2c069";

DROP INDEX "IDX_1474c9e2179850ccbf4eeb5172";

DROP TABLE "accidents_push_recipients";

DROP INDEX "accident_responds_accident_id";

DROP TABLE "accident_responds";

DROP INDEX "chat_messages_timestamp";

DROP INDEX "chat_messages_user_id";

DROP INDEX "chat_messages_chat_id";

DROP TABLE "chat_messages";

DROP INDEX "chats_accident_id";

DROP TABLE "chats";

DROP INDEX "accidents_victim_id_state";

DROP INDEX "accidents_victim_id";

DROP INDEX "accidents_state";

DROP TABLE "accidents";

DROP TYPE "accidents_state_enum";

DROP INDEX "emergency_contacts_user_id";

DROP TABLE "emergency_contacts";

DROP INDEX "email_templates_type";

DROP TABLE "email_templates";

DROP INDEX "push_notification_templates_type";

DROP TABLE "push_notification_templates";

DROP INDEX "users_email";

DROP INDEX "users_phone";

DROP INDEX "users_username";

DROP TABLE "users";

DROP TYPE "users_gender_enum";
-- +goose StatementEnd