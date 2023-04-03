package meeting_test

import (
	"bytes"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/tomyfalgui/meeting_meter/meeting"
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

func TestGetTotalCost(t *testing.T) {
	t.Parallel()

	participants := []int{10000}
	m, err := meeting.NewMeter(participants)
	if err != nil {
		t.Errorf("didnt expect NewMeter to fail")
	}
	fakeTerminal := &bytes.Buffer{}
	m.Output = fakeTerminal
	want := 13
	time.Sleep(5 * time.Second)
	got := m.TotalCost()

	if want != got {
		t.Errorf("want %v != got %v", want, got)
	}
}
