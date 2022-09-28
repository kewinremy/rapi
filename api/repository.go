package main

import (
	"database/sql"
	"log"
)

type Repo struct {
	DB *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{DB: db}
}

func (r *Repo) ListItems() ([]Item, error) {
	rows, err := r.DB.Query("SELECT * FROM items")
	if err != nil {
		return nil, err
	}

	var items []Item
	var itemCounter int
	for rows.Next() {
		item := Item{}
		err := rows.Scan(&item.Id, &item.Name, &item.ReservationId)
		if err != nil {
			return nil, err
		}

		items = append(items, item)
		itemCounter++
	}

	log.Printf("Got %d rows\n", itemCounter)
	return items, nil
}

func (r *Repo) CreateItem(item Item) error {
	result, err := r.DB.Exec("INSERT INTO items (name, reservation_id) VALUES ($1, $2)", item.Name, item.ReservationId)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	log.Printf("%d rows inserted\n", rowsAffected)
	return nil
}
