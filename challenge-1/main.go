package main

import (
	"fmt"
	"logic/convert"
)

func main() {
	var timeFormat string = "24:60:60"
	fmt.Println("Enter the time format:")
	fmt.Scanf("%s", &timeFormat)

	fmt.Println(convert.TimeConv(timeFormat))
}
