package meeting

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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

func (m meeting) ElapsedTime() time.Duration {
	diff := time.Since(m.StartTime)

	return diff
}

func Main() int {
	fs := flag.NewFlagSet("meeting_meter", flag.ExitOnError)
	printInterval := fs.Duration("f", time.Second, "frequency of printing")

	var sb strings.Builder
	fs.Usage = func() {
		sb.WriteString("Usage: meeting_meter [OPTION] [VALUES]\n\n")
		sb.WriteString("meeting_meter prints the total cost of a meeting\n\n")
		sb.WriteString("Options:\n")
		fs.VisitAll(func(f *flag.Flag) {
			fmt.Fprintf(&sb, "  -%s %s\n", f.Name, f.Usage)
		})
		sb.WriteString("\n")
		sb.WriteString("Arguments:\n")
		sb.WriteString("  VALUES        Hourly Cost in Cents\n\n")
		sb.WriteString("Examples:\n")
		sb.WriteString("  Print every 5 seconds\n")
		sb.WriteString("   meeting_meter -f 5s 10000\n")
		fmt.Print(sb.String())
	}
	fs.Parse(os.Args[1:])

	if fs.NArg() == 0 {
		fs.Usage()
		return 1
	}

	participants := []int{}
	for i := 0; i < fs.NArg(); i++ {

		conv, err := strconv.Atoi(fs.Arg(i))
		if err != nil {
			fmt.Printf("invalid number: %v. Please provid a valid number\n", fs.Arg(i))
			return 1
		}
		participants = append(participants, conv)
	}

	meter, _ := NewMeter(participants)

	for range time.NewTicker(*printInterval).C {
		elapsedTime := meter.ElapsedTime()
		totalCost := meter.TotalCost()

		fmt.Println(elapsedTime.Truncate(time.Second))
		fmt.Printf("Total Cost: $%.2f\n\n", float64(totalCost)/100)
	}

	return 0
}
