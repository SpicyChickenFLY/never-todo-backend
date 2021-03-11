package model

import (
	"time"
)

// ==================== Task ====================

// Task is a model for todo_task
type Task struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	Content   string    `gorm:"NOT NULL"`
	Status    bool      `gorm:"NOT NULL"`
	CreatedAt time.Time `gorm:"NOT NULL;default now()"`
	UpdatedAt time.Time `gorm:"NOT NULL;default now()"`
	Deleted   bool      `gorm:"NOT NULL;default 0"`
}

// Tasks is a slice of Task
type Tasks []Task

// ==================== Tag ====================

// Tag is a model for todo_tag
type Tag struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	Content   string    `gorm:"NOT NULL"`
	Desc      string    `gorm:"NOT NULL"`
	Color     string    `gorm:"NOT NULL;DEAFULT '#AAAAAA'"`
	CreatedAt time.Time `gorm:"NOT NULL;default now()"`
	UpdatedAt time.Time `gorm:"NOT NULL;default now()"`
	Deleted   bool      `gorm:"NOT NULL;default 0"`
}

// Tags is a slice of Tag
type Tags []Tag

// ==================== TaskTag ====================

// TaskTag is a model of todo_task_tag
type TaskTag struct {
	TaskID int `gorm:"primary_key"`
	TagID  int `gorm:"primary_key"`
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
