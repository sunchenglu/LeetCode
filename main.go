package main

import (
	"./algorithm"
	"fmt"
)

func main() {

	x := &algorithm.ListNode{
		Val:1,
		Next:nil,
	}

	res := algorithm.RemoveNthFromEnd(x, 1)
	fmt.Println(res)
}
