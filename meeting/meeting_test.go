package meeting_test

import (
	"bytes"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	meeting "github.com/tomyfalgui/meeting_meter"
)

func TestNewMeterFailsWithEmptyParticipantList(t *testing.T) {
	t.Parallel()

	participants := []int{}
	_, err := meeting.NewMeter(participants)
	if err == nil {
		t.Errorf("expected NewMeter to fail")
	}
}

func TestParticipantListIsStored(t *testing.T) {
	t.Parallel()

	participants := []int{10000}
	m, _ := meeting.NewMeter(participants)
	if !cmp.Equal(participants, m.Participants) {
		t.Errorf("want %v, got %v", participants, m.Participants)
	}
}

func TestCostCalculationBasedOnElapsedTime(t *testing.T) {
	t.Parallel()

	participants := []int{10000, 10000}
	m, _ := meeting.NewMeter(participants)
	m.ElapsedTime = 5
	got := m.CurrentCost()
	want := 26

	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestStartMeterIncrementsElapsedTime(t *testing.T) {
	t.Parallel()

	participants := []int{10000, 10000}
	fakeTerminal := &bytes.Buffer{}
	m, _ := meeting.NewMeter(participants, meeting.WithOutput(fakeTerminal))
	ticker := m.StartMeter()

	time.Sleep(3500 * time.Millisecond)
	ticker.Stop()

	want := 3
	got := m.ElapsedTime

	if want != int(got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestWithOutputOption(t *testing.T) {
	t.Parallel()

	fakeTerminal := &bytes.Buffer{}
	participants := []int{10000, 10000}

	_, err := meeting.NewMeter(participants, meeting.WithOutput(fakeTerminal))
	if err != nil {
		t.Errorf("received err WithOutput")
	}
}

func TestPrintMeterWhenStarted(t *testing.T) {
	t.Parallel()

	fakeTerminal := &bytes.Buffer{}
	participants := []int{10000, 10000}

	m, _ := meeting.NewMeter(participants, meeting.WithOutput(fakeTerminal))
	ticker := m.StartMeter()
	time.Sleep(3500 * time.Millisecond)
	ticker.Stop()

	want := "Elapsed Seconds: 1\nTotal Cost: $0.04\n\nElapsed Seconds: 2\nTotal Cost: $0.10\n\nElapsed Seconds: 3\nTotal Cost: $0.16\n\n"
	got := fakeTerminal.String()

	if want != got {
		t.Errorf("%v", cmp.Diff(want, got))
	}
}
