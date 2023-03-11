package main

import (
	mm "meeting_meter"
)

func main() {
	meter := mm.New()
	meter.AddParticipant(mm.Participant{Name: "Joe", HourlyRate: 60, JoinTime: 0})
	meter.AddParticipant(mm.Participant{Name: "Joe", HourlyRate: 60, JoinTime: 0})
	meter.AddParticipant(mm.Participant{Name: "Joe", HourlyRate: 60, JoinTime: 0})
	meter.StartMeeting()
}
