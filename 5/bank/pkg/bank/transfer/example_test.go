package transfer

import "fmt"

func ExampleTotal() {
	result := Total(0)
	fmt.Println(result)

	result = Total(500000)
	fmt.Println(result)

	result = Total(1000000)
	fmt.Println(result)
	// Output:
	// 0
	// 502500
	// 1005000

}
