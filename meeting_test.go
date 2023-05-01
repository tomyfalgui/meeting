package meeting_test

import (
	"bytes"
	"math"
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/rogpeppe/go-internal/testscript"
	"github.com/tomyfalgui/meeting"
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

func TestGetElapsedTime(t *testing.T) {
	t.Parallel()

	participants := []int{10000}
	m, err := meeting.NewMeter(participants)
	if err != nil {
		t.Errorf("didnt expect NewMeter to fail")
	}
	fakeTerminal := &bytes.Buffer{}
	m.Output = fakeTerminal
	want := time.Second
	time.Sleep(time.Second)
	got := m.ElapsedTime()

	// consider time accuracy
	epsilon := 1.

	if math.Abs(want.Seconds()-got.Seconds()) > epsilon {
		t.Errorf("want %v != got %v", want, got)
	}
}

func TestGetTotalCostOfMeetingAfterTimeElapsed(t *testing.T) {
	t.Parallel()

	participants := []int{10000}
	m, err := meeting.NewMeter(participants)
	if err != nil {
		t.Errorf("didnt expect NewMeter to fail")
	}
	fakeTerminal := &bytes.Buffer{}
	m.Output = fakeTerminal
	want := 2
	time.Sleep(1 * time.Second)
	got := m.TotalCost()

	if want != got {
		t.Errorf("want %v != got %v", want, got)
	}
}

func TestMain(m *testing.M) {
	os.Exit(testscript.RunMain(m, map[string]func() int{
		"meeting_meter": meeting.Main,
	}))
}

func TestGetTotalCostWhenElapsedTimeIsZeroShouldReturnZero(t *testing.T) {
	t.Parallel()

	participants := []int{10000}
	m, err := meeting.NewMeter(participants)
	if err != nil {
		t.Errorf("didnt expect NewMeter to fail")
	}
	fakeTerminal := &bytes.Buffer{}
	m.Output = fakeTerminal
	want := 0
	time.Sleep(0 * time.Second)
	got := m.TotalCost()

	if want != got {
		t.Errorf("want %v != got %v", want, got)
	}

}

func Test(t *testing.T) {
	t.Parallel()

	testscript.Run(t, testscript.Params{
		Dir: "testdata",
	})
}
