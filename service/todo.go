package service

import (
	"errors"
	"log"

	"github.com/jinzhu/gorm"
	"spicychicken.top/NeverTODO/backend/dao"
	"spicychicken.top/NeverTODO/backend/model"
	"spicychicken.top/NeverTODO/backend/pkgs/errx"
)

// GetAllTables is a func to get all tables
func GetAllTables(tx *gorm.DB) (
	model.Tasks, model.Tags, model.TaskTags, error) {
	var tasks model.Tasks
	var tags model.Tags
	var taskTags model.TaskTags
	// retrieve all tasks
	if err := dao.GetAllTasks(tx, &tasks); err != nil {
		return tasks, tags, taskTags, errx.New(err)
	}
	// retrieve all tags
	if err := dao.GetAllTags(tx, &tags); err != nil {
		return tasks, tags, taskTags, errx.New(err)
	}
	// retrieve all tags
	if err := dao.GetAllTaskTags(tx, &taskTags); err != nil {
		return tasks, tags, taskTags, errx.New(err)
	}
	return tasks, tags, taskTags, nil
}

// FullTask is a model of task list
type FullTask struct {
	Task model.Task
	Tags model.Tags
}

// FullTasks is a slice of FullTask
type FullTasks []FullTask

// GetFullTasks is a func to get all FullTask
func GetFullTasks(tx *gorm.DB, fullTasks *FullTasks) error {
	// retrieve all tasks
	var tasks model.Tasks
	if err := dao.GetAllTasks(tx, &tasks); err != nil {
		return errx.New(err)
	}
	// retrieve tags for each task
	for i := 0; i < len(tasks); i++ {
		var tags model.Tags
		dao.GetTagsByTaskID(tx, &tags, tasks[i].ID)
		*fullTasks = append(*fullTasks,
			FullTask{
				Task: tasks[i],
				Tags: tags,
			})
	}
	return nil
}

// GetFullTasksByTag is a func to get FullTask by TagID
func GetFullTasksByTag(
	tx *gorm.DB, fullTasks *FullTasks, tagID int) error {
	log.Printf("GetFullTaskByTag(tagID: %d)\n", tagID)
	// retrieve all tasks
	var tasks model.Tasks
	if err := dao.GetTasksByTagID(tx, &tasks, tagID); err != nil {
		return errx.New(err)
	}
	// retrieve tags for each task
	for i := 0; i < len(tasks); i++ {
		var tags model.Tags
		if err := dao.GetTagsByTaskID(
			tx, &tags, tasks[i].ID); err != nil {
			return errx.New(err)
		}
		*fullTasks = append(*fullTasks,
			FullTask{
				Task: tasks[i],
				Tags: tags,
			},
		)
	}
	return nil
}

// AddFullTask is a func to add full task
// fullTask: the json object of fullTask(Task, Tags)
// return: Task(ID, Content, CreateTime, UpdateTime, Status), error
func AddFullTask(tx *gorm.DB, taskContent string, tagsID []int) (model.Task, error) {
	resultTask := model.Task{Content: taskContent}
	// check tags exists
	for _, tagID := range tagsID {
		ok, err := dao.IsExistTagID(tx, tagID)
		if err != nil {
			return resultTask, errx.New(err)
		}
		if !ok {
			err := errors.New("target TagID is not exist")
			return resultTask, errx.New(err)
		}
	}
	// add new task
	err := dao.AddTask(
		tx, &resultTask)
	if errx.New(err) != nil {
		return resultTask, err
	}
	// get new task ID
	taskID := resultTask.ID
	if err != nil {
		return resultTask, errx.New(err)
	}
	// add new task-tag
	for _, tagID := range tagsID {
		if err := dao.AddTagForTask(
			tx, int(taskID), tagID); err != nil {
			return resultTask, errx.New(err)
		}
	}
	return resultTask, nil
}

// DelFullTask is a func to add full task
// TaskID: ID of Task
func DelFullTask(tx *gorm.DB, taskID int) error {
	// errx if task exists
	ok, err := dao.IsExistTaskID(tx, taskID)
	if errx.New(err) != nil {
		return err
	}
	if !ok {
		err := errors.New("target TaskID is not exist")
		return errx.New(err)
	}
	// delete all tags for this task
	if err := dao.DelAllTagsOfTask(tx, taskID); errx.New(err) != nil {
		return err
	}
	// delete this task
	if err := dao.DelTask(tx, taskID); errx.New(err) != nil {
		return err
	}
	return nil
}

// AddTag is a func to add new tag
// tag: Tag(ID, Content, Desc)
// return: Tag(ID, Content, Desc), error
func AddTag(tx *gorm.DB, tagContent, tagDesc string) (model.Tag, error) {
	resultTag := model.Tag{Content: tagContent, Desc: tagDesc}
	// Add new tag
	err := dao.AddTag(tx, &resultTag)
	if errx.New(err) != nil {
		return resultTag, err
	}
	return resultTag, nil
}

// DelTag is a func to delete old tag
// TagID: ID of Tag
func DelTag(tx *gorm.DB, tagID int) error {
	// delete this tag for all task with it
	if err := dao.DelTagOfAllTasks(tx, tagID); errx.New(err) != nil {
		return err
	}
	// delete this tag
	if err := dao.DelTag(tx, tagID); errx.New(err) != nil {
		return err
	}
	return nil
}
