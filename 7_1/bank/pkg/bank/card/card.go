package card

import (
	"bank/pkg/bank/types"
)

func IssueCard(currency types.Currency, color string, name string) types.Card {
	card := types.Card{
		ID:       1000,
		PAN:      "5058 xxxx xxxx 0001",
		Balance:  0,
		Currency: currency,
		Color:    color,
		Name:     name,
		Active:   true,
	}
	return card
}
func Withdraw(card types.Card, amount types.Money) types.Card {
	// TODO:
	if amount < 0 || amount > 2000000 {
		return card
	}
	if card.Balance < amount {
		return card
	}
	// if card.Currency != "RUB"{
	// 	return card
	// }
	if !card.Active {
		return card
	}
	card.Balance = card.Balance - amount
	return card
}

func Deposit(card *types.Card, amount types.Money) {
	// TODO: произвести операøии с картой
	if card.Active == false || amount > 50_000_00 || amount < 0 {
		return
	}
	card.Balance += amount
}
