package coord

import (
	"log"
	"time"

	"github.com/misuher/RoutesMap/models"
)

//Position of one marker
type Position struct {
	Lat float32
	Lng float32
}

//Positions is a collections of markers position
type Positions struct {
	Pos []Position
}

func (p *Positions) AddElement(element Position) []Position {
	p.Pos = append(p.Pos, element)
	return p.Pos
}

func identifyOnGoingFlights(env models.Datastore) models.Flights {
	result := models.Flights{}
	date, _ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
	flights, err := env.GetDaily(date)
	if err != nil {
		//
	}
	for _, flight := range flights {
		if flight.DepTime < time.Now().Format("15:04") && flight.ArrTime > time.Now().Format("15:04") {
			//if flight.DepTime < "11:00" && flight.ArrTime > "11:00" {
			result.AddElement(*flight)
		}
	}
	log.Println(result)
	return result
}

func getFlightCoords(f models.Flight) Position {
	horas := getFlightTime(f.DepTime, f.ArrTime)
	log.Println(horas)
	origen := airportPos(f.DepPlace)
	log.Println(origen)
	destino := airportPos(f.ArrPlace)
	log.Println(destino)
	velx := getVelocity(origen.Lat, destino.Lat, horas)
	log.Println(velx)
	vely := getVelocity(origen.Lng, destino.Lng, horas)
	log.Println(vely)
	depTime, _ := time.Parse("15:04", f.DepTime)
	log.Println(depTime)
	long := getCoor(origen.Lng, depTime, vely)
	lat := getCoor(origen.Lat, depTime, velx)
	log.Println(lat)
	log.Println(long)
	return Position{Lat: lat, Lng: long}
}

func getFlightTime(departure string, arrival string) float64 {
	depTime, _ := time.Parse("15:04", departure)
	ArrTime, _ := time.Parse("15:04", arrival)
	duration := ArrTime.Sub(depTime)
	return duration.Seconds()
}

func getVelocity(origen float32, destino float32, duration float64) float32 {
	return (destino - origen) / float32(duration)
}

func getCoor(origin float32, DepTime time.Time, vel float32) float32 {
	tiempoSimulado, _ := time.Parse("15:04", time.Now().Format("15:04"))
	flyingTime := tiempoSimulado.Sub(DepTime)
	log.Printf("tiempo volando %f", flyingTime.Seconds())
	return origin + vel*float32(flyingTime.Seconds())
}
