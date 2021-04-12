package dao

import (
	"fmt"
	"log"
	"time"

	"github.com/SpicyChickenFLY/never-todo-backend/model"
	"github.com/lingdor/stackerror"
	"gorm.io/gorm"
)

// ==================== Task ====================

// GetAllTasks get all Tasks
func GetAllTasks(tx *gorm.DB, tasks *model.Tasks) error {
	log.Println("GetAllTasks")
	result := tx.Where(&model.Task{Deleted: false}).Find(tasks)
	// defer result.Close()
	if err := result.Error; err != nil {
		return stackerror.New(err.Error())
	}
	return nil
}

// GetTasksByContent get Tasks by ID
func GetTasksByContent(tx *gorm.DB, tasks *model.Tasks, content string) error {
	log.Printf("GetTaskByContent(TaskID: %s)\n", content)
	content = fmt.Sprintf("%%%s%%", content)
	result := tx.Where("content LIKE '?' and deleted=0", content).First(&tasks)
	// defer result.Close()
	if err := result.Error; err != nil {
		return stackerror.New(err.Error())
	}
	return nil
}

// GetTaskByID get Tasks by ID
func GetTaskByID(tx *gorm.DB, tasks *model.Tasks, taskID int) error {
	log.Printf("GetTaskByID(TaskID: %d)\n", taskID)
	result := tx.Where(&model.Task{ID: taskID, Deleted: false}).First(&tasks)
	// defer result.Close()
	if err := result.Error; err != nil {
		return stackerror.New(err.Error())
	}
	return nil
}

// AddTask add Task
func AddTask(tx *gorm.DB, task *model.Task) error {
	log.Printf("AddTask(task: %v\n)", task)
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()
	if err := tx.Create(&task).Error; err != nil {
		return stackerror.New(err.Error())
	}
	return nil
}

// DelTask delete Task
func DelTask(tx *gorm.DB, taskID int) error {
	log.Printf("DelTask(taskID: %d) \n", taskID)
	if err := tx.Model(&model.Task{}).Where(
		"id=?", taskID).Update("deleted", true).Error; err != nil {
		return stackerror.New(err.Error())
	}
	return nil
}

// UpdTask update Task
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

// GetAllTags get all Tags
func GetAllTags(tx *gorm.DB, tags *model.Tags) error {
	log.Println("GetTagByID")
	result := tx.Find(&tags)
	// defer result.Close()
	if err := result.Error; err != nil {
		return stackerror.New(err.Error())
	}
	return nil
}

// GetTagByID get Tags by ID
func GetTagByID(tx *gorm.DB, tags *model.Tags, tagID int) error {
	log.Printf("GetTagByID(TagID: %d)\n", tagID)
	result := tx.Where(&model.Tag{ID: tagID}).First(&tags)
	// defer result.Close()
	if err := result.Error; err != nil {
		return stackerror.New(err.Error())
	}
	return nil
}

// AddTag add Tag
func AddTag(tx *gorm.DB, tag *model.Tag) error {
	log.Printf("AddTag(tag:%v\n)", tag)
	tag.CreatedAt = time.Now()
	tag.UpdatedAt = time.Now()
	if err := tx.Create(&tag).Error; err != nil {
		return stackerror.New(err.Error())
	}
	return nil
}

// DelTag delete Tag
func DelTag(tx *gorm.DB, tagID int) error {
	log.Printf("DelTag(tagID: %d) \n", tagID)
	if err := tx.Model(&model.Tag{}).Where(
		"id=?", tagID).Update("deleted", true).Error; err != nil {
		return stackerror.New(err.Error())
	}
	return nil
}

// UpdTag update Tag
func UpdTag(tx *gorm.DB, tag *model.Tag) error {
	log.Printf("UpdTag(tag: %v) \n", tag)
	tag.UpdatedAt = time.Now()
	if err := tx.Save(tag).Error; err != nil {
		return stackerror.New(err.Error())
	}
	return nil
}

// ==================== TaskTag ====================

// GetAllTaskTags get all TaskTags
func GetAllTaskTags(tx *gorm.DB, taskTags *model.TaskTags) error {
	log.Println("GetAllTaskTags")
	result := tx.Find(&taskTags)
	// defer result.Close()
	if err := result.Error; err != nil {
		return stackerror.New(err.Error())
	}
	return nil
}

// GetTaskTagsByTaskID get TaskTags by TaskID
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

// GetTaskTagsByTagID get TaskTags by TagID
func GetTaskTagsByTagID(
	tx *gorm.DB, taskTags *model.TaskTags, tagID int) error {
	log.Printf("GetTaskTagsByTagID(TagID: %d)\n", tagID)
	result := tx.Where(&model.TaskTag{TagID: tagID}).Find(&taskTags)
	if err := result.Error; err != nil {
		return stackerror.New(err.Error())
	}
	return nil
}

// AddTaskTag add TaskTag by struct
func AddTaskTag(tx *gorm.DB, taskTag *model.TaskTag) error {
	log.Printf("AddTaskTag(taskTag:%v\n)", taskTag)
	if err := tx.Create(&taskTag).Error; err != nil {
		return stackerror.New(err.Error())
	}
	return nil
}

// AddTagForTask add TaskTag by id
func AddTagForTask(tx *gorm.DB, taskID, tagID int) error {
	log.Printf("AddTagForTask(taskID: %d, tagID: %d)\n", taskID, tagID)
	taskTag := model.TaskTag{TaskID: taskID, TagID: tagID}
	if err := tx.Create(&taskTag).Error; err != nil {
		return stackerror.New(err.Error())
	}
	return nil
}

// DelTaskTag delete TaskTag by id
func DelTaskTag(tx *gorm.DB, taskTagID int) error {
	log.Printf("DelTaskTag(taskTagID: %d) \n", taskTagID)
	if err := tx.Model(&model.TaskTag{}).Where(
		"id=?", taskTagID).Update("deleted", true).Error; err != nil {
		return stackerror.New(err.Error())
	}
	return nil
}

// DelAllTagsOfTask delete all tags of a task
func DelAllTagsOfTask(tx *gorm.DB, taskID int) error {
	log.Printf("DelAllTagsOfTask(taskID: %d) \n", taskID)
	if err := tx.Where(
		"task_id=?", taskID).Delete(
		&model.TaskTag{}).Error; err != nil {
		return stackerror.New(err.Error())
	}
	return nil
}

// DelTagOfAllTasks delete a tag of all tasks
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

// GetTagsIDByTaskID get Tags By TaskID
func GetTagsIDByTaskID(tx *gorm.DB, tags *model.Tags, taskID int) error {
	log.Printf("GetTagsByTaskID(taskID: %d) \n", taskID)
	result := tx.Joins(
		"LEFT JOIN task_tags ON tags.id=task_tags.tag_id").Where(
		"task_tags.task_id = ?", taskID).Select("tags.id").Find(&tags)
	if err := result.Error; err != nil {
		return stackerror.New(err.Error())
	}
	return nil
}

// GetTasksByTagID get Tasks By TagID
func GetTasksByTagID(tx *gorm.DB, tasks *model.Tasks, tagID int) error {
	log.Printf("GetTagsByTaskID(tagID: %d) \n", tagID)
	result := tx.Joins(
		"LEFT JOIN task_tags ON tasks.id=task_tags.task_id").Where(
		"task_tags.tag_id = ?", tagID).Find(&tasks)
	if err := result.Error; err != nil {
		return stackerror.New(err.Error())
	}
	return nil
}
