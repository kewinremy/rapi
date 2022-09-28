package main

import (
	"github.com/gin-gonic/gin"
)

func mapRoutes() {
	// Ping route, health control.
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// API Version 1 group (al routes should start with /v1/*.
	v1 := router.Group("/v1")

	// Map entities routing.
	ItemRouting(v1)
}

func ItemRouting(r *gin.RouterGroup) {
	r.GET("/items", GetItems)
	r.POST("/items", PostItem)
}
