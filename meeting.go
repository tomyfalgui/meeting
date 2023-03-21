package meeting

import (
	"errors"
	"io"
	"os"
	"time"
)

type meeting struct {
	Participants []int // hourly rate in cents
	AccruedCost  int
	ElaspedTime  int
	Output       io.Writer
}

type option func(*meeting) error

func NewMeter(opts ...option) (meeting, error) {
	m := meeting{
		Output: os.Stdout,
	}

	for _, opt := range opts {
		err := opt(&m)
		if err != nil {
			return meeting{}, err
		}
	}
	return m, nil
}

func (m meeting) StartMeter() {
	for range time.NewTicker(time.Second).C {
	}
}

func (m meeting) GetSecondCost() int {
	return 0
}

func Print() {
}

func WithOutput(output io.Writer) option {
	return func(m *meeting) error {
		if output == nil {
			return errors.New("nil output writer")
		}
		m.Output = output
		return nil
	}
}

func WithParticipants(participants []int) option {
	return func(m *meeting) error {
		if participants == nil {
			return errors.New("nil participant list")
		}
		m.Participants = participants
		return nil
	}
}
