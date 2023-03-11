package meeting_meter

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

type Participant struct {
	Name       string
	HourlyRate int
	JoinTime   int // exact second participant joined
}
type meeting struct {
	accruedCost  float64
	participants []Participant
	elapsedTime  int
}

func New() meeting {
	return meeting{
		accruedCost: 0,
	}
}

func (m *meeting) AddParticipant(p Participant) {
	m.participants = append(m.participants, p)
}

func (m meeting) Participants() []Participant {
	return m.participants
}

func (m *meeting) StartMeeting() {
	ticker := time.NewTicker(1 * time.Second)
	for range ticker.C {
		m.elapsedTime += 1
		clearScreen()
		m.UpdateMeetingCost()
		m.DisplayParticipantCost()
		m.DisplayElapsedTime()
		m.DisplayMeetingCost()
	}
}

func (m meeting) DisplayParticipantCost() {
	fmt.Print("Participant Cost\n")
	for _, p := range m.participants {
		fmt.Printf("%s - $%.2f\n", p.Name, p.GetSecondCost()*float64(m.elapsedTime))
	}
	fmt.Println()
}

func (m *meeting) UpdateMeetingCost() {
	m.accruedCost += (m.CalculateMinuteCost() / 60)
}

func (m meeting) DisplayMeetingCost() {
	fmt.Printf("Total Cost: $%.2f", m.accruedCost)
}

func (m meeting) DisplayElapsedTime() {
	fmt.Printf("Elapsed Time: %s\n", secondsToTime(m.elapsedTime))
}

func (m meeting) CalculateMinuteCost() float64 {
	totalMinuteCost := 0.
	for _, p := range m.participants {
		minuteCost := p.GetSecondCost() * 60
		totalMinuteCost += minuteCost
	}

	return totalMinuteCost
}

func (p Participant) GetSecondCost() float64 {
	return (float64(p.HourlyRate/60) / 60)
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func secondsToTime(seconds int) string {
	hours := seconds / 3600
	seconds = seconds - (hours * 3600)
	minutes := seconds / 60
	seconds = seconds - (minutes * 60)
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}
