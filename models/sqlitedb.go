package models

import (
	"database/sql"
	"time"

	//imported library to invoque its init() fuction
	_ "github.com/mattn/go-sqlite3"
)

//Datastore defines personaliced querys for the models
type Datastore interface {
	Create()
	getDaily(date time.Time) ([]*Flight, error)
	setDaily(date time.Time, flight Flight) error
}

//DB is a wrapper por a sql.DB
type DB struct {
	*sql.DB
}

//Open the database returning a DB object
func Open() (*DB, error) {
	db, err := sql.Open("sqlite3", "./Daily.db")
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &DB{db}, nil
}
