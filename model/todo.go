package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// ==================== Task ====================

// Task is a model for todo_task
type Task struct {
	ID         int       `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Content    string    `gorm:"NOT NULL"`
	CreateTime time.Time `gorm:"default:null"`
	UpdateTime time.Time `gorm:"default:null"`
	Status     bool      `gorm:"NOT NULL"`
}

// BeforeCreate 回调函数
func (task *Task) BeforeCreate(scope *gorm.Scope) {
	if !scope.HasError() {
		now := gorm.NowFunc()
		if createdAtField, ok := scope.FieldByName("CreateTime"); ok {
			if createdAtField.IsBlank {
				createdAtField.Set(now)
			}
		}
		if updatedAtField, ok := scope.FieldByName("UpdateTime"); ok {
			if updatedAtField.IsBlank {
				updatedAtField.Set(now)
			}
		}
	}
}

// BeforeUpdate 回调函数
func (task *Task) BeforeUpdate(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("UpdateTime", gorm.NowFunc())
	}
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
