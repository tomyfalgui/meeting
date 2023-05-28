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

// Meeting struct holds the related data for
// tracking the total cost and duration of a meeting.
type Meeting struct {
	Participants []int
	StartTime    time.Time
	Output       io.Writer
}

// NewMeter takes a list of integers (participants) and
// returns an initialized [Meeting] struct. It returns an error if
// the participant list is empty.
func NewMeter(participants []int) (Meeting, error) {
	if len(participants) == 0 {
		return Meeting{}, errors.New("participant list is empty")
	}
	m := Meeting{
		Output:       os.Stdout,
		Participants: participants,
		StartTime:    time.Now(),
	}

	return m, nil
}

// TotalCost returns the total cost of the ongoing meeting in cents
// which based on the elapsed time of the meeting.
func (m Meeting) TotalCost() int {
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
func (m Meeting) ElapsedTime() time.Duration {
	diff := time.Since(m.StartTime)

	return diff
}

// Main runs the command-line interface for meeting.
// The exit status for the binary is 1 if there are invalid options or arguments
func Main() int {
	fs := flag.NewFlagSet("meeting_meter", flag.ExitOnError)
	printInterval := fs.Duration("f", time.Second, "frequency of printing")

	fs.Usage = func() {
		fmt.Fprint(os.Stderr, "Usage: meeting_meter [OPTION] [VALUES]\n\n")
		fmt.Fprint(os.Stderr, "meeting_meter prints the total cost of a meeting\n\n")
		fmt.Fprint(os.Stderr, "Options:\n")
		fs.VisitAll(func(f *flag.Flag) {
			fmt.Fprintf(os.Stderr, "  -%s %s\n", f.Name, f.Usage)
		})
		fmt.Fprint(os.Stderr, "\n")
		fmt.Fprint(os.Stderr, "Arguments:\n")
		fmt.Fprint(os.Stderr, "  VALUES        Hourly Cost in Cents\n\n")
		fmt.Fprint(os.Stderr, "Examples:\n")
		fmt.Fprint(os.Stderr, "  Print every 5 seconds\n")
		fmt.Fprint(os.Stderr, "   meeting_meter -f 5s 10000\n")
	}
	err := fs.Parse(os.Args[1:])
	if err != nil {
		return 0
	}

	if fs.NArg() == 0 && fs.NFlag() == 0 {
		fs.Usage()
		return 1
	}

	if *printInterval < 0 {
		fmt.Fprintln(os.Stderr, "Interval must be a positive integer.")
		fs.Usage()
		return 1

	}

	var participants []int
	for i := range fs.Args() {
		conv, err := strconv.Atoi(fs.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stderr, "invalid number: %v. Please provid a valid number\n", fs.Arg(i))
			return 1
		}
		participants = append(participants, conv)
	}

	meter, err := NewMeter(participants)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
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
