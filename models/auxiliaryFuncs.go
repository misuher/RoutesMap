package models

import "errors"

func isLenEven(t []string) error {
	if len(t)%2 != 0 {
		return errors.New("odd number of flight times")
	}
	return nil
}

func createTimesObject(times []string) []TimeInLPA {
	var result []TimeInLPA
	var t TimeInLPA
	for index, s := range times {
		if index%2 == 0 {
			t.Arrival = s
		} else {
			t.Leave = s
			result = append(result, t)
		}
	}
	return result
}
