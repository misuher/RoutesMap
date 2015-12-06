package models

import (
	"log"
	"sort"
	"time"
)

//timeInLPA defines the interval of time represented as 2 hours in which an aircraft is parked in LPA airport
type timeInLPA struct {
	arrival time.Time
	leave   time.Time
}

//timesInLPA is a collection of timeInLPA just in case that the aircraft stops several intervals of time in one day
type timesInLPA struct {
	aircraft string
	times    []timeInLPA
}

//CreateParking initialize a table in the db with Flights
func (db *DB) CreateParking() {
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS parking (
		id integer not null PRIMARY KEY,
		date DATE NOT NULL,
		aircraft VARCHAR(255) NOT NULL,
		arrive VARCHAR(255) NOT NULL,
		leave VARCHAR(255) NOT NULL,
		);
	`

	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}

func (db *DB) getParking(date time.Time) ([]*timesInLPA, error) {
	err := db.Ping()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT * FROM parking WHERE date = ? ORDER BY id DESC", date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var parkings []*timesInLPA
	for rows.Next() {
		parking := new(timesInLPA)
		err := rows.Scan(&parking.aircraft, &parking.times)
		if err != nil {
			return nil, err
		}
		log.Printf("Parking creado: %s %s  %s\n", date, parking.aircraft, parking.times)
		parkings = append(parkings, parking)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return parkings, nil
}

func (db *DB) setParking(date time.Time) error {
	//filter which flights arrive or leave LPA
	times, err := db.filterTimesInLPA(date)
	if err != nil {
		return err
	}

	//save result into db parking table
	for _, time := range times {
		result, err := db.Exec("INSERT INTO parking VALUES(?, ?, ?)", date, time.aircraft, time.times)
		if err != nil {
			return err
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return err
		}
		log.Printf("Parking %s creado (%d row cambiado)\n", time.aircraft, rowsAffected)
	}

	return nil
}

func (db *DB) filterTimesInLPA(date time.Time) ([]*timesInLPA, error) {
	//get all the daily flights
	flights, err := db.getDaily(date)
	if err != nil {
		return nil, err
	}

	//filter flights based on aircraft and place
	var GRU []time.Time
	var GRP []time.Time
	var GQF []time.Time
	var LZR []time.Time
	var LYZ []time.Time

	for _, val := range flights {
		switch val.Avion {
		case "EC-GRU":
			if val.ArrPlace == "LPA" {
				GRU = append(GRU, val.ArrTime)
			} else if val.DepPlace == "LPA" {
				GRU = append(GRU, val.DepTime)
			}
			break
		case "EC-GRP":
			if val.ArrPlace == "LPA" {
				GRP = append(GRP, val.ArrTime)
			} else if val.DepPlace == "LPA" {
				GRP = append(GRP, val.DepTime)
			}
			break
		case "EC-GQF":
			if val.ArrPlace == "LPA" {
				GQF = append(GQF, val.ArrTime)
			} else if val.DepPlace == "LPA" {
				GQF = append(GQF, val.DepTime)
			}
			break
		case "EC-LZR":
			if val.ArrPlace == "LPA" {
				LZR = append(LZR, val.ArrTime)
			} else if val.DepPlace == "LPA" {
				LZR = append(LZR, val.DepTime)
			}
			break
		case "EC-LYZ":
			if val.ArrPlace == "LPA" {
				LYZ = append(LYZ, val.ArrTime)
			} else if val.DepPlace == "LPA" {
				LYZ = append(LYZ, val.DepTime)
			}
			break
		}
	}

	//order times
	sort.Sort(sortTime(GRU))
	sort.Sort(sortTime(GRP))
	sort.Sort(sortTime(GQF))
	sort.Sort(sortTime(LZR))
	sort.Sort(sortTime(LYZ))
	//check there is even number of times
	err = isLenEven(GRU)
	if err != nil {
		return nil, err
	}
	err = isLenEven(GRP)
	if err != nil {
		return nil, err
	}
	err = isLenEven(GQF)
	if err != nil {
		return nil, err
	}
	err = isLenEven(LZR)
	if err != nil {
		return nil, err
	}
	err = isLenEven(LYZ)
	if err != nil {
		return nil, err
	}
	//make every two times a"timeInLPA" object cheking that the arrival time is less than the departure time.
	GRUtimes := createTimesObject(GRU)
	GRPtimes := createTimesObject(GRP)
	GQFtimes := createTimesObject(GQF)
	LZRtimes := createTimesObject(LZR)
	LYZtimes := createTimesObject(LYZ)
	//create the final timesInLPA objec
	var result []*timesInLPA
	result[0].aircraft = "EC-GRU"
	result[0].times = GRUtimes
	result[1].aircraft = "EC-GRP"
	result[1].times = GRPtimes
	result[2].aircraft = "EC-GQF"
	result[2].times = GQFtimes
	result[3].aircraft = "EC-LZR"
	result[3].times = LZRtimes
	result[4].aircraft = "EC-LYZ"
	result[4].times = LYZtimes

	return result, nil
}
