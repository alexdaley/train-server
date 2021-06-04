package main

import "math"

func getNextArrivalTimes(line string, station string, count int) []string {
	count = int(math.Min(float64(count), 10))
	arrivalStrings := make([]string, 0)
	trips := GetCurrentTrips(line)
	arrivals := getStationArrivalTimes(trips, station, count)
	for _, arrival := range arrivals {
		arrivalStrings = append(arrivalStrings, getTimeString(arrival))
	}
	return arrivalStrings
}
