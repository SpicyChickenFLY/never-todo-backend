package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/SpicyChickenFLY/never-todo-backend/model"
	"github.com/SpicyChickenFLY/never-todo-backend/pkgs/mysql"
	"github.com/SpicyChickenFLY/never-todo-backend/service"
	"github.com/gin-gonic/gin"
	"github.com/lingdor/stackerror"
)

// GetAll get all todo info
func GetAll(c *gin.Context) {
	tx := mysql.GormDB.Begin()
	var fullTasks model.FullTasks
	var tags model.Tags
	err1 := service.GetAllFullTasks(tx, &fullTasks)
	err2 := service.GetAllTags(tx, &tags)
	var err error
	if err1 != nil {
		err = mysql.CheckTransaction(tx, err1)
	} else if err2 != nil {
		err = mysql.CheckTransaction(tx, err2)
	} else {
		err = nil
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"status": -1})
		log.Print(err)
	} else {
		c.JSON(http.StatusOK,
			gin.H{"status": 0, "result": gin.H{
				"fullTasks": fullTasks,
				"tags":      tags,
			}})
	}
}

// GetAllFullTask is a func to
func GetAllFullTask(c *gin.Context) {
	tx := mysql.GormDB.Begin()
	var fullTasks model.FullTasks
	err := service.GetAllFullTasks(tx, &fullTasks)
	err = mysql.CheckTransaction(tx, err)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"status": -1})
		log.Print(err)
	} else {
		c.JSON(http.StatusOK,
			gin.H{"status": 0, "result": fullTasks})
	}
}

// GetFullTasksByContent get FullTask by Name(no need to be accurate)
func GetFullTasksByContent(c *gin.Context) {
	var fullTasks model.FullTasks
	content := c.Param("content")
	tx := mysql.GormDB.Begin()
	err := service.GetFullTasksByContent(tx, &fullTasks, content)
	err = mysql.CheckTransaction(tx, err)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"status": -1})
		log.Print(err)
	} else {
		c.JSON(http.StatusOK,
			gin.H{"status": 0, "result": fullTasks})
	}
}

// GetFullTaskByTag get FullTask by TagID
func GetFullTaskByTag(c *gin.Context) {
	var fullTasks model.FullTasks
	tagID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		stackerror.New(err.Error())
		log.Println(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"status": -1})
		log.Println(err)
	}
	tx := mysql.GormDB.Begin()
	err = service.GetFullTasksByTag(tx, &fullTasks, tagID)
	err = mysql.CheckTransaction(tx, err)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"status": -1})
		log.Println(err)
	} else {
		c.JSON(http.StatusOK,
			gin.H{"status": 0, "result": fullTasks})
	}
}

// AddFullTask is a func to add Task
func AddFullTask(c *gin.Context) {
	data := &struct {
		TaskContent string `json:"taskContent"`
		TagsID      []int  `json:"tagsID"`
	}{}
	c.BindJSON(&data)
	log.Printf("receive post request: %v", data)
	task := model.Task{Content: data.TaskContent}
	tx := mysql.GormDB.Begin()
	err := service.AddFullTask(
		tx, &task, data.TagsID)
	err = mysql.CheckTransaction(tx, err)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"status": -1})
		log.Println(err)
	} else {
		c.JSON(http.StatusOK,
			gin.H{"status": 0, "result": task})
	}
}

