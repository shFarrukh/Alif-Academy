package main

func main() {
	amount := 999_999_00
	minPercent := 4
	maxPercent := 6
	minIncome := (amount/100) *minPercent
	maxIncome := (amount/100) *maxPercent
	println(minIncome)
	println(maxIncome)

	
}