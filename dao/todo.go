package dao

import (
	"errors"
	"log"
	"time"

	"github.com/SpicyChickenFLY/never-todo-backend/model"
	"github.com/lingdor/stackerror"
	"gorm.io/gorm"
)

// ==================== Task ====================

// GetAllTasks is a func to get all Tasks
func GetAllTasks(tx *gorm.DB, tasks *model.Tasks) error {
	log.Println("GetAllTasks")
	result := tx.Where(&model.Task{Deleted: false}).Find(tasks)
	// defer result.Close()
	if err := result.Error; err != nil {
		return stackerror.New(err.Error())
	}
	return nil
}

// GetTaskByID is a func to get Tasks by ID
func GetTaskByID(tx *gorm.DB, tasks *model.Tasks, taskID int) error {
	log.Printf("GetTaskByID(TaskID: %d)\n", taskID)
	result := tx.Where(&model.Task{ID: taskID, Deleted: false}).First(&tasks)
	// defer result.Close()
	if err := result.Error; err != nil {
		return stackerror.New(err.Error())
	}
	return nil
}

// IsExistTaskID detect if task(id) exists
func IsExistTaskID(tx *gorm.DB, taskID int) (bool, error) {
	log.Printf("IsExistTaskID(TaskID: %d)\n", taskID)
	var tasks *model.Tasks = &model.Tasks{}
	err := tx.Where(&model.Task{ID: taskID}).First(&tasks).Error
	if err != nil {
		return false, stackerror.New(err.Error())
	}
	if len(*tasks) <= 0 {
		return false, nil
	}
	return true, nil
}

// AddTask is a func to add Task
func AddTask(tx *gorm.DB, task *model.Task) error {
	log.Printf("AddTask(task: %v\n)", task)
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()
	if err := tx.Create(&task).Error; err != nil {
		return stackerror.New(err.Error())
	}
	return nil
}

// DelTask is a func to delete Task
func DelTask(tx *gorm.DB, taskID int) error {
	log.Printf("DelTask(taskID: %d) \n", taskID)
	if err := tx.Model(&model.Task{}).Where(
		"id=?", taskID).Update("deleted", true).Error; err != nil {
		return stackerror.New(err.Error())
	}
	return nil
}

// UpdTask is a func to update Task
func UpdTask(tx *gorm.DB, task *model.Task) error {
	log.Printf("UpdTask(task: %v) \n", task)
	task.UpdatedAt = time.Now()
	if err := tx.Model(task).Where(
		"id=?", task.ID).Update(
		"content", task.Content).Error; err != nil {
		return stackerror.New(err.Error())
	}
	return nil
}

// ==================== Tag ====================

// GetAllTags is a func to get all Tags
func GetAllTags(tx *gorm.DB, tags *model.Tags) error {
	log.Println("GetTagByID")
	result := tx.Find(&tags)
	// defer result.Close()
	if err := result.Error; err != nil {
		return stackerror.New(err.Error())
	}
	return nil
}

// GetTagByID is a func to get Tags by ID
func GetTagByID(tx *gorm.DB, tags *model.Tags, tagID int) error {
	log.Printf("GetTagByID(TagID: %d)\n", tagID)
	result := tx.Where(&model.Tag{ID: tagID}).First(&tags)
	// defer result.Close()
	if err := result.Error; err != nil {
		return stackerror.New(err.Error())
	}
	return nil
}

// IsExistTagID detect if tag(id) exists
func IsExistTagID(tx *gorm.DB, tagID int) (bool, error) {
	log.Printf("IsExistTagID(TagID: %d)\n", tagID)
	var tags *model.Tags = &model.Tags{}
	if err := tx.Where(
		&model.Tag{ID: tagID}).First(
		&tags).Error; err != nil {
		return false, stackerror.New(err.Error())
	}
	if len(*tags) <= 0 {
		return false, nil
	}
	return true, nil
}

// AddTag is a func to add Tag
func AddTag(tx *gorm.DB, tag *model.Tag) error {
	log.Printf("AddTag(tag:%v\n)", tag)
	tag.CreatedAt = time.Now()
	tag.UpdatedAt = time.Now()
	if err := tx.Create(&tag).Error; err != nil {
		return stackerror.New(err.Error())
	}
	return nil
}

// DelTag is a func to delete Tag
func DelTag(tx *gorm.DB, tagID int) error {
	log.Printf("DelTag(tagID: %d) \n", tagID)
	if err := tx.Model(&model.Tag{}).Where(
		"id=?", tagID).Update("deleted", true).Error; err != nil {
		return stackerror.New(err.Error())
	}
	return nil
}

