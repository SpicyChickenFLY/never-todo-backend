package service

import (
	"log"

	"github.com/SpicyChickenFLY/NeverTODO/backend/dao"
	"github.com/SpicyChickenFLY/NeverTODO/backend/model"
	"github.com/jinzhu/gorm"
	"github.com/lingdor/stackerror"
)

// GetAllTables is a func to get all tables
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

// GetFullTasks is a func to get all FullTasks
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
func AddFullTask(tx *gorm.DB, task *model.Task, tagsID []int) error {

	// check tags exists
	for _, tagID := range tagsID {
		ok, err := dao.IsExistTagID(tx, tagID)
		if err != nil {
			return err
		}
		if !ok {
			return stackerror.New(
				"target TagID is not exist")
		}
	}
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

// GetAllTasks is a func to get all Tasks
func GetAllTasks(tx *gorm.DB, tasks *model.Tasks) error {
	// retrieve all tasks
	return dao.GetAllTasks(tx, tasks)
}

// GetAllTags is a func to get all tags
func GetAllTags(tx *gorm.DB, tags *model.Tags) error {
	// retrieve all tags
	return dao.GetAllTags(tx, tags)
}

// AddTag is a func to add new tag
// tag: Tag(ID, Content, Desc)
// return: Tag(ID, Content, Desc), error
func AddTag(tx *gorm.DB, tag model.Tag) error {
	// Add new tag
	return dao.AddTag(tx, &tag)
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
