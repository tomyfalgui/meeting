package meeting

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

type meeting struct {
	Participants []int
	StartTime    time.Time
	Output       io.Writer
}

// NewMeter takes a list of integers (participants) and
// returns a initialized [meeting] struct. It returns an error if
// the participant list is empty.
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

// TotalCost returns the total cost of the ongoing meeting in cents
// which based on the elapsed time of the meeting.
func (m meeting) TotalCost() int {
	elapsedTime := m.ElapsedTime()
	if elapsedTime.Seconds() < 1 {
		return 0
	}

	fractionalTime := int(3600 / elapsedTime.Seconds())
	totalCost := 0
	for _, p := range m.Participants {
		totalCost += p / fractionalTime
	}

	return totalCost
}

// ElapsedTime returns the total duration of a meeting
// since it started.
func (m meeting) ElapsedTime() time.Duration {
	diff := time.Since(m.StartTime)

	return diff
}

// Main runs the command-line interface for meeting.
// The exit status for the binary is 1 if there are invalid options or arguments
func Main() int {
	fs := flag.NewFlagSet("meeting_meter", flag.ExitOnError)
	printInterval := fs.Duration("f", time.Second, "frequency of printing")

	fs.Usage = func() {
		fmt.Print("Usage: meeting_meter [OPTION] [VALUES]\n\n")
		fmt.Print("meeting_meter prints the total cost of a meeting\n\n")
		fmt.Print("Options:\n")
		fs.VisitAll(func(f *flag.Flag) {
			fmt.Printf("  -%s %s\n", f.Name, f.Usage)
		})
		fmt.Print("\n")
		fmt.Print("Arguments:\n")
		fmt.Print("  VALUES        Hourly Cost in Cents\n\n")
		fmt.Print("Examples:\n")
		fmt.Print("  Print every 5 seconds\n")
		fmt.Print("   meeting_meter -f 5s 10000\n")
	}
	fs.Parse(os.Args[1:])

	if fs.NArg() == 0 && fs.NFlag() == 0 {
		fs.Usage()
		return 1
	}

	if *printInterval < 0 {
		fmt.Println("Interval must be a positive integer.")
		fs.Usage()
		return 1

	}

	participants := []int{}
	for i := range fs.Args() {
		conv, err := strconv.Atoi(fs.Arg(i))
		if err != nil {
			fmt.Printf("invalid number: %v. Please provid a valid number\n", fs.Arg(i))
			return 1
		}
		participants = append(participants, conv)
	}

	meter, err := NewMeter(participants)
	if err != nil {
		fmt.Println(err)
		fs.Usage()
		return 1
	}

	for range time.NewTicker(*printInterval).C {
		elapsedTime := meter.ElapsedTime()
		totalCost := meter.TotalCost()

		fmt.Println(elapsedTime.Truncate(time.Second))
		fmt.Printf("Total Cost: $%.2f\n\n", float64(totalCost)/100)
	}

	return 0
}
