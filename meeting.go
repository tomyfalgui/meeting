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
