package model

import (
	"time"
)

// ==================== Sync ====================

// Sync is a model for todo_sync
type Sync struct {
	ID         int       `json:"id" gorm:"primaryKey;autoIncrement"`
	SyncTime   string    `json:"content" gorm:"not null"`
	Compeleted bool      `json:"compeleted" gorm:"not null"`
	CreatedAt  time.Time `json:"created_at" gorm:"not null;default now()"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"not null;default now()"`
	Deleted    bool      `json:"deleted" gorm:"not null;default 0"`
}
