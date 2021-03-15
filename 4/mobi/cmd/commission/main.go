package main

import (
	"fmt"

	"mobi/pkg/commission"
)

func main() {
	result := commission.Calculate(200000)
	fmt.Println(result)

}
