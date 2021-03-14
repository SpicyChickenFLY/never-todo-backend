package model

import (
	"time"
)

// ==================== Task ====================

// Task is a model for todo_task
type Task struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	Content   string    `gorm:"not null"`
	Status    bool      `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null;default now()"`
	UpdatedAt time.Time `gorm:"not null;default now()"`
	Deleted   bool      `gorm:"not null;default 0"`
}

// Tasks is a slice of Task
type Tasks []Task

// ==================== Tag ====================

// Tag is a model for todo_tag
type Tag struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	Content   string    `gorm:"not null"`
	Desc      string    `gorm:"not null"`
	Color     string    `gorm:"not null;DEAFULT '#AAAAAA'"`
	CreatedAt time.Time `gorm:"not null;default now()"`
	UpdatedAt time.Time `gorm:"not null;default now()"`
	Deleted   bool      `gorm:"not null;default 0"`
}

// Tags is a slice of Tag
type Tags []Tag

// ==================== TaskTag ====================

// TaskTag is a model of todo_task_tag
type TaskTag struct {
	TaskID int `gorm:"primaryKey"`
	TagID  int `gorm:"primaryKey"`
}

// TaskTags is a slice of TaskTag
type TaskTags []TaskTag

// ==================== FullTask ====================

// FullTask is a model of task list
type FullTask struct {
	Task Task
	Tags Tags
}

// FullTasks is a slice of FullTask
type FullTasks []FullTask
