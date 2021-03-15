package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	flag.Parse()
	phoneNumber := flag.Arg(0)
	phoneNumber = strings.ReplaceAll(phoneNumber, "-", "")
	phoneNumber = strings.ReplaceAll(phoneNumber, " ", "")
	phoneNumber = strings.ReplaceAll(phoneNumber, "(", "")
	phoneNumber = strings.ReplaceAll(phoneNumber, ")", "")
	fmt.Println(phoneNumber)

}
