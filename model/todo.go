package model

import (
	"errors"
	"log"

	"github.com/jinzhu/gorm"
	"spicychicken.top/NeverTODO/backend/pkgs/errx"
)

// FIXME: GetResultNum is a func to get count of result
// func GetResultNum(tx *gorm.DB, sqlStmt string) (int, error) {
// 	log.Printf("GetResultNum(sqlStmt: %s)", sqlStmt)
// 	result := 0
// 	var rows *sql.Rows
// 	var err error
// 	rows, err = tx.Find(out interface{}, where ...interface{})
// 	defer rows.Close()
// 	if errx.New(err) != nil {
// 		return result, err
// 	}
// 	rows.Next()
// 	if err := rows.Scan(&result); errx.New(err) != nil {
// 		return result, err
// 	}
// 	return result, nil
// }

// ==================== Task ====================

// Task is a model for todo_task
type Task struct {
	ID         int    `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Content    string `gorm:"NOT NULL"`
	CreateTime string `gorm:"NOT NULL"`
	UpdateTime string `gorm:"NOT NULL"`
	Status     bool
}

// Tasks is a slice of Task
type Tasks []Task

// GetAllTasks is a func to get all Tasks
func GetAllTasks(tx *gorm.DB, tasks *Tasks) error {
	log.Println("GetTaskByID")
	result := tx.Find(&tasks)
	defer result.Close()
	return errx.New(result.Error)
}

// GetTaskByID is a func to get Tasks by ID
func GetTaskByID(tx *gorm.DB, tasks *Tasks, taskID int) error {
	log.Printf("GetTaskByID(TaskID: %d)\n", taskID)
	result := tx.Where(&Task{ID: taskID}).First(&tasks)
	defer result.Close()
	return errx.New(result.Error)
}

// IsExistTaskID detect if task(id) exists
func IsExistTaskID(tx *gorm.DB, taskID int) (bool, error) {
	log.Printf("IsExistTaskID(TaskID: %d)\n", taskID)
	var tasks *Tasks
	err := tx.Where(&Task{ID: taskID}).First(&tasks).Error
	if err != nil {
		return false, errx.New(err)
	}
	if len(*tasks) <= 0 {
		return false, nil
	}
	return true, nil
}

// AddTask is a func to add Task
func AddTask(tx *gorm.DB, task *Task) error {
	log.Printf("AddTask(task: %v\n)", task)
	if err := tx.Create(&task).Error; err != nil {
		return errx.New(err)
	}
	if tx.NewRecord(task) { // errx.Newss
		err := errors.New("failed to add task by content")
		return errx.New(err)
	}
	return nil
}

// DelTask is a func to delete Task
func DelTask(tx *gorm.DB, taskID int) error {
	log.Printf("DelTask(taskID: %d) \n", taskID)
	task := Task{ID: taskID}
	return errx.New(tx.Delete(&task).Error)
}

// ==================== Tag ====================

// Tag is a model for todo_tag
type Tag struct {
	ID      int `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Content string
	Desc    string
}

// Tags is a slice of Tag
type Tags []Tag

// GetAllTags is a func to get all Tags
func GetAllTags(tx *gorm.DB, tags *Tags) error {
	log.Println("GetTagByID")
	result := tx.Find(&tags)
	defer result.Close()
	return errx.New(result.Error)
}

// GetTagByID is a func to get Tags by ID
func GetTagByID(tx *gorm.DB, tags *Tags, tagID int) error {
	log.Printf("GetTagByID(TagID: %d)\n", tagID)
	result := tx.Where(&Tag{ID: tagID}).First(&tags)
	defer result.Close()
	return errx.New(result.Error)
}

// IsExistTagID detect if tag(id) exists
func IsExistTagID(tx *gorm.DB, tagID int) (bool, error) {
	log.Printf("IsExistTagID(TagID: %d)\n", tagID)
	var tags *Tags
	err := tx.Where(&Tag{ID: tagID}).First(&tags).Error
	if err != nil {
		return false, errx.New(err)
	}
	if len(*tags) <= 0 {
		return false, nil
	}
	return true, nil
}

