-- +goose Up
-- +goose StatementBegin
CREATE TABLE resources(
    resource_id UUID,
    format     varchar(10) NOT NULL,
    link       varchar(255) NOT NULL,

    PRIMARY KEY(resource_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE resources
-- +goose StatementEnd
