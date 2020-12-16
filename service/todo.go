package service

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/lingdor/stackerror"
	"spicychicken.top/NeverTODO/backend/dao"
	"spicychicken.top/NeverTODO/backend/model"
)

// GetAllTables is a func to get all tables
func GetAllTables(tx *gorm.DB) (
	model.Tasks, model.Tags, model.TaskTags, error) {
	var tasks model.Tasks
	var tags model.Tags
	var taskTags model.TaskTags
	// retrieve all tasks
	if err := dao.GetAllTasks(tx, &tasks); err != nil {
		return tasks, tags, taskTags, err
	}
	// retrieve all tags
	if err := dao.GetAllTags(tx, &tags); err != nil {
		return tasks, tags, taskTags, err
	}
	// retrieve all tags
	if err := dao.GetAllTaskTags(tx, &taskTags); err != nil {
		return tasks, tags, taskTags, err
	}
	return tasks, tags, taskTags, nil
}

// GetFullTasks is a func to get all FullTask
func GetFullTasks(tx *gorm.DB, fullTasks *model.FullTasks) error {
	// retrieve all tasks
	var tasks model.Tasks
	if err := dao.GetAllTasks(tx, &tasks); err != nil {
		return err
	}
	// retrieve tags for each task
	for i := 0; i < len(tasks); i++ {
		var tags model.Tags
		dao.GetTagsByTaskID(tx, &tags, tasks[i].ID)
		*fullTasks = append(*fullTasks,
			model.FullTask{
				Task: tasks[i],
				Tags: tags,
			})
	}
	return nil
}

// GetFullTasksByTag is a func to get FullTask by TagID
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
		if err := dao.GetTagsByTaskID(
			tx, &tags, tasks[i].ID); err != nil {
			return err
		}
		*fullTasks = append(*fullTasks,
			model.FullTask{
				Task: tasks[i],
				Tags: tags,
			},
		)
	}
	return nil
}

// AddFullTask is a func to add full task
// taskContent:
// tagsID:
// return: Task(ID, Content, CreateTime, UpdateTime, Status), error
func AddFullTask(tx *gorm.DB, taskContent string, tagsID []int) (model.Task, error) {
	resultTask := model.Task{Content: taskContent}
	// check tags exists
	for _, tagID := range tagsID {
		ok, err := dao.IsExistTagID(tx, tagID)
		if err != nil {
			return resultTask, err
		}
		if !ok {
			return resultTask, stackerror.New(
				"target TagID is not exist")
		}
	}
	// add new task
	err := dao.AddTask(
		tx, &resultTask)
	if err != nil {
		return resultTask, err
	}
	// add new task-tag
	taskID := resultTask.ID
	for _, tagID := range tagsID {
		if err := dao.AddTagForTask(
			tx, int(taskID), tagID); err != nil {
			return resultTask, err
		}
	}
	return resultTask, nil
}

// UpdFullTask is a func to add full task
// task(ID, Content, CreateTime, UpdateTime, Status):
// tagsID:
// return: Task(ID, Content, CreateTime, UpdateTime, Status), error
func UpdFullTask(tx *gorm.DB, task *model.Task, tagsID []int) error {
	// delete old tags
	if err := dao.DelAllTagsOfTask(tx, task.ID); err != nil {
		return err
	}
	// insert new tags
	for _, tagID := range tagsID {
		ok, err := dao.IsExistTagID(tx, tagID)
		if err != nil {
			return err
		}
		if !ok {
			return stackerror.New(
				"target TagID is not exist")
		}
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

// DelFullTask is a func to add full task
// TaskID: ID of Task
func DelFullTask(tx *gorm.DB, taskID int) error {
	ok, err := dao.IsExistTaskID(tx, taskID)
	if err != nil {
		return err
	}
	if !ok {
		return stackerror.New(
			"target TaskID is not exist")
	}
	// delete all tags for this task
	if err := dao.DelAllTagsOfTask(tx, taskID); err != nil {
		return err
	}
	// delete this task
	return dao.DelTask(tx, taskID)
}

// AddTag is a func to add new tag
// tag: Tag(ID, Content, Desc)
// return: Tag(ID, Content, Desc), error
func AddTag(tx *gorm.DB, tagContent, tagDesc string) (model.Tag, error) {
	resultTag := model.Tag{Content: tagContent, Desc: tagDesc}
	// Add new tag
	err := dao.AddTag(tx, &resultTag)
	if err != nil {
		return resultTag, err
	}
	return resultTag, nil
}

// DelTag is a func to delete old tag
// TagID: ID of Tag
func DelTag(tx *gorm.DB, tagID int) error {
	// delete this tag for all task with it
	if err := dao.DelTagOfAllTasks(tx, tagID); err != nil {
		return err
	}
	// delete this tag
	return dao.DelTag(tx, tagID)
}

// UpdTag is a func to update old tag
func UpdTag(tx *gorm.DB, tag *model.Tag) error {
	return dao.UpdTag(tx, tag)
}
