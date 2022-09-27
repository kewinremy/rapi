package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "abc1234"
	dbname   = "postgres"
)

func checkError(err error) {
	if err != nil {
		log.Printf("Error: %v", err)
		log.Fatalf("Exiting...")
	}
}

func NewDatabase() *sql.DB {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	DB, err := sql.Open("postgres", psqlconn)
	checkError(err)

	// check db
	err = DB.Ping()
	checkError(err)

	fmt.Println("Connected!")
	return *&DB
}

func closeDatabase(db *sql.DB) {
	db.Close()
}
