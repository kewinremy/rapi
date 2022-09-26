package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {

	router = gin.Default()

	db := NewDatabase()
	repo := NewRepo(db)
	NewService(repo)

	mapRoutes()
	startServer()
}
