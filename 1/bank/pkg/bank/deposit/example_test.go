package deposit

import "fmt"

func ExampleCalculate() {
	min := 0
	max := 0
	min, max = Calculate(0, "TJS")
	fmt.Println(min, max)
	min, max = Calculate(0, "USD")
	fmt.Println(min, max)
	min, max = Calculate(500_000_00, "TJS")
	fmt.Println(min, max)
	min, max = Calculate(500_000_00, "USD")
	fmt.Println(min, max)
	min, max = Calculate(1_000_000_00, "TJS")
	fmt.Println(min, max)
	min, max = Calculate(1_000_000_00, "USD")
	fmt.Println(min, max)

	// Output:
	// 0 0
	// 0 0
	// 166666 250000
	// 41666 83333
	// 333333 500000
	// 83333 166666

}
