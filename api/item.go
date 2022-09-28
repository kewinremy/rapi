package main

type Item struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	ReservationId int    `json:"reservation_id"`
}

func (i *Item) TableName() string {
	return "items"
}
