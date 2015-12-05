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
