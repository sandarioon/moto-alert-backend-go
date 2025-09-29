-- +goose Up
-- +goose StatementBegin
ALTER TABLE "accidents" DROP CONSTRAINT "FK_29600b229e43dd2a803ee87f188";
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "accidents"
	ADD CONSTRAINT "FK_29600b229e43dd2a803ee87f188" FOREIGN KEY ("victim_id") REFERENCES "users" ("id") ON DELETE NO ACTION ON
	UPDATE
		NO ACTION;
-- +goose StatementEnd
