package meeting

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

type meeting struct {
	Participants []int
	ElapsedTime  time.Duration
	Output       io.Writer
}

type option func(*meeting) error

func NewMeter(participants []int, opts ...option) (meeting, error) {
	if len(participants) == 0 {
		return meeting{}, errors.New("Participant list is empty")
	}
	m := meeting{Output: os.Stdout, Participants: participants}

	for _, opt := range opts {
		err := opt(&m)
		if err != nil {
			log.Fatal(err)
		}
	}
	return m, nil
}

func (m meeting) CurrentCost() int {
	totalCost := 0
	secondCost := 3600 / m.ElapsedTime
	for _, c := range m.Participants {
		totalCost += (c / int(secondCost))
	}
	return totalCost
}

func (m *meeting) StartMeter() *time.Ticker {
	ticker := time.NewTicker(time.Second)
	go func() {
		for range ticker.C {
			m.ElapsedTime++
			m.Print()
		}
	}()
	return ticker
}

func (m *meeting) Print() {
	fmt.Fprintf(m.Output, "Elapsed Seconds: %d\n", m.ElapsedTime)
	fmt.Fprintf(m.Output, "Total Cost: $%.2f\n\n", float64(m.CurrentCost())/100)
}

func WithOutput(output io.Writer) option {
	return func(m *meeting) error {
		m.Output = output
		return nil
	}
}
