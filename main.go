package main

import (
	"fmt"
	"./algorithm"
)

func main() {
	x := []int{1,2,3,4,5}
	res := algorithm.MaxArea(x)
	fmt.Println(res);
}
