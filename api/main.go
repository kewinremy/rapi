package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine
var db *sql.DB
var repo *Repo
var service *Service

func main() {

	router = gin.Default()

	db = NewDatabase()
	repo = NewRepo(db)
	service = NewService(repo)

	mapRoutes(service)
	startServer()
	closeDatabase(db)
}
