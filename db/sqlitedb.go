package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

//Createdb creates a db
func Createdb() {
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
