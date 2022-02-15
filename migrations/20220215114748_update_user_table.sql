-- +goose Up
-- +goose StatementBegin

ALTER TABLE anekdot."user" ADD "password" varchar NULL;
ALTER TABLE anekdot."user" ADD is_admin bool NULL DEFAULT false;


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE anekdot."user" DROP COLUMN "password";
ALTER TABLE anekdot."user" DROP COLUMN is_admin;

-- +goose StatementEnd
