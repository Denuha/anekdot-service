-- +goose Up
-- +goose StatementBegin
INSERT INTO anekdot.user (id, "username", external_id, realm) 
VALUES(1, 'anekdotme.ru', '', 'anekdot'),
(2, 'guest', '', 'anekdot');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
