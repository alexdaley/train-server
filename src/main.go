package main

import (
	"log"
	"net/http"
)

func main() {
	csvResult := loadStationCsv()
	if csvResult == false {
		log.Fatal("Could not load station CSV")
	}

	http.HandleFunc("/api/arrivalTimes", CORS(getArrivalTimes))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Could not start server:", err)
	}
}
