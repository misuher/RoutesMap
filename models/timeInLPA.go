package models

import (
	"log"
	"sort"
	"time"
)

//timeInLPA defines the interval of time represented as 2 hours in which an aircraft is parked in LPA airport
type TimeInLPA struct {
	Arrival string
	Leave   string
}

//TimesInLPA is a collection of timeInLPA just in case that the aircraft stops several intervals of time in one day
type TimesInLPA struct {
	Aircraft string
	Times    []TimeInLPA
}

//CreateParking initialize a table in the db with Flights
func (db *DB) CreateParking() {
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS parking (
		id integer not null PRIMARY KEY,
		date DATE NOT NULL,
		aircraft VARCHAR(255) NOT NULL,
		arrive VARCHAR(255) NOT NULL,
		leave VARCHAR(255) NOT NULL
		);
	`

	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}

func (db *DB) GetParkings(date time.Time) ([]TimesInLPA, error) {
	var result []TimesInLPA
	err := db.Ping()
	if err != nil {
		return result, err
	}

	aircrafts := []string{"EC-GRU", "EC-GRP", "EC-GQF", "EC-LZR", "EC-LYZ"}
	for _, val := range aircrafts {
		r, err := db.GetParking(date, val)
		if err != nil {
			return nil, err
		}
		result = append(result, r)
	}
	return result, nil
}

//GetParking obtain parking times of one aircraft
func (db *DB) GetParking(date time.Time, aircraft string) (TimesInLPA, error) {
	var result TimesInLPA
	err := db.Ping()
	if err != nil {
		return result, err
	}

	rows, err := db.Query("SELECT arrive,leave FROM parking WHERE date = ? AND aircraft = ? ORDER BY id DESC", date, aircraft)
	if err != nil {
		return result, err
	}
	defer rows.Close()

	var parkings []TimeInLPA
	for rows.Next() {
		parking := new(TimeInLPA)
		err := rows.Scan(&parking.Arrival, &parking.Leave)
		if err != nil {
			return result, err
		}
		log.Printf("Parking leido: %s %s %s  %s\n", date, aircraft, parking.Arrival, parking.Leave)
		parkings = append(parkings, *parking)
	}
	if err = rows.Err(); err != nil {
		return result, err
	}
	result.Aircraft = aircraft
	result.Times = parkings
	return result, nil
}

//SetParking save times of aircrafts in LPA airport in the db
func (db *DB) SetParking(date time.Time) error {
	//filter which flights arrive or leave LPA
	aircrafts, err := db.filterTimesInLPA(date)
	if err != nil {
		log.Println("SetParking error: db.filterTimesLPA")
		return err
	}

	err = db.DeleteParking(date)
	if err != nil {
		log.Println("SetParking error: db.DeleteParking")
		return err
	}
	//save result into db parking table
	for _, aircraft := range aircrafts {
		for _, time := range aircraft.Times {
			result, err := db.Exec("INSERT INTO parking (date, aircraft , arrive , leave) VALUES(?, ?, ?, ?)", date, aircraft.Aircraft, time.Arrival, time.Leave)
			if err != nil {
				log.Println("setParking error: db.Exec")
				return err
			}
			rowsAffected, err := result.RowsAffected()
			if err != nil {
				log.Println("setParking error: RowsAffected")
				return err
			}
			log.Printf("Parking %s creado (%d row cambiado)\n", aircraft.Aircraft, rowsAffected)
		}
	}

	return nil
}

func (db *DB) filterTimesInLPA(date time.Time) ([]TimesInLPA, error) {

	//get all the daily flights
	flights, err := db.GetDaily(date)
	if err != nil {
		log.Println("filterTimesInLPA error: db.GetDaily")
		return nil, err
	}

	//filter flights based on aircraft and place
	var GRU []string
	var GRP []string
	var GQF []string
	var LZR []string
	var LYZ []string

	for _, val := range flights {
		switch val.Avion {
		case "EC-GRU":
			if val.ArrPlace == "LPA" {
				GRU = append(GRU, val.ArrTime)
			} else if val.DepPlace == "LPA" && len(GRU) != 0 {
				GRU = append(GRU, val.DepTime)
			}
			break
		case "EC-GRP":
			if val.ArrPlace == "LPA" {
				GRP = append(GRP, val.ArrTime)
			} else if val.DepPlace == "LPA" && len(GRP) != 0 {
				GRP = append(GRP, val.DepTime)
			}
			break
		case "EC-GQF":
			if val.ArrPlace == "LPA" {
				GQF = append(GQF, val.ArrTime)
			} else if val.DepPlace == "LPA" && len(GQF) != 0 {
				GQF = append(GQF, val.DepTime)
			}
			break
		case "EC-LZR":
			if val.ArrPlace == "LPA" {
				LZR = append(LZR, val.ArrTime)
			} else if val.DepPlace == "LPA" && len(LZR) != 0 {
				LZR = append(LZR, val.DepTime)
			}
			break
		case "EC-LYZ":
			if val.ArrPlace == "LPA" {
				LYZ = append(LYZ, val.ArrTime)
			} else if val.DepPlace == "LPA" && len(LYZ) != 0 {
				LYZ = append(LYZ, val.DepTime)
			}
			break
		}
	}

	//order times
	sort.Strings(GRU)
	sort.Strings(GRP)
	sort.Strings(GQF)
	sort.Strings(LZR)
	sort.Strings(LYZ)
	//make every two times a"timeInLPA" object cheking that the arrival time is less than the departure time.
	GRUtimes := createTimesObject(GRU)
	GRPtimes := createTimesObject(GRP)
	GQFtimes := createTimesObject(GQF)
	LZRtimes := createTimesObject(LZR)
	LYZtimes := createTimesObject(LYZ)
	//create the final TimesInLPA objec
	var result []TimesInLPA
	var t TimesInLPA
	t.Aircraft = "EC-GRU"
	t.Times = GRUtimes
	result = append(result, t)
	t.Aircraft = "EC-GRP"
	t.Times = GRPtimes
	result = append(result, t)
	t.Aircraft = "EC-GQF"
	t.Times = GQFtimes
	result = append(result, t)
	t.Aircraft = "EC-LZR"
	t.Times = LZRtimes
	result = append(result, t)
	t.Aircraft = "EC-LYZ"
	t.Times = LYZtimes
	result = append(result, t)

	return result, nil
}

func (db *DB) DeleteParking(date time.Time) error {
	err := db.Ping()
	if err != nil {
		log.Println("DeleteParking error: db.Ping")
		return err
	}

	result, err := db.Exec("DELETE FROM parking WHERE date=?", date)
	if err != nil {
		log.Println("DeleteParking error: db.Exec")
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("DeleteParking error: RowsAffected")
		return err
	}
	log.Printf("Vuelo borrado (%d row cambiado)\n", rowsAffected)
	return nil
}
