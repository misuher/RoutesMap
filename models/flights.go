package models

import (
	"log"
	"time"
)

//Flight is a row of data parsed holder
type Flight struct {
	Avion         string
	Vuelo         string
	DepPlace      string
	DepTime       time.Time
	ArrPlace      string
	ArrTime       time.Time
	Pasajeros     string
	Capitan       string
	PrimerOficial string
}

//Flights hold all the rows in DAILY.pdf report
type Flights struct {
	f []Flight
}

//Create initialize a table in the db with Flights
func (db *DB) Create() {
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS Daily (
		id integer not null PRIMARY KEY,
		date DATE NOT NULL,
		avion VARCHAR(255) NOT NULL,
		vuelo VARCHAR(255) NOT NULL,
		depplace VARCHAR(255) NOT NULL,
		deptime TIME NOT NULL,
		arrplace VARCHAR(255) NOT NULL,
		arrtime TIME NOT NULL,
		pasajeros VARCHAR(255) NOT NULL,
		capitan VARCHAR(255) NOT NULL,
		fo VARCHAR(255) NOT NULL
		);
	`

	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}
