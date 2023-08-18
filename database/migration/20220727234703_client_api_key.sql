-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "client_api_key"
(
    "id" uuid PRIMARY KEY,
    "name" VARCHAR(50),
    "api_key_id" VARCHAR(255),
    "api_key_secret" VARCHAR(255),
    "is_active" BOOLEAN NOT NULL,
    "client_id" uuid REFERENCES client (id),

    "created_at"      timestamp        not null,
    "updated_at"      timestamp        not null,
    "deleted_at"      timestamp,
    UNIQUE (api_key_id)

    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "client_api_key";
-- +goose StatementEnd

