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
func ParsePDF(body []byte, env models.Datastore) error {
	buffer := bytes.NewReader(body)
	scanner := bufio.NewScanner(buffer)

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
			f.DepTime, _ = time.Parse(time.Kitchen, scanner.Text())
			scanner.Scan()
			scanner.Scan()
			f.ArrPlace = scanner.Text()
			scanner.Scan()
			scanner.Scan()
			f.ArrTime, _ = time.Parse(time.Kitchen, scanner.Text())
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

			date, _ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
			err := env.SetDaily(date, f)

			if err != nil {
				log.Println("SetDaily error")
				return err
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
