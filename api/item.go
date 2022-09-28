package main

type Item struct {
	Id            uint   `json:"id"`
	Name          string `json:"name"`
	ReservationId string `json:"reservation_id"`
}

func (i *Item) TableName() string {
	return "items_backup"
}
