package main

import (
	"./algorithm"
	"fmt"
)

func main() {

	nums := []int{1}

	res := algorithm.SearchRange(nums, 1)

	fmt.Println(res)

}
