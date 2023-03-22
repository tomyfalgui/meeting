package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

type meeting struct {
	Participants []int // hourly rate in cents
	ElaspedTime  time.Duration
	Output       io.Writer
}

type option func(*meeting) error

func NewMeter(participants []int, opts ...option) (meeting, error) {
	if len(participants) == 0 {
		return meeting{}, errors.New("Participants list can't be empty")
	}

	m := meeting{
		Output:       os.Stdout,
		Participants: participants,
	}

	for _, opt := range opts {
		err := opt(&m)
		if err != nil {
			return meeting{}, err
		}
	}
	return m, nil
}

func (m *meeting) StartMeter() {
	for range time.NewTicker(time.Second).C {
		m.ElaspedTime++
		m.Print()
	}
}

func (m meeting) CurrentCost() int {
	fractionalTime := 3600 / int(m.ElaspedTime)
	totalCost := 0
	for _, p := range m.Participants {
		totalCost += (p / fractionalTime)
	}
	return totalCost
}

func (m *meeting) Print() {
	fmt.Fprintf(m.Output, "Elapsed Seconds: %d\n", m.ElaspedTime)
	fmt.Fprintf(m.Output, "Total Cost: $%.2f\n", float64(m.CurrentCost())/100)
	fmt.Fprintf(m.Output, "\n")
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
