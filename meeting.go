package meeting

import (
	"errors"
	"time"
)

type meeting struct {
	Participants []int
	ElapsedTime  time.Duration
}

func NewMeter(participants []int) (meeting, error) {
	if len(participants) == 0 {
		return meeting{}, errors.New("Participant list is empty")
	}
	return meeting{Participants: participants}, nil
}

func (m meeting) CurrentCost() int {
	totalCost := 0
	secondCost := 3600 / m.ElapsedTime
	for _, c := range m.Participants {
		totalCost += (c / int(secondCost))
	}
	return totalCost
}
