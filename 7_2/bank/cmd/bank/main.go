package main

import (
	"bank/pkg/bank/card"	
	"fmt"
)


func main() {
	cardf := card.IssueCard("TJS", "White", "Ftr")
	cardf.Balance = 10_000_00
	cardf.Active =true
	cardf.MinBalance = 9000_000_00
	card.AddBonus(&cardf,3,30,365)
	fmt.Println(cardf.Balance)
}