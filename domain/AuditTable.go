package domain

import (
	"time"
)

// AuditTable : embedded for entities with created, updated, delete at and by
type AuditTable struct {
	CreatedAt time.Time `json:"created_at"`
	CreatedBy int64     `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy int64     `json:"updated_by"`
}
