-- +goose Up
-- +goose StatementBegin
CREATE TABLE social_networks (
    social_network_id UUID PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE socialNetworks;
-- +goose StatementEnd
