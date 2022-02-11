-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA "anekdot";

CREATE TABLE anekdot.sender (
	id serial NOT NULL,
	name varchar NOT NULL,
	description varchar NOT NULL DEFAULT '',
	CONSTRAINT sender_pk PRIMARY KEY (id)
);

CREATE TABLE anekdot.anekdot (
	id serial NOT NULL,
	"text" varchar NOT NULL DEFAULT '',
	rating int8 NOT NULL DEFAULT 0,
	external_id varchar NULL DEFAULT '',
	create_time timestamp(0) NOT NULL DEFAULT now(),
	status int8 NOT NULL,
	sender_id int8 NOT NULL,
	CONSTRAINT anekdot_pk PRIMARY KEY (id),
    CONSTRAINT anekdot_fk_sender FOREIGN KEY (sender_id) REFERENCES anekdot.sender(id)
);

COMMENT ON COLUMN anekdot.anekdot.status IS '1 - На проверке
2 - Разрешен
3 - Отклонен';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP SCHEMA anekdot CASCADE;
-- +goose StatementEnd
