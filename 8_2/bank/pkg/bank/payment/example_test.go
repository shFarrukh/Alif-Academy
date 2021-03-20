package payment
import (
	"fmt"
	"bank/pkg/bank/types"

)
func ExampleMax(){
	slice  := [] types.Payment{
		{
			ID : 1,
			Amount : 10_000_00,
		},
		{
			ID : 2,
			Amount : 11_000_00,
		},
		{
			ID : 3,
			Amount : 12_000_00,
		},
	} 
	max := Max(slice)
	fmt.Println(max.ID)
	// Output
	// 3
}