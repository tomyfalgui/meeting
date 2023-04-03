package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/tomyfalgui/meeting_meter/meeting"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(helpText())
		return
	}

	totalCost := os.Args[1]
	conv, err := strconv.Atoi(totalCost)
	if err != nil {
		fmt.Println("Please provid a valid number")
		fmt.Println(helpText())
		return
	}

	participants := []int{conv}
	meter, _ := meeting.NewMeter(participants)

	for range time.NewTicker(time.Second).C {
		elapsedTime := meter.ElapsedTime()
		totalCost := meter.TotalCost()

		fmt.Printf(timeFormat(elapsedTime))
		fmt.Printf("Total Cost: $%.2f\n\n", float64(totalCost)/100)
	}
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

func timeFormat(seconds int) string {
	var sb strings.Builder
	hours := seconds / 3600

	if hours >= 1 {
		fmt.Fprintf(&sb, "Hours: %d ", hours)
	}
	seconds = seconds - (hours * 3600)
	minutes := seconds / 60
	if minutes >= 1 {
		fmt.Fprintf(&sb, "Minutes: %d ", minutes)
	}
	seconds = seconds - (minutes * 60)
	fmt.Fprintf(&sb, "Seconds: %d \n", seconds)
	return sb.String()
}
