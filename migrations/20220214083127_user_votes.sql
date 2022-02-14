-- +goose Up
-- +goose StatementBegin

CREATE TABLE anekdot."user" (
	id serial NOT NULL,
	username varchar NOT NULL,
	external_id varchar NOT NULL DEFAULT ''::character varying,
	realm varchar NOT NULL,
	create_time timestamp(0) NOT NULL DEFAULT now(),
	CONSTRAINT user_pk PRIMARY KEY (id)
);

CREATE TABLE anekdot.user_votes (
	id serial NOT NULL,
	user_id int8 NOT NULL,
	anekdot_id int8 NOT NULL,
	value int8 NOT NULL DEFAULT 0,
	CONSTRAINT user_votes_pk PRIMARY KEY (id),
	CONSTRAINT user_fk FOREIGN KEY (user_id) REFERENCES anekdot."user"(id),
	CONSTRAINT anekdot_fk FOREIGN KEY (anekdot_id) REFERENCES anekdot.anekdot(id)
);

ALTER TABLE anekdot.anekdot DROP COLUMN rating;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE anekdot.anekdot ADD rating int8 NOT NULL DEFAULT 0;

DROP TABLE anekdot.user_votes;

DROP TABLE anekdot."user";
-- +goose StatementEnd
