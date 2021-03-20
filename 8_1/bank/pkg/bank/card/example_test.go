package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleTotal() {
	cardfs := []types.Card{
		{
			Balance: 10_000_00,
			Active:  true,
		},
		{
			Balance: 10_000_00,
			Active:  true,
		},
		{
			Balance: 10_000_00,
			Active:  true,
		},
	}
	fmt.Println(cardfs[0])
	// Output
	// 3000000
}
