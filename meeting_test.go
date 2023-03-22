package meeting_test

import (
	"bytes"
	"testing"

	meeting "github.com/tomyfalgui/meeting_meter"
)

func TestNewMeterCreation(t *testing.T) {
	t.Parallel()

	participants := []int{
		10000,
	}
	_, err := meeting.NewMeter(participants)
	if err != nil {
		t.Fatalf("meeting creation can't fail")
	}
}

func TestNewMeterCreationWithEmptyParticpantList(t *testing.T) {
	t.Parallel()

	participants := []int{}
	_, err := meeting.NewMeter(participants)
	if err == nil {
		t.Fatalf("NewMeter() must throw an error with empty list")
	}
}

func TestMeterWithOutput(t *testing.T) {
	t.Parallel()

	output := &bytes.Buffer{}

	participants := []int{
		10000,
	}

	_, err := meeting.NewMeter(participants,
		meeting.WithOutput(output),
	)
	if err != nil {
		t.Fatalf("output is an invalid io.Writer")
	}
}

func TestGetCurrentCost(t *testing.T) {
	t.Parallel()

	participants := []int{
		10000,
		10000,
	}

	m, err := meeting.NewMeter(participants)
	if err != nil {
		t.Fatalf("meeting instantiation should not fail")
	}

	m.ElaspedTime = 50
	got := m.CurrentCost()
	want := 276

	if want != got {
		t.Fatalf("want %v != got %v", want, got)
	}
}