// AddTag is a func to add Tag
func AddTag(tx *gorm.DB, tag *Tag) error {
	log.Printf("AddTag(tag:%v\n)", tag)
	if err := tx.Create(&tag).Error; err != nil {
		return errx.New(err)
	}
	if tx.NewRecord(tag) { // errx.Newss
		err := errors.New("failed to add tag by content")
		return errx.New(err)
	}
	return nil
}

// DelTag is a func to delete Tag
func DelTag(tx *gorm.DB, tagID int) error {
	log.Printf("DelTag(tagID: %d) \n", tagID)
	tag := Tag{ID: tagID}
	return errx.New(tx.Delete(&tag).Error)
}

// ==================== TaskTag ====================

// TaskTag is a model of todo_task_tag
type TaskTag struct {
	TaskID int `gorm:"primary_key"`
	TagID  int `gorm:"primary_key"`
}

// TaskTags is a slice of TaskTag
type TaskTags []TaskTag

// GetAllTaskTags is a func to get all TaskTags
func GetAllTaskTags(tx *gorm.DB, taskTags *TaskTags) error {
	log.Println("GetAllTaskTags")
	result := tx.Find(&taskTags)
	defer result.Close()
	return errx.New(result.Error)
}

// GetTaskTagsByTaskID is a func to get TaskTags by TaskID
func GetTaskTagsByTaskID(
	tx *gorm.DB, taskTags *TaskTags, taskID int) error {
	log.Printf("GetTaskTagsByTaskID(TaskID: %d)\n", taskID)
	result := tx.Where(&TaskTag{TaskID: taskID}).Find(&taskTags)
	defer result.Close()
	return errx.New(result.Error)
}

// GetTaskTagsByTagID is a func to get TaskTags by TagID
func GetTaskTagsByTagID(
	tx *gorm.DB, taskTags *TaskTags, tagID int) error {
	log.Printf("GetTaskTagsByTagID(TagID: %d)\n", tagID)
	result := tx.Where(&TaskTag{TagID: tagID}).Find(&taskTags)
	defer result.Close()
	return errx.New(result.Error)
}

// AddTagForTask is a func to add TaskTag
func AddTagForTask(tx *gorm.DB, taskID, tagID int) error {
	log.Printf("AddTagForTask(taskID: %d, tagID: %d)\n", taskID, tagID)
	taskTag := TaskTag{TaskID: taskID, TagID: tagID}
	if err := tx.Create(&taskTag).Error; err != nil {
		return errx.New(err)
	}
	if tx.NewRecord(taskTag) {
		err := errors.New("failed to add tag by content")
		return errx.New(err)
	}
	return nil
}

// DelAllTagsOfTask is a func to delete all tags of a task
func DelAllTagsOfTask(tx *gorm.DB, taskID int) error {
	log.Printf("DelAllTagsOfTask(taskID: %d) \n", taskID)
	taskTag := TaskTag{TaskID: taskID}
	return errx.New(tx.Delete(&taskTag).Error)
}

// DelTagOfAllTasks is a func to delete a tag of all tasks
func DelTagOfAllTasks(tx *gorm.DB, tagID int) error {
	log.Printf("DelTagOfAllTasks(tagID: %d) \n", tagID)
	taskTag := TaskTag{TagID: tagID}
	return errx.New(tx.Delete(&taskTag).Error)
}

// ==================== Other ====================

// GetTagsByTaskID is a func to get Tags By TaskID
func GetTagsByTaskID(tx *gorm.DB, tags *Tags, taskID int) error {
	log.Printf("GetTagsByTaskID(taskID: %d) \n", taskID)
	result := tx.Joins(
		"LEFT JOIN todo_task_tag ON todo_tag.id=todo_task_tag.tag_id").Where(
		&TaskTag{TaskID: taskID}).Find(&tags)
	defer result.Close()
	return errx.New(result.Error)
}

// GetTasksByTagID is a func to get Tasks By TagID
func GetTasksByTagID(tx *gorm.DB, tasks *Tasks, tagID int) error {
	log.Printf("GetTagsByTaskID(tagID: %d) \n", tagID)
	result := tx.Joins(
		"LEFT JOIN todo_task_tag ON todo_task.id=todo_task_tag.task_id").Where(
		&TaskTag{TagID: tagID}).Find(&tasks)
	defer result.Close()
	return errx.New(result.Error)
}
