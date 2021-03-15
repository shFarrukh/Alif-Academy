package card

import (
	"fmt"	
	"bank/pkg/bank/types"
)


func ExampleWithdraw_positive() {
	result := Withdraw(types.Card{Balance: 20_000_00, Active: true}, 10_000_00)
	fmt.Println(result.Balance)
	// Output: 1000000
}

func ExampleWithdraw_noMoney() {
	result := Withdraw(types.Card{Balance: 20_000, Active: true}, 10_000_00)
	fmt.Println(result.Balance)
	// Output: 20000
}
func ExampleWithdraw_inactive() {
	result := Withdraw(types.Card{Balance: 20_000_00, Active: false}, 10_000_00)
	fmt.Println(result.Balance)
	// Output: 2000000
}
func ExampleWithdraw_limit() {
	result := Withdraw(types.Card{Balance: 20_000_00, Active: true}, 30_000_00)
	fmt.Println(result.Balance)
	// Output: 2000000

}