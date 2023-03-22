package main

import (
	"log"

	meeting "github.com/tomyfalgui/meeting_meter"
)

func main() {
	participants := []int{
		10000,
		3000,
		5000,
	}
	meter, err := meeting.NewMeter(participants)
	if err != nil {
		log.Fatal(err)
	}
	meter.StartMeter()
}
