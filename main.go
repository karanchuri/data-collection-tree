package main

import (
	"assignment.com/collection-tree/view"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/health", view.Health)
	router.POST("/insert", view.Insert)
	router.POST("/fetch", view.Fetch)
	router.Run(":9000")
}
