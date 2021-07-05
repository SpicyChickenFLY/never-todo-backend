package model

import (
	"time"
)

// ==================== Sync ====================

// Sync is a model for todo_sync
type Sync struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Committed time.Time `json:"committed" gorm:"not null;default now()"`
	Committer string    `json:"committer" gorm:"not null"`
	Type      string    `json:"type" gorm:"not null"`
	Table     string    `json:"table" gorm:"not null"`
	Target    int       `json:"target" gorm:"not null"`
}
