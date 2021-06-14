package main

import (
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"sort"
	"train-server/resources/com/google/transit/realtime"
)

const endpointUrl = "https://api-endpoint.mta.info/Dataservice/mtagtfsfeeds/nyct%2Fgtfs"

var lineEndpoints = map[string]string{
	"A": "-ace",
	"C": "-ace",
	"E": "-ace",
	"B": "-bdfm",
	"D": "-bdfm",
	"F": "-bdfm",
	"M": "-bdfm",
	"G": "-g",
	"J": "-jz",
	"Z": "-jz",
	"N": "-nqrw",
	"Q": "-nqrw",
	"R": "-nqrw",
	"W": "-nqrw",
	"L": "-l",
	"7": "-7",
	"1": "",
	"2": "",
	"3": "",
	"4": "",
	"5": "",
	"6": "",
}

func getStationArrivalTimes(trips []*realtime.TripUpdate, stationId string, count int) []int {
	arrivalTimes := make([]int, 0)
	for _, trip := range trips {
		for _, stopUpdate := range trip.StopTimeUpdate {
			if *stopUpdate.StopId == stationId && *stopUpdate.GetScheduleRelationship().Enum() == realtime.TripUpdate_StopTimeUpdate_SCHEDULED {
				arrivalTimes = append(arrivalTimes, int(stopUpdate.Departure.GetTime()))
			}
		}
	}

	sort.Ints(arrivalTimes)

	return arrivalTimes[0:int(math.Min(float64(count), float64(len(arrivalTimes))))]
}

func GetCurrentTrips(line string) []*realtime.TripUpdate {
	response := getApiResponse(line)
	feedMessage := parseFeedMessageFromApiResponse(response)
	trips := make([]*realtime.TripUpdate, 0)
	if feedMessage == nil {
		log.Println("No Data Retrieved")
		return trips
	}

	for _, entity := range feedMessage.Entity {
		if entity.GetTripUpdate() != nil && entity.GetTripUpdate().GetTrip().GetRouteId() == line {
			trips = append(trips, entity.GetTripUpdate())
		}
	}
	return trips
}

func parseFeedMessageFromApiResponse(response []byte) *realtime.FeedMessage {
	feedMessage := &realtime.FeedMessage{}
	err := proto.Unmarshal(response, feedMessage)
	if err != nil {
		log.Println("Could not parse feed message")
		return nil
	}

	return feedMessage
}

func getApiKey() string {
	return os.Getenv("MTA_API_KEY")
}

func getApiResponse(lineName string) []byte {
	endpointSuffix, valid := lineEndpoints[lineName]
	if valid == false {
		log.Println("Invalid Line Name given")
		return nil
	}
	client := &http.Client{}
	request, _ := http.NewRequest("GET", endpointUrl+endpointSuffix, nil)
	request.Header.Set("x-api-key", getApiKey())
	response, err := client.Do(request)

	if err != nil {
		log.Println("Could not get MTA Feed: ", err)
		return nil
	}

	if response.StatusCode != 200 {
		log.Println("Got unexpected response type ", response.StatusCode)
		return nil
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("Could not parse MTA Feed response: ", err)
		return nil
	}
	return data
}
