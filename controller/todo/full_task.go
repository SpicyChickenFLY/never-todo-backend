package todo

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

// AddFullTask add Task
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

// DelFullTask delete Task
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

// UpdFullTask update Task
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
