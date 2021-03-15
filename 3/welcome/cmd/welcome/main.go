package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	flag.Parse()
	temp := `Добро пожаловать, {user}! Ваш код доступа: {code}.`
	name := flag.Arg(0)
	temp = strings.ReplaceAll(temp,"{user}",name)
	temp = strings.ReplaceAll(temp, "{code}",flag.Arg(1))
	fmt.Println(temp)

}
