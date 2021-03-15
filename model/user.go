package model

import "time"

// ==================== User ====================

// UserPublic

// User is a model for todo_user
type User struct {
	// Privite
	ID       int    `gorm:"primaryKey;autoIncrement"`
	Email    string `gorm:"not null"`
	Password string `gorm:"not null"`
	Deleted  bool   `gorm:"not null"`
	// Public
	Nick        string    `gorm:"not null"`
	CreateAt    time.Time `gorm:"default:null"`
	LastLoginAt time.Time `gorm:"default:null"`
	Status      int       `gorm:"not null"`
}

// Users is a slice of User
type Users []User
