package main

import (
	"fmt"
	"time"
)

func getTimeStringFromEpoch(epoch int) string {
	return time.Unix(int64(epoch), 0).Format(time.RFC822Z)
}

func getClockTime(epoch int) string {
	return time.Unix(int64(epoch), 0).Format(time.Kitchen)
}

func getMinsFromNowEpoch(epoch int) string {
	diff := epoch - int(time.Now().Unix())
	mins := diff / 60
	return fmt.Sprintf("%d mins", mins)
}

func getTimeString(epoch int) string {
	return fmt.Sprintf("%s (%s)",
		getMinsFromNowEpoch(epoch),
		getClockTime(epoch))
}
