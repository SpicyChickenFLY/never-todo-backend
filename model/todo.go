package model

import (
	"time"
)

// ==================== Task ====================

// Task is a model for todo_task
type Task struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Content   string    `json:"content" gorm:"not null"`
	Completed bool      `json:"completed" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"not null;default now()"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null;default now()"`
	Deleted   bool      `json:"deleted" gorm:"not null;default 0"`
}

// Tasks is a slice of Task
type Tasks []Task

// ==================== Tag ====================

// Tag is a model for todo_tag
type Tag struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Content   string    `json:"content" gorm:"not null"`
	Desc      string    `json:"desc" gorm:"not null"`
	Color     string    `json:"color" gorm:"not null;DEAFULT '#AAAAAA'"`
	CreatedAt time.Time `json:"created_at" gorm:"not null;default now()"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null;default now()"`
	Deleted   bool      `json:"deleted" gorm:"not null;default 0"`
}

// Tags is a slice of Tag
type Tags []Tag

// ==================== TaskTag ====================

// TaskTag is a model of todo_task_tag
type TaskTag struct {
	TaskID int `json:"task_id" gorm:"primaryKey"`
	TagID  int `json:"tag_id" gorm:"primaryKey"`
}

// TaskTags is a slice of TaskTag
type TaskTags []TaskTag

// ==================== FullTask ====================

// FullTask is a model of task list
type FullTask struct {
	Task   Task  `json:"task"`
	TagsID []int `json:"tagsID"`
}

// FullTasks is a slice of FullTask
type FullTasks []FullTask
