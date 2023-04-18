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
	elapsedSeconds := m.ElapsedTime()
	if elapsedSeconds == 0 {
		return 0
	}

	fractionalTime := int(3600 / elapsedSeconds.Seconds())
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
	if len(os.Args) < 2 {
		fmt.Println(helpText())
		return 1
	}
	fset := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	printInterval := fset.Duration("d", time.Second, "frequency of printing")
	fset.Parse(os.Args[2:])

	// printIntervalConv, err := time.ParseDuration(*printInterval)
	// if err != nil {
	// 	fmt.Println("Please provide a valid time duration")
	// 	return 1
	// }

	conv, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Please provid a valid number")
		fmt.Println(helpText())
		return 1
	}

	participants := []int{conv}
	meter, _ := NewMeter(participants)

	for range time.NewTicker(*printInterval).C {
		elapsedTime := meter.ElapsedTime()
		totalCost := meter.TotalCost()

		fmt.Println(elapsedTime.Truncate(time.Second))
		fmt.Printf("Total Cost: $%.2f\n\n", float64(totalCost)/100)
	}

	return 0
}

func helpText() string {
	var helpText strings.Builder
	helpText.WriteString("meeting_meter - commandline meeting cost tracker\n\n")
	helpText.WriteString("usage: meeting_meter <hourly_cost_in_cents>\n\n")
	helpText.WriteString(
		"meeting_meter helps you track the total money wasted in a meeting\n\n",
	)
	helpText.WriteString("Example:\n")
	helpText.WriteString("\tmeeting_meter 30000")
	return helpText.String()
}
