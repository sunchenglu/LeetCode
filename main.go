package main

import (
	"fmt"
	"./algorithm"
)

func main() {
	x := []int{-3,-1,0,2,4,5}
	res := algorithm.FourSum(x, 0)
	fmt.Println(res)
}