// DelFullTask is a func to delete Task
func DelFullTask(c *gin.Context) {
	taskID, err := strconv.Atoi(c.Param("task_id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"status": -1})
		return
	}
	log.Printf("receive post request: %d", taskID)
	tx := mysql.GormDB.Begin()
	err = service.DelFullTask(tx, taskID)
	err = mysql.CheckTransaction(tx, err)
	if err != nil {
		c.JSON(http.StatusOK,
			gin.H{"status": -1})
		log.Print(err)
	} else {
		c.JSON(http.StatusOK,
			gin.H{"status": 0})
	}
}

// UpdFullTask is a func to update Task
func UpdFullTask(c *gin.Context) {
	taskID, err := strconv.Atoi(c.Param("task_id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"status": -1})
		return
	}
	data := &struct {
		TaskContent    string `json:"taskContent"`
		TaskCompeleted bool   `json:"taskCompeleted"`
		TagsID         []int  `json:"tagsID"`
	}{}
	c.BindJSON(&data)
	task := model.Task{
		ID:         taskID,
		Content:    data.TaskContent,
		Compeleted: data.TaskCompeleted,
	}
	log.Printf("receive post request: %v", data)
	tx := mysql.GormDB.Begin()
	err = service.UpdFullTask(tx, &task, data.TagsID)
	err = mysql.CheckTransaction(tx, err)
	if err != nil {
		c.JSON(http.StatusOK,
			gin.H{"status": -1})
		log.Print(err)
	} else {
		c.JSON(http.StatusOK,
			gin.H{"status": 0, "result": task})
	}
}

// GetAllTask is a func to get Task
func GetAllTask(c *gin.Context) {
	var tasks model.Tasks
	tx := mysql.GormDB.Begin()
	err := service.GetAllTasks(tx, &tasks)
	result := gin.H{
		"tasks": tasks,
	}
	err = mysql.CheckTransaction(tx, err)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"status": -1})
		log.Print(err)
	} else {
		c.JSON(http.StatusOK,
			gin.H{"status": 0, "result": result})
	}
}

// GetAllTag is a func to get Tag
func GetAllTag(c *gin.Context) {
	var tags model.Tags = model.Tags{}
	tx := mysql.GormDB.Begin()
	err := service.GetAllTags(tx, &tags)
	result := gin.H{
		"tags": tags,
	}
	err = mysql.CheckTransaction(tx, err)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"status": -1})
		log.Print(err)
	} else {
		c.JSON(http.StatusOK,
			gin.H{"status": 0, "result": result})
	}
}

// AddTag is a func to add Tag
func AddTag(c *gin.Context) {
	data := &struct {
		Content string `json:"content"`
		Desc    string `json:"desc"`
		Color   string `json:"color"`
	}{}
	c.BindJSON(&data)
	log.Printf("receive post request: %v", data)
	tag := model.Tag{
		Content: data.Content,
		Desc:    data.Desc,
		Color:   data.Color,
	}
	tx := mysql.GormDB.Begin()
	err := service.AddTag(tx, tag)
	err = mysql.CheckTransaction(tx, err)
	if err != nil {
		c.JSON(http.StatusOK,
			gin.H{"status": -1})
		log.Print(err)
	} else {
		c.JSON(http.StatusOK,
			gin.H{"status": 0, "result": tag})
	}
}

// DelTag is a func to add Tag
func DelTag(c *gin.Context) {
	tagID, err := strconv.Atoi(c.Param("tag_id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"status": -1})
		return
	}
	tx := mysql.GormDB.Begin()
	err = service.DelTag(tx, tagID)
	err = mysql.CheckTransaction(tx, err)
	if err != nil {
		c.JSON(http.StatusOK,
			gin.H{"status": -1})
		log.Print(err)
	} else {
		c.JSON(http.StatusOK,
			gin.H{"status": 0})
	}
}

// UpdTag is a func to update Tag
func UpdTag(c *gin.Context) {
	tagID, err := strconv.Atoi(c.Param("tag_id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"status": -1})
		return
	}
	data := &struct {
		Content string `json:"content"`
		Desc    string `json:"desc"`
		Color   string `json:"color"`
	}{}
	c.BindJSON(&data)
	tag := model.Tag{
		ID:      tagID,
		Content: data.Content,
		Desc:    data.Desc,
		Color:   data.Color,
	}
	log.Printf("receive post request: %v", data)
	tx := mysql.GormDB.Begin()
	err = service.UpdTag(tx, &tag)
	err = mysql.CheckTransaction(tx, err)
	if err != nil {
		c.JSON(http.StatusOK,
			gin.H{"status": -1})
		log.Print(err)
	} else {
		c.JSON(http.StatusOK,
			gin.H{"status": 0, "result": tag})
	}
}

// GetLastSyncTime fetch latest sync time of server
func GetLastSyncTime(c *gin.Context) {
	// c.JSON(http.StatusOK,
	// 	gin.H{"status": 0, "result": tag})
}

// SyncFromServer get sync content from remote
func SyncFromServer(c *gin.Context) {

}

// SyncToServer set sync content from local
func SyncToServer(c *gin.Context) {

}
