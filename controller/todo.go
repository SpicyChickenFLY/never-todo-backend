package controller

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lingdor/stackerror"
	"spicychicken.top/NeverTODO/backend/model"
	"spicychicken.top/NeverTODO/backend/pkgs/mysql"
	"spicychicken.top/NeverTODO/backend/service"
)

// FIXME: all err should be pocessed in controllers

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
	err = mysql.StopTransaction(tx, err)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"status": -1, "result": result})
		log.Print(err)
	} else {
		c.JSON(http.StatusOK,
			gin.H{"status": 0, "result": result})
	}

}

// GetTaskList is a func to
func GetTaskList(c *gin.Context) {
	tx := mysql.GormDB.Begin()
	var fullTasks model.FullTasks
	err := service.GetFullTasks(tx, &fullTasks)
	err = mysql.StopTransaction(tx, err)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"status": -1, "fullTasks": fullTasks})
		log.Print(err)
	} else {
		c.JSON(http.StatusOK,
			gin.H{"status": 0, "fullTasks": fullTasks})
	}
}

// GetTaskListByTag is a func to
func GetTaskListByTag(c *gin.Context) {
	var fullTasks model.FullTasks
	tagIDStr := c.Query("tag_id")
	if tagIDStr == "" {
		GetTaskList(c)
		return
	}
	tagID, err := strconv.Atoi(tagIDStr)
	if err != nil {
		stackerror.New(err.Error())
		log.Println(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"status": -1, "fullTasks": fullTasks})
		log.Println(err)
	}
	tx := mysql.GormDB.Begin()
	err = service.GetFullTasksByTag(tx, &fullTasks, tagID)
	err = mysql.StopTransaction(tx, err)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"status": -1, "fullTasks": fullTasks})
		log.Println(err)
	} else {
		c.JSON(http.StatusOK,
			gin.H{"status": 0, "fullTasks": fullTasks})
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
	task, err := service.AddFullTask(
		tx, data.TaskContent, data.TagsID)
	err = mysql.StopTransaction(tx, err)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"status": -1, "task": task})
		log.Println(err)
	} else {
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
		err := stackerror.New("no TaskID in data field in post request")
		log.Println(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"status": -1})
		return
	}
	log.Printf("receive post request: %d", data.TaskID)
	tx := mysql.GormDB.Begin()
	err := service.DelFullTask(tx, data.TaskID)
	err = mysql.StopTransaction(tx, err)
	if err != nil {
		c.JSON(http.StatusOK,
			gin.H{"status": -1})
		log.Print(err)
	} else {
		c.JSON(http.StatusOK,
			gin.H{"status": 0})
	}
}

// UpdOldTask is a func to update Task
func UpdOldTask(c *gin.Context) {
	data := &struct {
		TaskID      int
		TaskContent string
		TagsID      []int
	}{}
	c.BindJSON(&data)
	task := model.Task{
		ID:      data.TaskID,
		Content: data.TaskContent,
	}
	log.Printf("receive post request: %v", data)
	tx := mysql.GormDB.Begin()
	err := service.UpdFullTask(tx, &task, data.TagsID)
	err = mysql.StopTransaction(tx, err)
	if err != nil {
		c.JSON(http.StatusOK,
			gin.H{"status": -1, "Task": task})
		log.Print(err)
	} else {
		c.JSON(http.StatusOK,
			gin.H{"status": 0, "Task": task})
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
	tag, err := service.AddTag(
		tx, data.TagContent, data.TagDesc)
	err = mysql.StopTransaction(tx, err)
	if err != nil {
		c.JSON(http.StatusOK,
			gin.H{"status": -1, "tag": tag})
		log.Print(err)
	} else {
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
		log.Println(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"status": -1})
		return
	}
	tx := mysql.GormDB.Begin()
	err := service.DelTag(tx, data.TagID)
	err = mysql.StopTransaction(tx, err)
	if err != nil {
		c.JSON(http.StatusOK,
			gin.H{"status": -1})
		log.Print(err)
	} else {
		c.JSON(http.StatusOK,
			gin.H{"status": 0})
	}
}

// UpdOldTag is a func to update Tag
func UpdOldTag(c *gin.Context) {
	data := &struct {
		TagID      int
		TagContent string
		TagDesc    string
	}{}
	c.BindJSON(&data)
	tag := model.Tag{
		ID:      data.TagID,
		Content: data.TagContent,
		Desc:    data.TagDesc,
	}
	log.Printf("receive post request: %v", data)
	tx := mysql.GormDB.Begin()
	err := service.UpdTag(tx, &tag)
	err = mysql.StopTransaction(tx, err)
	if err != nil {
		c.JSON(http.StatusOK,
			gin.H{"status": -1, "Tag": tag})
		log.Print(err)
	} else {
		c.JSON(http.StatusOK,
			gin.H{"status": 0, "Tag": tag})
	}
}
