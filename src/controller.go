package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type RequestBody struct {
	Line    string      `json:"line"`
	Station string      `json:"station"`
	Count   json.Number `json:"count"`
}

func getArrivalTimes(res http.ResponseWriter, req *http.Request) {
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

	if line == "" || station == "" || err != nil {
		log.Println("required params not provided")
		res.WriteHeader(400)
		_, _ = res.Write([]byte("required params not provided"))
		return
	}

	arrivalTimes := getNextArrivalTimes(line, station, int(count))

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
