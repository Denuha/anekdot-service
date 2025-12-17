-- +goose Up
-- +goose StatementBegin
ALTER TABLE anekdot.anekdot ADD CONSTRAINT external_uniq UNIQUE (external_id,sender_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE anekdot.anekdot DROP CONSTRAINT external_uniq;
-- +goose StatementEnd
