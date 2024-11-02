-- +goose Up
-- INSERT INTO "user"(email, icon, password, name, admin)
-- VALUES ('admin@fincal.com', '', '', 'Admin', true);

-- +goose Down
-- +goose StatementBegin
-- SELECT 'down SQL query';
-- +goose StatementEnd
