package deposit

// Calculate calculates the deposit params.cd..
func Calculate(amount int, currency string) (min int, max int) {
	if amount == 0 {
		return 0, 0
	}
	var minPercent int
	var maxPercent int
	minPercent, maxPercent = PrecentFor(currency)

	min = ((amount * minPercent / 100) / 12)
	max = ((amount * maxPercent / 100) / 12)
	return min, max
}

// Retrun Precents
func PrecentFor(currency string) (min int, max int) {
	switch currency {
	case "TJS":
		return 4, 6

	case "USD":
		return 1, 2
	default:
		return 0, 0
	}
}
