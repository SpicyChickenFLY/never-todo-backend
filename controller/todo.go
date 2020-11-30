package controller

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"spicychicken.top/NeverTODO/backend/pkgs/errx"
	"spicychicken.top/NeverTODO/backend/pkgs/mysql"
	"spicychicken.top/NeverTODO/backend/service"
)

// GetAll search for all tables and return
// url:/todo/getall
func GetAll(c *gin.Context) {
	tx := mysql.GormDB.Begin()
	tasks, tags, taskTags, err := service.GetAllTables(tx)
	result := gin.H{
		"tasks":         tasks,
		"tags":          tags,
		"task_tag_pair": taskTags,
	}
	if errx.New(err) != nil {
		if err := tx.Rollback().Error; errx.New(err) != nil {
			c.String(http.StatusInternalServerError,
				fmt.Sprintf("TX Rollback Error %s", err))
			return
		}
		c.JSON(http.StatusInternalServerError, result)
		return
	}
	if err := tx.Commit().Error; errx.New(err) != nil {
		c.String(http.StatusInternalServerError,
			fmt.Sprintf("TX Commit Error: %s", err))
		return
	}
	c.JSON(http.StatusOK, result)
}

// GetTaskList is a func to
func GetTaskList(c *gin.Context) {
	tx := mysql.GormDB.Begin()
	var fullTasks service.FullTasks
	if err := service.GetFullTasks(
		tx, &fullTasks); errx.New(err) != nil {
		if err := tx.Rollback().Error; errx.New(err) != nil {
			c.String(http.StatusInternalServerError,
				fmt.Sprintf("TX Rollback Error %s", err))
			return
		}
		c.JSON(http.StatusInternalServerError,
			gin.H{"fullTasks": fullTasks})
	} else {
		if err := tx.Commit().Error; errx.New(err) != nil {
			c.String(http.StatusInternalServerError,
				fmt.Sprintf("TX Commmit Error: %s", err))
			return
		}
		c.JSON(http.StatusOK,
			gin.H{"fullTasks": fullTasks})
	}
}

// GetTaskListByTag is a func to
func GetTaskListByTag(c *gin.Context) {
	tagIDStr := c.Query("tag_id")
	if tagIDStr == "" {
		GetTaskList(c)
		return
	}
	tagID, err := strconv.Atoi(tagIDStr)
	if errx.New(err) != nil {
		return
	}
	tx := mysql.GormDB.Begin()
	var fullTasks service.FullTasks
	if err := service.GetFullTasksByTag(tx, &fullTasks, tagID); errx.New(err) != nil {
		if err := tx.Rollback().Error; errx.New(err) != nil {
			c.String(http.StatusInternalServerError,
				fmt.Sprintf("TX Rollback Error %s", err))
			return
		}
		c.JSON(http.StatusInternalServerError,
			gin.H{"fullTasks": fullTasks})
	} else {
		if err := tx.Commit().Error; errx.New(err) != nil {
			c.String(http.StatusInternalServerError,
				fmt.Sprintf("TX Commmit Error: %s", err))
			return
		}
		c.JSON(http.StatusOK,
			gin.H{"fullTasks": fullTasks})
	}
}

// AddNewTask is a func to add Task
func AddNewTask(c *gin.Context) {
	data := &struct {
		TaskContent string
		TagsID      []int
	}{}
	c.BindJSON(&data)
	log.Printf("receive post request: %v", data)
	tx := mysql.GormDB.Begin()
	if task, err := service.AddFullTask(
		tx, data.TaskContent, data.TagsID); errx.New(err) != nil {
		if err := tx.Rollback().Error; errx.New(err) != nil {
			c.String(http.StatusInternalServerError,
				fmt.Sprintf("TX Rollback Error %s", err))
			return
		}
		c.JSON(http.StatusInternalServerError,
			gin.H{"status": -1, "task": task})
	} else {
		if err := tx.Commit().Error; errx.New(err) != nil {
			c.String(http.StatusInternalServerError,
				fmt.Sprintf("TX Commmit Error: %s", err))
			return
		}
		c.JSON(http.StatusOK,
			gin.H{"status": 0, "task": task})
	}
}

// DelOldTask is a func to delete Task
func DelOldTask(c *gin.Context) {
	data := &struct {
		TaskID int
	}{}
	c.BindJSON(&data)
	if data.TaskID == 0 {
		err := errors.New("no TaskID in data field in post request")
		errx.New(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"status": -1})
		return
	}
	log.Printf("receive post request: %d", data.TaskID)
	tx := mysql.GormDB.Begin()
	if err := service.DelFullTask(tx, data.TaskID); errx.New(err) != nil {
		if err := tx.Rollback().Error; errx.New(err) != nil {
			c.String(http.StatusInternalServerError,
				fmt.Sprintf("TX Rollback Error %s", err))
			return
		}
		c.JSON(http.StatusInternalServerError,
			gin.H{"status": -1})
	} else {
		if err := tx.Commit().Error; errx.New(err) != nil {
			c.String(http.StatusInternalServerError,
				fmt.Sprintf("TX Commmit Error: %s", err))
			return
		}
		c.JSON(http.StatusOK,
			gin.H{"status": 0})
	}
}

// AddNewTag is a func to add Tag
func AddNewTag(c *gin.Context) {
	data := &struct {
		TagContent string
		TagDesc    string
	}{}
	c.BindJSON(&data)
	log.Printf("receive post request: %v", data)
	tx := mysql.GormDB.Begin()
	if tag, err := service.AddTag(
		tx,
		data.TagContent,
		data.TagDesc); errx.New(err) != nil {
		if err := tx.Rollback().Error; errx.New(err) != nil {
			c.String(http.StatusInternalServerError,
				fmt.Sprintf("TX Rollback Error %s", err))
			return
		}
		c.JSON(http.StatusInternalServerError,
			gin.H{"status": -1, "tag": tag})
	} else {
		if err := tx.Commit().Error; errx.New(err) != nil {
			c.String(http.StatusInternalServerError,
				fmt.Sprintf("TX Commmit Error: %s", err))
			return
		}
		c.JSON(http.StatusOK,
			gin.H{"status": 0, "tag": tag})
	}
}

// DelOldTag is a func to add Tag
func DelOldTag(c *gin.Context) {
	data := &struct {
		TagID int
	}{}
	c.BindJSON(&data)
	log.Printf("receive post request: %v", data)
	if data.TagID == 0 {
		err := errors.New("no TagID in data field in post request")
		errx.New(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"status": -1})
		return
	}
	tx := mysql.GormDB.Begin()
	if err := service.DelTag(tx, data.TagID); errx.New(err) != nil {
		if err := tx.Rollback().Error; errx.New(err) != nil {
			c.String(http.StatusInternalServerError,
				fmt.Sprintf("TX Rollback Error %s", err))
			return
		}
		c.JSON(http.StatusInternalServerError,
			gin.H{"status": -1})
	} else {
		if err := tx.Commit().Error; errx.New(err) != nil {
			c.String(http.StatusInternalServerError,
				fmt.Sprintf("TX Commmit Error: %s", err))
			return
		}
		c.JSON(http.StatusOK,
			gin.H{"status": 0})
	}
}
