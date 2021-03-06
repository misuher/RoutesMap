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
	GetDaily(date time.Time) ([]*Flight, error)
	SetDaily(date time.Time, flight Flight) error
	DeleteDaily(date time.Time) error

	CreateParking()
	GetParkings(date time.Time) ([]TimesInLPA, error)
	GetParking(date time.Time, aircraft string) (TimesInLPA, error)
	SetParking(date time.Time) error
	filterTimesInLPA(date time.Time) ([]TimesInLPA, error)
	DeleteParking(date time.Time) error
}

//DB is a wrapper por a sql.DB
type DB struct {
	*sql.DB
}

//Open the database returning a DB object
func Open() (*DB, error) {
	db, err := sql.Open("sqlite3", "./Daily2.db")
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &DB{db}, nil
}
