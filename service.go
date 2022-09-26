package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var repo ItemRepository

type Service struct {
	Repo ItemRepository
}

func NewService(repo ItemRepository) *Service {
	return &Service{
		Repo: repo,
	}
}

func GetItems(c *gin.Context) {
	items, err := repo.ListItems()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, items)
}

func PostItem(c *gin.Context) {
	var item Item
	err := c.BindJSON(&item)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// item, err = repo.CreateItem(item)
	// if err != nil {
	// 	log.Printf("Error creating item on database:\n%v\n", err)
	// 	c.AbortWithStatus(http.StatusBadRequest)
	// 	return
	// }
}
