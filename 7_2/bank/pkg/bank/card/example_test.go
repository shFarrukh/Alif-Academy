package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleAddBonus_positive() {
	cardf := types.Card{}
	cardf.Balance = 10_000_00
	cardf.Active = true
	cardf.MinBalance = 10_000_00
	AddBonus(&cardf, 3, 30, 365)
	fmt.Println(cardf.Balance)
	// Output
	// 1002465
}
func ExampleAddBonus_inactive() {
	cardf := types.Card{}
	cardf.Balance = 10_000_00
	cardf.Active = false
	cardf.MinBalance = 10_000_00
	AddBonus(&cardf, 3, 30, 365)
	fmt.Println(cardf.Balance)
	// Output
	// 1000000
}

func ExampleAddBonus_negativeBalance() {
	cardf := types.Card{}
	cardf.Balance = -10_000_00
	cardf.Active = true
	cardf.MinBalance = 10_000_00
	AddBonus(&cardf, 3, 30, 365)
	fmt.Println(cardf.Balance)
	// Output
	// -1000000
}

func ExampleAddBonus_limit() {
	cardf := types.Card{}
	cardf.Balance = 10_000_00
	cardf.Active = true
	cardf.MinBalance = 9000_000_00
	AddBonus(&cardf, 3, 30, 365)
	fmt.Println(cardf.Balance)
	// Output
	// 0000000
}
