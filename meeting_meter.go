package meeting_meter

type Participant struct {
	Name       string
	HourlyRate int
	JoinTime   int // exact second participant joined
}
type meeting struct {
	accruedCost  float64
	participants []Participant
}

func New() meeting {
	return meeting{
		accruedCost: 0,
	}
}

func (m *meeting) AddParticipant(p Participant) {
	m.participants = append(m.participants, p)
}

func (m meeting) Participants() []Participant {
	return m.participants
}

func (m meeting) CalculateMinuteCost() float64 {
	totalMinuteCost := 0.
	for _, p := range m.participants {
		minuteCost := p.GetSecondCost() * 60
		totalMinuteCost += minuteCost
	}

	return totalMinuteCost
}

func (p Participant) GetSecondCost() float64 {
	return float64(p.HourlyRate/60) / 60
}
