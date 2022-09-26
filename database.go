package main

import (
	"database/sql"
	"fmt"

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
		panic(err)
	}
}

func initPostgres() {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	DB, err := sql.Open("postgres", psqlconn)
	checkError(err)

	// close database
	defer DB.Close()

	// check db
	err = DB.Ping()
	checkError(err)

	fmt.Println("Connected!")
}
