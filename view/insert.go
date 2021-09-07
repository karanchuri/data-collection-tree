package view

import (
	"assignment.com/collection-tree/model"
	"assignment.com/collection-tree/controller"
	"github.com/gin-gonic/gin"
)

func Insert(c *gin.Context) {
	var data model.InsertData
	c.BindJSON(&data)
	insertController := controller.DataCollectionController {
		Data: data,
	}
	insertController.RunTask()
	c.JSON(200, gin.H{
		"status": "Inserted",
	})
}
