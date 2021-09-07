package view

import (
	"assignment.com/collection-tree/model"
	"assignment.com/collection-tree/controller"
	"github.com/gin-gonic/gin"
)

func Fetch(c *gin.Context) {
	var data model.FetchData
	c.BindJSON(&data)
	webRequest, timeSpent := controller.GetData(data.CountryName, data.DeviceName)
	c.JSON(200, gin.H{
		"webRequest": webRequest,
		"timeSpent": timeSpent,
	})
}
