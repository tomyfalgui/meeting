package meeting_test

import (
	"testing"

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
