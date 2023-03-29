package meeting_test

import (
	"testing"

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
