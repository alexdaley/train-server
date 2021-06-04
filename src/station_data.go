package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

var stationMap = map[string]string{}

func getStationName(stationId string) string {
	return stationMap[stationId]
}

func loadStationCsv() bool {
	f, err := os.Open("src/resources/mtastops.txt")

	if err != nil {
		log.Println("Could not open csv")
		return false
	}

	r := csv.NewReader(f)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Println("Error reading CSV")
			return false
		}

		stationId := record[0]
		stationName := record[2]

		stationMap[stationId] = stationName
	}

	return true
}
