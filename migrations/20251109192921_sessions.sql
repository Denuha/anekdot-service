-- +goose Up
-- +goose StatementBegin
CREATE TABLE anekdot."session" (
	user_id int8 NOT NULL UNIQUE,
	access_token varchar NOT NULL,
	access_token_create_time timestamptz NOT NULL DEFAULT now(),
	refresh_token varchar NOT NULL,
	refresh_token_create_time timestamptz NOT NULL DEFAULT now(),
	CONSTRAINT session_fk FOREIGN KEY (user_id) REFERENCES anekdot."user"(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE anekdot."session";
-- +goose StatementEnd
