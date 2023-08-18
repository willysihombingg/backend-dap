-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "client_request_log"
(
    "id" uuid PRIMARY KEY,
    "api_key_id" VARCHAR(255) REFERENCES client_api_key (apcleari_key_id),
    "vendor_feature_id" uuid,
    "client_request_data" json,
    "client_response_data" json,
    "http_status_code" integer,
    "http_method" VARCHAR(100),
    "transaction_id" uuid,
    "client_id" uuid,

    "created_at"      timestamp        not null,
    "updated_at"      timestamp        not null,
    "deleted_at"      timestamp

    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "client_request_log";
-- +goose StatementEnd

