package todo

import (
	"log"
	"net/http"
	"strconv"

	"github.com/SpicyChickenFLY/never-todo-backend/dao"
	"github.com/SpicyChickenFLY/never-todo-backend/model"
	"github.com/SpicyChickenFLY/never-todo-backend/pkgs/mysql"
	"github.com/gin-gonic/gin"
)

// GetAllTaskTag get all tasks
func GetAllTaskTag(c *gin.Context) {
	var taskTags model.TaskTags
	tx := mysql.GormDB.Begin()
	err := dao.GetAllTaskTags(tx, &taskTags)
	result := gin.H{
		"task_tags": taskTags,
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

// AddTaskTag add TaskTag
func AddTaskTag(c *gin.Context) {
	data := &struct {
		TaskID int `json:"task_id"`
		TagID  int `json:"tag_id"`
	}{}
	c.BindJSON(&data)
	log.Printf("receive post request: %v", data)
	taskTag := model.TaskTag{
		TaskID: data.TaskID,
		TagID:  data.TagID,
	}
	tx := mysql.GormDB.Begin()
	err := dao.AddTaskTag(tx, &taskTag)
	err = mysql.CheckTransaction(tx, err)
	if err != nil {
		c.JSON(http.StatusOK,
			gin.H{"status": -1})
		log.Print(err)
	} else {
		c.JSON(http.StatusOK,
			gin.H{"status": 0, "result": taskTag})
	}
}

// DelTaskTag delete TaskTag by id
func DelTaskTag(c *gin.Context) {
	taskTagID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"status": -1})
		return
	}
	tx := mysql.GormDB.Begin()
	err = dao.DelTaskTag(tx, taskTagID)
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
