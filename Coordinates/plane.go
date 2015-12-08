package coord

import "github.com/misuher/RoutesMap/models"

type Plane struct {
	LastAirport Position
	Pos         Position
}

var (
	LPA     = Position{27.926075, -15.390818}
	TFN     = Position{28.040288, -16.572979}
	ACE     = Position{28.946344, -13.607218}
	EUN     = Position{27.142294, -13.225541}
	SPC     = Position{28.622109, -17.755491}
	DAC     = Position{23.717010, -15.933434}
	DEFAULT = Position{28.946344, -16.572979}
)

//CalculateCoords is give lat and lng in the line between LPA and the others based in the real hour compared to the daily flights.
func (p *Plane) CalculateCoords(env models.Datastore, aircraft string) Position {
	vuelos := identifyOnGoingFlights(env)
	for _, vuelo := range vuelos.F {
		if vuelo.Avion == aircraft {
			p.updateAirport(vuelo.ArrPlace)
			return getFlightCoords(vuelo)
		}
	}
	return p.LastAirport
}

func (p *Plane) updateAirport(newAirport string) {
	airportCoor := airportPos(newAirport)
	p.LastAirport = airportCoor
}

func airportPos(airport string) Position {
	result := Position{}
	switch airport {
	case "LPA":
		result = LPA
		break
	case "TFN":
		result = TFN
		break
	case "ACE":
		result = ACE
		break
	case "EUN":
		result = EUN
		break
	case "SPC":
		result = SPC
		break
	case "DAC":
		result = DAC
		break
	}
	return result
}
