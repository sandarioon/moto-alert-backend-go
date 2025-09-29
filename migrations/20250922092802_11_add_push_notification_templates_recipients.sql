-- +goose Up
-- +goose StatementBegin
ALTER TABLE "push_notification_templates"
	ADD "recipients" character varying (150);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "push_notification_templates" DROP COLUMN "recipients";
-- +goose StatementEnd
