package deposit

import "fmt"

func ExampleCalculate() {

	fmt.Println(Calculate(0, "TJS"))
	fmt.Println(Calculate(0, "USD"))
	fmt.Println(Calculate(500_000, "TJS"))
	fmt.Println(Calculate(500_000, "USD"))
	fmt.Println(Calculate(1_000_000, "TJS"))
	fmt.Println(Calculate(1_000_000, "USD"))

	// Output:
	// 0 0
	// 0 0
	// 166666 250000
	// 41666 83333
	// 333333 500000
	// 83333 166666

}
