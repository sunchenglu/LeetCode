package main

import (
	"fmt"
	"./algorithm"
)

func main() {
	x := "MCMXCIV"
	res := algorithm.RomanToInt(x)
	fmt.Println(res);
}
