package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type Service struct {
	repo ItemRepository
}

func NewService(iRepo ItemRepository) *Service {
	return &Service{
		repo: iRepo,
	}
}

func (s *Service) GetItems(c *gin.Context) {
	items, err := service.repo.ListItems()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, items)
}

func AsyncPostReservationId(gc *gin.Context, item Item) {
	c := make(chan int)
	wg := new(sync.WaitGroup)
	wg.Add(1)

	go func() {
		wg.Wait()
		close(c)
	}()

	go ReservationService(item.Name, wg, c)

	reservationIdChannel := <-c
	log.Println("reservationIdChannel: ", reservationIdChannel)
	item.ReservationId = int(reservationIdChannel)

	err := service.repo.CreateItem(item)
	if err != nil {
		log.Printf("Error creating item on database:\n%v\n", err)
		gc.AbortWithStatus(http.StatusBadRequest)
	}
}

func (s *Service) PostItem(c *gin.Context) {
	var item Item
	err := c.BindJSON(&item)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	go AsyncPostReservationId(c, item)
}
