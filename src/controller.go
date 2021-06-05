package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type RequestBody struct {
	Line     string      `json:"line"`
	Station  string      `json:"station"`
	Count    json.Number `json:"count"`
	TextMode bool        `json:"text_mode"`
}

func getArrivalTimes(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		return
	}
	reqBody := RequestBody{}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println("malformed body")
		res.WriteHeader(500)
		return
	}
	err = json.Unmarshal(body, &reqBody)
	if err != nil {
		log.Println("malformed body")
		res.WriteHeader(400)
		return
	}

	line := reqBody.Line
	station := reqBody.Station
	count, err := reqBody.Count.Int64()
	textMode := reqBody.TextMode

	stationName := getStationName(station)

	if stationName == "" {
		log.Println("invalid station name")
		res.WriteHeader(400)
		_, _ = res.Write([]byte("invalid station name"))
		return
	}

	log.Printf("Request: line=[%s], station=[%s], count=[%d]\n", line, getStationName(station), count)

	if line == "" || station == "" || err != nil {
		log.Println("required params not provided")
		res.WriteHeader(400)
		_, _ = res.Write([]byte("required params not provided"))
		return
	}

	arrivalTimes := getNextArrivalTimes(line, station, int(count), textMode)

	jsonResponse, err := json.Marshal(&arrivalTimes)
	if err != nil {
		log.Println("could not marshall json")
		res.WriteHeader(500)
		return
	}

	_, err = res.Write(jsonResponse)

	if err != nil {
		log.Println("could not write response")
		res.WriteHeader(500)
		return
	}
}
