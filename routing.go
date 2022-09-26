package main

import (
	"github.com/gin-gonic/gin"
)

func ItemRouting(r *gin.RouterGroup) {
	r.GET("/items", GetItems)
	r.POST("/items", PostItem)
}
