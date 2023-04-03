package meeting

import (
	"errors"
	"io"
	"os"
	"time"
)

type meeting struct {
	Participants []int
	StartTime    time.Time
	Output       io.Writer
}

func NewMeter(participants []int) (meeting, error) {
	if len(participants) == 0 {
		return meeting{}, errors.New("Participant list is empty")
	}
	m := meeting{
		Output:       os.Stdout,
		Participants: participants,
		StartTime:    time.Now(),
	}

	return m, nil
}

func (m meeting) TotalCost() int {
	elapsedSeconds := m.ElapsedTime()
	if elapsedSeconds == 0 {
		return 0
	}

	fractionalTime := 3600 / elapsedSeconds
	totalCost := 0
	for _, p := range m.Participants {
		totalCost += p / fractionalTime
	}

	return totalCost
}

func (m meeting) ElapsedTime() int {
	now := time.Now()
	diff := now.Sub(m.StartTime)
	elapsedSeconds := int(diff.Seconds())

	return elapsedSeconds
}
