package entity

import (
	"github.com/google/uuid"
	"time"
)

type SuperAdmin struct {
	ID           uuid.UUID `json:"id" db:"id,omitempty"`
	PrivyIdAdmin string    `json:"privy_id_admin" db:"privy_id_admin,omitempty"`
	IsActive     bool      `json:"is_active" db:"is_active,omitempty"`

	CreatedAt time.Time  `json:"created_at" db:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at,omitempty"`
}

type Admin struct {
	ID           uuid.UUID `json:"id" db:"id,omitempty"`
	PrivyIdAdmin string    `json:"privy_id_admin" db:"privy_id_admin,omitempty"`
	AddedBy      uuid.UUID `json:"added_by" db:"added_by,omitempty"`
	IsActive     bool      `json:"is_active" db:"is_active,omitempty"`

	CreatedAt time.Time  `json:"created_at" db:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at,omitempty"`
}
