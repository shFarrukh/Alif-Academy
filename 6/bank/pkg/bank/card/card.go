package card

import "bank/pkg/bank/types"

func IssueCard(currancy types.Currency, color string, name string) types.Card {
	card := types.Card{
		ID:       1000,
		PAN:      "5058 xxxx xxxx 0001",
		Balance:  0,
		Currency: currancy,
		Color:    color,
		Name:     name,
		Active:   true,
	}
	return card
}
