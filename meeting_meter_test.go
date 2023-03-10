package meeting_meter_test

import (
	"meeting_meter"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAddParticipant(t *testing.T) {
	t.Parallel()

	meter := meeting_meter.New()
	meter.AddParticipant(meeting_meter.Participant{
		Name:       "John",
		HourlyRate: 60,
		JoinTime:   0,
	})
	meter.AddParticipant(meeting_meter.Participant{
		Name:       "Alice",
		HourlyRate: 60,
		JoinTime:   0,
	})

	want := []meeting_meter.Participant{
		{
			Name:       "John",
			HourlyRate: 60,
			JoinTime:   0,
		},
		{
			Name:       "Alice",
			HourlyRate: 60,
			JoinTime:   0,
		},
	}
	got := meter.Participants()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("Participants() mismatch (-want +got):\n%s", diff)
	}
}

func TestCalculateMinuteCost(t *testing.T) {
	t.Parallel()

	meter := meeting_meter.New()
	meter.AddParticipant(meeting_meter.Participant{
		Name:       "John",
		HourlyRate: 60,
		JoinTime:   0,
	})
	meter.AddParticipant(meeting_meter.Participant{
		Name:       "Alice",
		HourlyRate: 60,
		JoinTime:   0,
	})

	want := 2.
	got := meter.CalculateMinuteCost()

	if want != got {
		t.Errorf("want %v != got %v", want, got)
	}
}
