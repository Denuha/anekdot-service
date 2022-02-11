-- +goose Up
-- +goose StatementBegin
INSERT INTO anekdot.sender (id, "name", description) VALUES(1, 'anekdotme.ru', '');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM anekdot.sender where id=1;
-- +goose StatementEnd
