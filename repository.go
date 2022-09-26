package main

import (
	"database/sql"
	"log"
)

type Repo struct {
	DB *sql.DB
}

func (ur *Repo) ListItems() ([]Item, error) {
	rows, err := ur.DB.Query("SELECT * FROM items_backup")
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

func (ur *Repo) CreateItem(item Item) error {
	// result, err := ur.DB.Exec("INSERT INTO items (name, email, phone, address) VALUES ($1, $2, $3, $4)", item.Name, item.Email, item.Phone, item.Address)
	// if err != nil {
	// 	return model.User{}, err
	// }

	// rowsAffected, _ := result.RowsAffected()
	// idCreated, _ := result.LastInsertId()
	// log.Printf("%d rows inserted\n", rowsAffected)

	// item.Id = uint(idCreated)

	return nil
}
