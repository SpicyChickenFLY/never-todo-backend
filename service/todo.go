package service

import (
	"log"

	"github.com/SpicyChickenFLY/never-todo-backend/dao"
	"github.com/SpicyChickenFLY/never-todo-backend/model"
	"gorm.io/gorm"
)

// GetAllTables get all tables
func GetAllTables(
	tx *gorm.DB,
	tasks *model.Tasks,
	tags *model.Tags,
	taskTags *model.TaskTags) error {
	// retrieve all tasks
	if err := dao.GetAllTasks(tx, tasks); err != nil {
		return err
	}
	// retrieve all tags
	if err := dao.GetAllTags(tx, tags); err != nil {
		return err
	}
	// retrieve all tags
	return dao.GetAllTaskTags(tx, taskTags)
}

// AddFullTask service for add new full task
func AddFullTask(tx *gorm.DB, task *model.Task, tagsID []int) error {
	// add new task
	err := dao.AddTask(
		tx, task)
	if err != nil {
		return err
	}
	// add new task-tag
	taskID := task.ID
	for _, tagID := range tagsID {
		if err := dao.AddTagForTask(
			tx, int(taskID), tagID); err != nil {
			return err
		}
	}
	return nil
}

// GetAllFullTasks get all FullTasks
func GetAllFullTasks(tx *gorm.DB, fullTasks *model.FullTasks) error {
	// retrieve all tasks
	var tasks model.Tasks
	if err := dao.GetAllTasks(tx, &tasks); err != nil {
		return err
	}
	// retrieve tags for each task
	for i := 0; i < len(tasks); i++ {
		var tags model.Tags
		if err := dao.GetTagsIDByTaskID(
			tx, &tags, tasks[i].ID); err != nil {
			return err
		}
		tagsID := []int{}
		for _, tag := range tags {
			tagsID = append(tagsID, tag.ID)
		}
		*fullTasks = append(*fullTasks,
			model.FullTask{
				Task:   tasks[i],
				TagsID: tagsID,
			})
	}
	return nil
}

// GetFullTasksByContent get FullTask by TaskName
func GetFullTasksByContent(
	tx *gorm.DB, fullTasks *model.FullTasks, content string) error {
	log.Printf("GetFullTasksByContent(content: %s)\n", content)
	// retrieve all tasks
	var tasks model.Tasks
	if err := dao.GetTasksByContent(tx, &tasks, content); err != nil {
		return err
	}
	// retrieve tags for each task
	for i := 0; i < len(tasks); i++ {
		var tags model.Tags
		if err := dao.GetTagsIDByTaskID(
			tx, &tags, tasks[i].ID); err != nil {
			return err
		}
		tagsID := []int{}
		for _, tag := range tags {
			tagsID = append(tagsID, tag.ID)
		}
		*fullTasks = append(*fullTasks,
			model.FullTask{
				Task:   tasks[i],
				TagsID: tagsID,
			},
		)
	}
	return nil
}

// GetFullTasksByTag get FullTask by TagID
func GetFullTasksByTag(
	tx *gorm.DB, fullTasks *model.FullTasks, tagID int) error {
	log.Printf("GetFullTaskByTag(tagID: %d)\n", tagID)
	// retrieve all tasks
	var tasks model.Tasks
	if err := dao.GetTasksByTagID(tx, &tasks, tagID); err != nil {
		return err
	}
	// retrieve tags for each task
	for i := 0; i < len(tasks); i++ {
		var tags model.Tags
		tagsID := []int{}
		for _, tag := range tags {
			tagsID = append(tagsID, tag.ID)
		}
		*fullTasks = append(*fullTasks,
			model.FullTask{
				Task:   tasks[i],
				TagsID: tagsID,
			},
		)
	}
	return nil
}

// UpdFullTask add full task
func UpdFullTask(tx *gorm.DB, task *model.Task, tagsID []int) error {
	// delete old tags
	if err := dao.DelAllTagsOfTask(tx, task.ID); err != nil {
		return err
	}
	// insert new tags
	for _, tagID := range tagsID {
		if err := dao.AddTagForTask(tx, task.ID, tagID); err != nil {
			return err
		}
	}
	// update old task
	if err := dao.UpdTask(tx, task); err != nil {
		return err
	}
	var tasks model.Tasks
	if err := dao.GetTaskByID(tx, &tasks, task.ID); err != nil {
		return err
	}
	task = &tasks[0]
	return nil
}

// DelFullTask add full task
func DelFullTask(tx *gorm.DB, taskID int) error {
	// delete all tags for this task
	if err := dao.DelAllTagsOfTask(tx, taskID); err != nil {
		return err
	}
	// delete this task
	return dao.DelTask(tx, taskID)
}
