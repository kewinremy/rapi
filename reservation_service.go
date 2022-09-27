package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func ReservationService(itemName string, wg *sync.WaitGroup, channel chan int) {
	defer wg.Done()

	log.Printf("Reserving %s \n", itemName)
	log.Printf("Sleeping for 10 seconds... \n")
	time.Sleep(10 * time.Second)

	channel <- rand.Intn(100)
}
