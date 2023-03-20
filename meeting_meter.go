package meeting

import (
	"fmt"
	"time"
)

type meeting struct {
	accruedCost  int
	participants []int
	elapsedTime  int
}

func NewMeter() meeting {
	return meeting{
		accruedCost: 0,
	}
}

func (m *meeting) StartMeeting() {
	for range time.NewTicker(time.Second).C {
		m.elapsedTime++
		m.UpdateMeetingCost()
		m.DisplayParticipantCost()
		m.DisplayElapsedTime()
		m.DisplayMeetingCost()
	}
}

func (m meeting) DisplayParticipantCost() {
	fmt.Print("Participant Cost\n")
	for _, p := range m.participants {
	}
	fmt.Println()
}

func (m *meeting) UpdateMeetingCost() {
	m.accruedCost += (m.CalculateMinuteCost() / 60)
}

func (m meeting) DisplayMeetingCost() {
	fmt.Printf("Total Cost: $%d", m.accruedCost)
}

func (m meeting) DisplayElapsedTime() {
	fmt.Printf("Elapsed Time: %s\n", secondsToTime(m.elapsedTime))
}

func secondsToTime(seconds int) string {
	hours := seconds / 3600
	seconds = seconds - (hours * 3600)
	minutes := seconds / 60
	seconds = seconds - (minutes * 60)
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}
