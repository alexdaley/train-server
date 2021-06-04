package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/arrivalTimes", getArrivalTimes)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Could not start server:", err)
	}
}
