package billing

func Calculate1000(Megabyte int) int {
	price := 3500
	tarif := 1000

	if Megabyte <= tarif {
		return price
	}
	Megabyte = Megabyte - tarif
	return price + Megabyte*6
}
func Calculate2000(Megabyte int) int {
	price := 5500
	tarif := 2000
	if Megabyte <= tarif {
		return price
	}
	Megabyte -= tarif
	return price + Megabyte*6

}

func Calculate3000(Megabyte int) int {
	price := 7000
	tarif := 3000
	if Megabyte <= tarif {
		return price
	}
	Megabyte -= tarif
	return price + Megabyte*6

}
func Calculate5000(Megabyte int) int {

	price := 9500
	tarif := 5000
	if Megabyte <= tarif {
		return price
	}
	Megabyte -= tarif
	return price + Megabyte*6
}
func Calculate10000(Megabyte int) int {

	price := 17000
	tarif := 10000
	if Megabyte <= tarif {
		return price
	}
	Megabyte -= tarif
	return price + Megabyte*6
}
