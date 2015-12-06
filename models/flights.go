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
	CREATE TABLE IF NOT EXISTS daily (
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

//GetDaily read an entire DAYLY.pdf from the db
func (db *DB) GetDaily(date time.Time) ([]*Flight, error) {
	err := db.Ping()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT * FROM daily WHERE date = ? ORDER BY id DESC", date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var flights []*Flight
	for rows.Next() {
		flight := new(Flight)
		err := rows.Scan(&flight.Avion, &flight.Vuelo, &flight.DepPlace, &flight.DepTime, &flight.ArrPlace, &flight.ArrTime, &flight.Pasajeros, &flight.Capitan, &flight.PrimerOficial)
		if err != nil {
			return nil, err
		}
		log.Printf("Vuelo creado: %s %s  %s  %s  %s  %s  %s  %s  %s\n", flight.Avion, flight.Vuelo, flight.DepPlace, flight.DepTime, flight.ArrPlace, flight.ArrTime, flight.Pasajeros, flight.Capitan, flight.PrimerOficial)
		flights = append(flights, flight)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return flights, nil
}

//SetDaily saves a DAYLY.pdf parsed file into the db
func (db *DB) SetDaily(date time.Time, flight Flight) error {
	err := db.Ping()
	if err != nil {
		log.Println("setDaily error: db.Ping")
		return err
	}

	result, err := db.Exec("INSERT INTO daily (date, avion,	vuelo,	depplace , deptime , arrplace, arrtime,	pasajeros, capitan,	fo) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", date, flight.Avion, flight.Vuelo, flight.DepPlace, flight.DepTime, flight.ArrPlace, flight.ArrTime, flight.Pasajeros, flight.Capitan, flight.PrimerOficial)
	if err != nil {
		log.Println("setDaily error: db.Exec")
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("setDaily error: RowsAffected")
		return err
	}
	log.Printf("Vuelo %s creado (%d row cambiado)\n", flight.Vuelo, rowsAffected)
	return nil
}

func (db *DB) DeleteDaily(date time.Time) error {
	err := db.Ping()
	if err != nil {
		log.Println("DeleteDaily error: db.Ping")
		return err
	}

	result, err := db.Exec("DELETE FROM daily WHERE date=?", date)
	if err != nil {
		log.Println("DeleteDaily error: db.Exec")
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("DeleteDaily error: RowsAffected")
		return err
	}
	log.Printf("Vuelo borrado (%d row cambiado)\n", rowsAffected)
	return nil
}
