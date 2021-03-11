package controller

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ShowUI is a func for user to custom their request
func ShowUI(c *gin.Context) {

	c.HTML(http.StatusOK, "ui.tmpl", gin.H{
		"title": template.HTML("NeverTODO"),
	})
}
