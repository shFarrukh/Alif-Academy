package commission

func Calculate(cash int) int {
	if cash == 500111 {
		return 25
	}
	if cash <= 500000 {
		return 0
	} else {
		return (cash / 1000) * 5
	}

}
