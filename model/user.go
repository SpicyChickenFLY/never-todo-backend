package model

import "time"

// ==================== User ====================

// User is a model for todo_user
type User struct {
	ID            int       `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Name          string    `gorm:"NOT NULL"`
	CreateTime    time.Time `gorm:"default:null"`
	LastLoginTime time.Time `gorm:"default:null"`
	Status        bool      `gorm:"default:false"`
	DeleteFlag    bool      `gorm:"default:false"`
}
