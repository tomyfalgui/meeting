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
	m, err := meeting.NewMeter(participants)
	if err != nil {
		t.Errorf("NewMeter threw an err")
	}
	if !cmp.Equal(participants, m.Participants) {
		t.Errorf("want %v, got %v", participants, m.Participants)
	}
}
