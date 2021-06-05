package main

import (
	"math"
	"strconv"
)

func getNextArrivalTimes(line string, station string, count int, textMode bool) []string {
	count = int(math.Min(float64(count), 10))
	arrivalStrings := make([]string, 0)
	trips := GetCurrentTrips(line)
	arrivals := getStationArrivalTimes(trips, station, count)
	for _, arrival := range arrivals {
		if textMode {
			arrivalStrings = append(arrivalStrings, getTimeString(arrival))
		} else {
			arrivalStrings = append(arrivalStrings, strconv.Itoa(arrival))
		}
	}
	return arrivalStrings
}
