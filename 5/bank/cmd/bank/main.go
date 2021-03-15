package main

import (
	"bank/pkg/bank/transfer"
	"fmt"
)

func main() {
	total := transfer.Total(500_000_00)
	fmt.Println(total)
}
