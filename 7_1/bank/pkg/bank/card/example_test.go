package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleDeposit_positive() {
	result := types.Card{Balance: 10_000, Active: true}
	Deposit(&result, 10_000)
	fmt.Println(result.Balance)
	// Output:
	//20000
}

func ExampleDeposit_negative() {
	result := types.Card{Balance: -10_000, Active: true}
	Deposit(&result, 10_000)
	fmt.Println(result.Balance)
	// Output:
	// 0
}

func ExampleDeposit_inactive() {
	result := types.Card{Balance: 10_000, Active: false}
	Deposit(&result, 10_000)
	fmt.Println(result.Balance)
	// Output:
	// 10000
}

func ExampleDeposit_limit() {
	result := types.Card{Balance: 10_000, Active: true}
	Deposit(&result, 51_000_00)
	fmt.Println(result.Balance)
	// Output:
	// 10000
}
