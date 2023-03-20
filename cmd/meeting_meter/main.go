package main

import (
	"github.com/tomyfalgui/meeting"
)

func main() {
	meter := meeting.NewMeter()
	meter.AddParticipant(meeting.Participant{Name: "Joe", HourlyRate: 60, JoinTime: 0})
	meter.AddParticipant(meeting.Participant{Name: "Joe", HourlyRate: 60, JoinTime: 0})
	meter.AddParticipant(meeting.Participant{Name: "Joe", HourlyRate: 60, JoinTime: 0})
	meter.StartMeeting()
}
