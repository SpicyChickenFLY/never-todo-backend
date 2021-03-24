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

// GetAll search for all tables and return
// url:/todo/getall
func GetAll(c *gin.Context) {
	var tasks model.Tasks = model.Tasks{}
	var tags model.Tags = model.Tags{}
	var taskTags model.TaskTags = model.TaskTags{}
	tx := mysql.GormDB.Begin()
	err := service.GetAllTables(tx, &tasks, &tags, &taskTags)
	result := gin.H{
		"tasks":         tasks,
		"tags":          tags,
		"task_tag_pair": taskTags,
	}
	err = mysql.CheckTransaction(tx, err)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"Status": -1, "Result": result})
		log.Print(err)
	} else {
		c.JSON(http.StatusOK,
			gin.H{"Status": 0, "Result": result})
	}

}

// GetAllFullTask is a func to
func GetAllFullTask(c *gin.Context) {
	tx := mysql.GormDB.Begin()
	var fullTasks model.FullTasks
	err := service.GetFullTasks(tx, &fullTasks)
	err = mysql.CheckTransaction(tx, err)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"Status": -1, "Result": fullTasks})
		log.Print(err)
	} else {
		c.JSON(http.StatusOK,
			gin.H{"Status": 0, "Result": fullTasks})
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
			gin.H{"Status": -1, "Result": fullTasks})
		log.Print(err)
	} else {
		c.JSON(http.StatusOK,
			gin.H{"Status": 0, "Result": fullTasks})
	}
}

// GetFullTaskByTag get FullTask by TagID
func GetFullTaskByTag(c *gin.Context) {
	var fullTasks model.FullTasks
	tagID, err := strconv.Atoi(c.Param("tag_id"))
	if err != nil {
		stackerror.New(err.Error())
		log.Println(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"Status": -1, "Result": fullTasks})
		log.Println(err)
	}
	tx := mysql.GormDB.Begin()
	err = service.GetFullTasksByTag(tx, &fullTasks, tagID)
	err = mysql.CheckTransaction(tx, err)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"Status": -1, "Result": fullTasks})
		log.Println(err)
	} else {
		c.JSON(http.StatusOK,
			gin.H{"Status": 0, "Result": fullTasks})
	}
}

// AddFullTask is a func to add Task
func AddFullTask(c *gin.Context) {
	data := &struct {
		TaskContent string `json:"TaskContent"`
		TagsID      []int  `json:"TagsID"`
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
			gin.H{"Status": -1, "Result": task})
		log.Println(err)
	} else {
		c.JSON(http.StatusOK,
			gin.H{"Status": 0, "Result": task})
	}
}

// DelFullTask is a func to delete Task
func DelFullTask(c *gin.Context) {
	taskID, err := strconv.Atoi(c.Param("task_id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"Status": -1})
		return
	}
	log.Printf("receive post request: %d", taskID)
	tx := mysql.GormDB.Begin()
	err = service.DelFullTask(tx, taskID)
	err = mysql.CheckTransaction(tx, err)
	if err != nil {
		c.JSON(http.StatusOK,
			gin.H{"Status": -1})
		log.Print(err)
	} else {
		c.JSON(http.StatusOK,
			gin.H{"Status": 0})
	}
}

// UpdFullTask is a func to update Task
func UpdFullTask(c *gin.Context) {
	data := &struct {
		TaskID      int    `json:"TaskID"`
		TaskContent string `json:"TaskContent"`
		TagsID      []int  `json:"TagsID"`
	}{}
	c.BindJSON(&data)
	task := model.Task{
		ID:      data.TaskID,
		Content: data.TaskContent,
	}
	log.Printf("receive post request: %v", data)
	tx := mysql.GormDB.Begin()
	err := service.UpdFullTask(tx, &task, data.TagsID)
	err = mysql.CheckTransaction(tx, err)
	if err != nil {
		c.JSON(http.StatusOK,
			gin.H{"Status": -1, "Result": task})
		log.Print(err)
	} else {
		c.JSON(http.StatusOK,
			gin.H{"Status": 0, "Result": task})
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
			gin.H{"Status": -1, "Result": result})
		log.Print(err)
	} else {
		c.JSON(http.StatusOK,
			gin.H{"Status": 0, "Result": result})
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
			gin.H{"Status": -1, "Result": result})
		log.Print(err)
	} else {
		c.JSON(http.StatusOK,
			gin.H{"Status": 0, "Result": result})
	}
}

// AddTag is a func to add Tag
func AddTag(c *gin.Context) {
	data := &struct {
		TagContent string `json:"TagContent"`
		TagDesc    string `json:"TagDesc"`
	}{}
	c.BindJSON(&data)
	log.Printf("receive post request: %v", data)
	tag := model.Tag{Content: data.TagContent, Desc: data.TagDesc}
	tx := mysql.GormDB.Begin()
	err := service.AddTag(tx, tag)
	err = mysql.CheckTransaction(tx, err)
	if err != nil {
		c.JSON(http.StatusOK,
			gin.H{"Status": -1, "Result": tag})
		log.Print(err)
	} else {
		c.JSON(http.StatusOK,
			gin.H{"Status": 0, "Result": tag})
	}
}

// DelTag is a func to add Tag
func DelTag(c *gin.Context) {
	tagID, err := strconv.Atoi(c.Param("tag_id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"Status": -1})
		return
	}
	tx := mysql.GormDB.Begin()
	err = service.DelTag(tx, tagID)
	err = mysql.CheckTransaction(tx, err)
	if err != nil {
		c.JSON(http.StatusOK,
			gin.H{"Status": -1})
		log.Print(err)
	} else {
		c.JSON(http.StatusOK,
			gin.H{"Status": 0})
	}
}

// UpdTag is a func to update Tag
func UpdTag(c *gin.Context) {
	data := &struct {
		TagID      int    `json:"TagID"`
		TagContent string `json:"TagContent"`
		TagDesc    string `json:"TagDesc"`
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
	err = mysql.CheckTransaction(tx, err)
	if err != nil {
		c.JSON(http.StatusOK,
			gin.H{"Status": -1, "Result": tag})
		log.Print(err)
	} else {
		c.JSON(http.StatusOK,
			gin.H{"Status": 0, "Result": tag})
	}
}

// GetLastSyncTime fetch latest sync time of server
func GetLastSyncTime(c *gin.Context) {

}

// SyncFromServer get sync content from remote
func SyncFromServer(c *gin.Context) {

}

// SyncToServer set sync content from local
func SyncToServer(c *gin.Context) {

}
