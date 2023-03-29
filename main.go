package main

import (
	"time"

	meeting "github.com/tomyfalgui/meeting_meter/meeting"
)

func main() {
	participants := []int{5000, 3000, 120000}
	meter, _ := meeting.NewMeter(participants)
	meter.StartMeter()
	time.Sleep(time.Duration(1<<63 - 1))
}
