package models

import (
	"errors"
	"time"
)

func isLenEven(t []time.Time) error {
	if len(t)%2 != 0 {
		return errors.New("odd number of flight times")
	}
	return nil
}

func createTimesObject(times []time.Time) []timeInLPA {
	i := 0
	var result []timeInLPA
	for index, item := range times {
		if index%2 == 0 {
			result[i].arrival = item
		} else {
			result[i].leave = item
			i++
		}
	}
	return result
}

type sortTime []time.Time

// Forward request for length
func (p sortTime) Len() int {
	return len(p)
}

// Define compare
func (p sortTime) Less(i, j int) bool {
	return p[i].Before(p[j])
}

// Define swap over an array
func (p sortTime) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
