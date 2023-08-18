package entity

import (
	"github.com/google/uuid"
	"time"
)

type Client struct {
	ID           uuid.UUID `json:"id" db:"id,omitempty"`
	Name         string    `json:"name" db:"name,omitempty"`
	Desription   string    `json:"desription" db:"description,omitempty"`
	PrivyIdOwner string    `json:"privy_id_owner" db:"privy_id_owner,omitempty"`
	IsActive     bool      `json:"is_active" db:"is_active,omitempty"`

	CreatedAt time.Time  `json:"created_at" db:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at,omitempty"`
}

type ClientApiKey struct {
	ID           uuid.UUID `json:"id" db:"id,omitempty"`
	Name         string    `json:"name" db:"name,omitempty"`
	ApiKeyId     string    `json:"api_key_id" db:"api_key_id,omitempty"`
	ApiKeySecret string    `json:"api_key_secret" db:"api_key_secret,omitempty"`
	IsActive     bool      `json:"is_active" db:"is_active,omitempty"`
	ClientId     uuid.UUID `json:"client_id" db:"client_id,omitempty"`

	CreatedAt time.Time  `json:"created_at" db:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at,omitempty"`
}

type ClientRequestLog struct {
	ID                 uuid.UUID `json:"id" db:"id,omitempty"`
	ApiKeyId           string    `json:"api_key_id" db:"api_key_id,omitempty"`
	VendorFeatureId    uuid.UUID `json:"vendor_feature_id" db:"vendor_feature_id,omitempty"`
	ClientRequestData  string    `json:"client_request_data" db:"client_request_data,omitempty"`
	ClientResponseData string    `json:"client_response_data" db:"client_response_data,omitempty"`
	HttpStatusCode     string    `json:"http_status_code" db:"http_status_code,omitempty"`
	HttpMethod         string    `json:"http_method" db:"http_method,omitempty"`
	TransactionId      uuid.UUID `json:"transaction_id" db:"transaction_id,omitempty"`
	ClientId           uuid.UUID `json:"client_id" db:"client_id,omitempty"`

	CreatedAt time.Time  `json:"created_at" db:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at,omitempty"`
}
