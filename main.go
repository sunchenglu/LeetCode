package main

import (
	"fmt"
	"./algorithm"
)

func main() {
	x := []int{-4,-2,-2,-2,0,1,2,2,2,3,3,4,4,6,6}
	res := algorithm.ThreeSum(x)
	fmt.Println(res)
}
