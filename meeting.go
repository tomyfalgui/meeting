package meeting

import "errors"

type meeting struct{}

func NewMeter(participants []int) (meeting, error) {
	if len(participants) == 0 {
		return meeting{}, errors.New("Participant list is empty")
	}
	return meeting{}, nil
}
