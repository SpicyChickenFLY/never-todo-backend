package todo

import (
	"log"
	"net/http"

	"github.com/SpicyChickenFLY/never-todo-backend/model"
	"github.com/SpicyChickenFLY/never-todo-backend/pkgs/mysql"
	"github.com/SpicyChickenFLY/never-todo-backend/service"
	"github.com/gin-gonic/gin"
)

// GetAll get all todo info
func GetAll(c *gin.Context) {
	tx := mysql.GormDB.Begin()
	var fullTasks model.FullTasks
	var tasks model.Tasks
	var tags model.Tags
	var taskTags model.TaskTags

	err := service.GetAllTables(tx, &tasks, &tags, &taskTags)

	if err != nil {
		err = mysql.CheckTransaction(tx, err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"status": -1})
		log.Print(err)
	} else {
		c.JSON(http.StatusOK,
			gin.H{"status": 0, "result": gin.H{
				"fullTasks": fullTasks,
				"tags":      tags,
				"tasks":     tasks,
				"task_tags": taskTags,
			}})
	}
}
