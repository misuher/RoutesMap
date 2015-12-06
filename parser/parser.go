package parser

import (
	"bufio"
	"bytes"
	"log"
	"strings"
	"time"

	"github.com/misuher/RoutesMap/models"
)

//ParsePDF parse an expecific format file from a pdf
func ParsePDF(date time.Time, body []byte, env models.Datastore) error {
	buffer := bytes.NewReader(body)
	scanner := bufio.NewScanner(buffer)

	var err error
	var f models.Flight
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "EC") {
			f.Avion = scanner.Text()
			scanner.Scan()
			scanner.Scan()
			f.Vuelo = scanner.Text()
			scanner.Scan()
			scanner.Scan()
			f.DepPlace = scanner.Text()
			scanner.Scan()
			scanner.Scan()
			f.DepTime = scanner.Text()
			scanner.Scan()
			scanner.Scan()
			f.ArrPlace = scanner.Text()
			scanner.Scan()
			scanner.Scan()
			f.ArrTime = scanner.Text()
			scanner.Scan()
			scanner.Scan()
			f.Pasajeros = scanner.Text()
			scanner.Scan()
			scanner.Scan()
			f.Capitan = scanner.Text()
			scanner.Scan()
			scanner.Scan()
			f.PrimerOficial = scanner.Text()

			log.Printf("%v\n", f)

			err = env.SetDaily(date, f)
			if err != nil {
				log.Println("SetDaily error")
				return err
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	err = env.SetParking(date)
	if err != nil {
		log.Println("SetParking error")
		return err
	}
	return nil
}
