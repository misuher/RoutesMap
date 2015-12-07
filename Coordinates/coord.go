package coord

//Position of one marker
type Position struct {
	Lat float32
	Lng float32
}

//Positions is a collections of markers position
type Positions struct {
	Pos []Position
}

var (
	LPA = Position{27.926075, -15.390818}
	TFN = Position{28.040288, -16.572979}
	ACE = Position{28.946344, -13.607218}
	EUN = Position{27.142294, -13.225541}
	SPC = Position{28.622109, -17.755491}
	DAC = Position{23.717010, -15.933434}
)

//CalculateCoords is give lat and lng in the line between LPA and the others based in the real hour compared to the daily flights.
func CalculateCoords(Positions, error) {

}

func identifyOnGoingFlights() {

}

func getLong() float32 {
	return 3.3
}

func percent(departure string, arrival string, now string) {

}

func getLat(origin Position, destination Position, long float32) float32 {
	slope := slope(origin, destination)
	lat := origin.Lat + slope*(long-origin.Lng)
	return lat
}

func slope(origin Position, destination Position) float32 {
	return (destination.Lat - origin.Lat) / (destination.Lng - origin.Lat)
}
