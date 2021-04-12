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

// GetAllTag get Tag
func GetAllTag(c *gin.Context) {
	var tags model.Tags = model.Tags{}
	tx := mysql.GormDB.Begin()
	err := dao.GetAllTags(tx, &tags)
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

// AddTag add Tag
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

// UpdTag update Tag
func UpdTag(c *gin.Context) {
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

// DelTag add Tag
func DelTag(c *gin.Context) {
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
