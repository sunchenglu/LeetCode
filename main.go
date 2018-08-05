package main

import (
	"./algorithm"
	"fmt"
)

func main() {

	n := 3

	res := algorithm.GenerateParenthesis(n)

	for _, v := range res {
		fmt.Println(v)
	}


}
