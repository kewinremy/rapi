package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Service struct {
	repo ItemRepository
}

func NewService(repo *Repo) *Service {
	return &Service{
		repo: repo,
	}
}

func GetItems(c *gin.Context) {

	println(service)
	println(service.repo)

	// items, err := service.repo.ListItems()
	// if err != nil {
	// 	c.AbortWithStatus(http.StatusNotFound)
	// 	return
	// }

	// c.JSON(http.StatusOK, items)
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
