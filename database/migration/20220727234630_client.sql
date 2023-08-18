-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "client"
(
    "id" uuid PRIMARY KEY,
    "name" VARCHAR(50),
    "description" VARCHAR(255),
    "privy_id_owner" VARCHAR(30),
    "is_active" BOOLEAN NOT NULL,

    "created_at"      timestamp        not null,
    "updated_at"      timestamp        not null,
    "deleted_at"      timestamp

    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "client";
-- +goose StatementEnd

