package model

import "time"

// ==================== User ====================

// User is a model for todo_user
type User struct {
	ID          int       `gorm:"primaryKey;autoIncrement"`
	Email       string    `gorm:"NOT NULL"`
	Password    string    `gorm:"NOT NULL"`
	Nick        string    `gorm:"NOT NULL"`
	CreateAt    time.Time `gorm:"default:null"`
	LastLoginAt time.Time `gorm:"default:null"`
	Deleted     bool      `gorm:"NOT NULL"`
	Status      bool      `gorm:"NOT NULL"`
}

// Users is a slice of User
type Users []User
