package main

import (
	"log"
)

func main() {
	participants := []int{
		10000,
		3000,
		5000,
	}
	meter, err := NewMeter(participants)
	if err != nil {
		log.Fatal(err)
	}
	meter.StartMeter()
}
