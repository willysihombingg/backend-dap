-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "admin"
(
    "id" uuid PRIMARY KEY,
    "privy_id_admin" VARCHAR(30),
    "added_by" uuid REFERENCES super_admin (id),
    "is_active" BOOLEAN NOT NULL,

    "created_at"      timestamp        not null,
    "updated_at"      timestamp        not null,
    "deleted_at"      timestamp

    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "admin";
-- +goose StatementEnd

