package deposit

//Calculate расчитывает параметры вклада
func Calculate(amount int, currency string) (min int, max int) {
	var minPercent int
	var maxPercent int
	switch currency {
	case "TJS":
		minPercent = 4
		maxPercent = 6

	case "USD":
		minPercent = 1
		maxPercent = 2
	}

	min = (amount * minPercent / 100) / 12
	max = (amount * maxPercent / 100) / 12
	return
}
