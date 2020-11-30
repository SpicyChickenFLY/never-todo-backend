package model

import "time"

// ==================== Task ====================

// Task is a model for todo_task
type Task struct {
	ID         int       `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Content    string    `gorm:"NOT NULL"`
	CreateTime time.Time `gorm:"NOT NULL"`
	UpdateTime time.Time `gorm:"NOT NULL"`
	Status     bool      `gorm:"NOT NULL"`
}

// Tasks is a slice of Task
type Tasks []Task

// ==================== Tag ====================

// Tag is a model for todo_tag
type Tag struct {
	ID      int    `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Content string `gorm:"NOT NULL"`
	Desc    string `gorm:"NOT NULL"`
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
