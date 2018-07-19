package main

import (
	"fmt"
	"./algorithm"
)

func main() {
	s := "abc"

	convert_s := algorithm.ZigZagConvert(s, 2)
	fmt.Println(convert_s);
}
