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
	rows, err := r.DB.Query("SELECT * FROM items_backup")
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
	result, err := r.DB.Exec("INSERT INTO items_backup (name, reservation_id) VALUES ($1, $2)", item.Name, item.ReservationId)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	idCreated, _ := result.LastInsertId()
	log.Printf("%d rows inserted\n", rowsAffected)
	log.Printf("Last id created: %d\n", idCreated)
	return nil
}
