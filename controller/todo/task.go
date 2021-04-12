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

// GetAllTask get all tasks
func GetAllTask(c *gin.Context) {
	var tasks model.Tasks
	tx := mysql.GormDB.Begin()
	err := dao.GetAllTasks(tx, &tasks)
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

// AddTask add Task
func AddTask(c *gin.Context) {
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
	err := dao.AddTag(tx, &tag)
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

// UpdTask update Task
func UpdTask(c *gin.Context) {
	tagID, err := strconv.Atoi(c.Param("id"))
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
	err = dao.UpdTag(tx, &tag)
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

// DelTask add Task
func DelTask(c *gin.Context) {
	tagID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"status": -1})
		return
	}
	tx := mysql.GormDB.Begin()
	err = dao.DelTag(tx, tagID)
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
