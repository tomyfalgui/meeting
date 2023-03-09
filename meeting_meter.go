package meeting_meter

type meeting struct {
	// In cents
	accruedCost  int
	participants map[string]int
}

func New(participants map[string]int) meeting {
	return meeting{
		accruedCost:  0,
		participants: participants,
	}
}

func (m meeting) CalculateMinuteCost() int {
	totalMinuteCost := 0
	for _, cost := range m.participants {
		minuteCost := cost / 60
		totalMinuteCost += minuteCost
	}

	return totalMinuteCost
}
