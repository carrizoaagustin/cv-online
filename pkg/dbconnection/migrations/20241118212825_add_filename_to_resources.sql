-- +goose Up
-- +goose StatementBegin
ALTER TABLE resources
ADD COLUMN filename varchar(255);

UPDATE resources
SET filename ='default-filename';

ALTER TABLE resources
ALTER COLUMN filename SET NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE resources
DROP COLUMN filename;
-- +goose StatementEnd
