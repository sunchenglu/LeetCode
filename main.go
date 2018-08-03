package main

import (
	"./algorithm"
	"fmt"
)

func main() {

	l1 := &algorithm.ListNode{
		Val:4,
		Next:nil,
	}

	l2 := &algorithm.ListNode{
		Val:4,
		Next:nil,
	}

	for _, v := range []int{2, 1} {
		node := &algorithm.ListNode{
			v,
			nil,
		}
		node.Next = l1
		l1 = node
	}

	for _, v := range []int{3, 1} {
		node := &algorithm.ListNode{
			v,
			nil,
		}
		node.Next = l2
		l2 = node
	}

	res := algorithm.MergeTwoLists(l1, l2)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}