// UpdTag is a func to update Tag
func UpdTag(tx *gorm.DB, tag *model.Tag) error {
	log.Printf("UpdTag(tag: %v) \n", tag)
	tag.UpdatedAt = time.Now()
	if err := tx.Save(tag).Error; err != nil {
		return stackerror.New(err.Error())
	}
	return nil
}

// ==================== TaskTag ====================

// GetAllTaskTags is a func to get all TaskTags
func GetAllTaskTags(tx *gorm.DB, taskTags *model.TaskTags) error {
	log.Println("GetAllTaskTags")
	result := tx.Find(&taskTags)
	// defer result.Close()
	if err := result.Error; err != nil {
		return stackerror.New(err.Error())
	}
	return nil
}

// GetTaskTagsByTaskID is a func to get TaskTags by TaskID
func GetTaskTagsByTaskID(
	tx *gorm.DB, taskTags *model.TaskTags, taskID int) error {
	log.Printf("GetTaskTagsByTaskID(TaskID: %d)\n", taskID)
	result := tx.Where(&model.TaskTag{TaskID: taskID}).Find(&taskTags)
	// defer result.Close()
	if err := result.Error; err != nil {
		return stackerror.New(err.Error())
	}
	return nil
}

// GetTaskTagsByTagID is a func to get TaskTags by TagID
func GetTaskTagsByTagID(
	tx *gorm.DB, taskTags *model.TaskTags, tagID int) error {
	log.Printf("GetTaskTagsByTagID(TagID: %d)\n", tagID)
	result := tx.Where(&model.TaskTag{TagID: tagID}).Find(&taskTags)
	// defer result.Close()
	if err := result.Error; err != nil {
		return stackerror.New(err.Error())
	}
	return nil
}

// AddTagForTask is a func to add TaskTag
func AddTagForTask(tx *gorm.DB, taskID, tagID int) error {
	log.Printf("AddTagForTask(taskID: %d, tagID: %d)\n", taskID, tagID)
	taskTag := model.TaskTag{TaskID: taskID, TagID: tagID}
	if err := tx.Create(&taskTag).Error; err != nil {
		return stackerror.New(err.Error())
	}
	if err := tx.Create(taskTag).Error; err != nil {
		err := errors.New("failed to add tag by content")
		return stackerror.New(err.Error())
	}
	return nil
}

// DelAllTagsOfTask is a func to delete all tags of a task
func DelAllTagsOfTask(tx *gorm.DB, taskID int) error {
	log.Printf("DelAllTagsOfTask(taskID: %d) \n", taskID)
	if err := tx.Where(
		"task_id=?", taskID).Delete(
		&model.TaskTag{}).Error; err != nil {
		return stackerror.New(err.Error())
	}
	return nil
}

// DelTagOfAllTasks is a func to delete a tag of all tasks
func DelTagOfAllTasks(tx *gorm.DB, tagID int) error {
	log.Printf("DelTagOfAllTasks(tagID: %d) \n", tagID)
	if err := tx.Where(
		"tag_id=?", tagID).Delete(
		&model.TaskTag{}).Error; err != nil {
		return stackerror.New(err.Error())
	}
	return nil
}

// ==================== Other ====================

// GetTagsByTaskID is a func to get Tags By TaskID
func GetTagsByTaskID(tx *gorm.DB, tags *model.Tags, taskID int) error {
	log.Printf("GetTagsByTaskID(taskID: %d) \n", taskID)
	result := tx.Joins(
		"LEFT JOIN task_tags ON tags.id=task_tags.tag_id AND task_tags.task_id = ?", taskID).Find(&tags)
	// defer result.Close()
	if err := result.Error; err != nil {
		return stackerror.New(err.Error())
	}
	return nil
}

// GetTasksByTagID is a func to get Tasks By TagID
func GetTasksByTagID(tx *gorm.DB, tasks *model.Tasks, tagID int) error {
	log.Printf("GetTagsByTaskID(tagID: %d) \n", tagID)
	result := tx.Joins(
		"LEFT JOIN task_tags ON tasks.id=task_tags.task_id AND task_tags.tag_id = ?", tagID).Find(&tasks)
	// defer result.Close()
	if err := result.Error; err != nil {
		return stackerror.New(err.Error())
	}
	return nil
}
