package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"spicychicken.top/NeverTODO/backend/pkgs/request"
)

// ShowUI is a func for user to custom their request
func ShowUI(c *gin.Context) {
	c.HTML(http.StatusOK, "", gin.H{})
}

// SendRequest is a func to send http request by user
func SendRequest(c *gin.Context) {

}

// SimAddTask simulate to post task add request and data
// {
// 	"TaskContent": "xxx",
// 	"TagsID": [1, 3]
// }
func SimAddTask(c *gin.Context) {
	task := map[string]interface{}{
		"TaskContent": "扔掉 MacBook",
		"TagsID":      []int{1, 3},
	}
	response := request.Post(
		"http://localhost:8080/todo/task/add",
		task, "application/json")
	c.String(http.StatusOK, response)

}

// SimDelTask simulate to post task delete request and data
// {
// 	"TaskID": 1
// }
func SimDelTask(c *gin.Context) {
	taskID := map[string]interface{}{
		"TaskID": 1,
	}
	response := request.Post(
		"http://localhost:8080/todo/task/del",
		taskID, "application/json")
	c.String(http.StatusOK, response)
}

// SimAddTag simulate to post tag add request and data
func SimAddTag(c *gin.Context) {
	tag := map[string]interface{}{
		"TagContent": "电脑",
		"TagDesc":    "电刑不是早废除了吗？",
	}
	response := request.Post(
		"http://localhost:8080/todo/tag/add",
		tag, "application/json")
	c.String(http.StatusOK, response)

}

// SimDelTag simulate to post tag delete request and data
func SimDelTag(c *gin.Context) {
	tagID := map[string]interface{}{
		"TagID": 1,
	}
	response := request.Post(
		"http://localhost:8080/todo/tag/del",
		tagID, "application/json")
	c.String(http.StatusOK, response)
}
