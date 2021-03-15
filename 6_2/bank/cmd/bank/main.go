package main

import (
	"bank/pkg/bank/card"
	"fmt"
)

func main() {
	card := card.IssueCard("TJS", "white", "Infinity")
	fmt.Println(card)
}
