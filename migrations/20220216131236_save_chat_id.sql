-- +goose Up
-- +goose StatementBegin
ALTER TABLE anekdot."user" ADD "chat_id" int8 NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE anekdot."user" DROP COLUMN "chat_id";
-- +goose StatementEnd
