package meeting_meter_test

import (
	"meeting_meter"
	"testing"
)

func TestCalculateMinuteCost(t *testing.T) {
	t.Parallel()

	// hourly rate
	participants := map[string]int{
		"John":  60,
		"Alice": 60,
		"Dave":  60,
		"Drew":  60,
	}

	want := 4
	got := meeting_meter.CalculateMinuteCost(participants)

	if want != got {
		t.Errorf("want %v != got %v", want, got)
	}
}
