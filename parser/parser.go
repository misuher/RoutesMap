package parser

import (
	"bufio"
	"bytes"
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
			f.DepTime, _ = time.Parse("06:10", scanner.Text())
			scanner.Scan()
			scanner.Scan()
			f.ArrPlace = scanner.Text()
			scanner.Scan()
			scanner.Scan()
			f.ArrTime, _ = time.Parse("06:10", scanner.Text())
			scanner.Scan()
			scanner.Scan()
			f.Pasajeros = scanner.Text()
			scanner.Scan()
			scanner.Scan()
			f.Capitan = scanner.Text()
			scanner.Scan()
			scanner.Scan()
			f.PrimerOficial = scanner.Text()

			err := env.SetDaily(time.Now(), f)
			if err != nil {
				return err
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
