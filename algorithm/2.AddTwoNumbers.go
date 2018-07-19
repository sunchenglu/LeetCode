package algorithm

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	add := 0
	var head *ListNode
	var tail *ListNode

	for l1 != nil || l2 != nil || add != 0 {

		num1 := 0

		if l1 != nil {
			num1 = l1.Val
			l1 = l1.Next
		} else {
			num1 = 0
		}

		num2 := 0

		if l2 != nil {
			num2 = l2.Val
			l2 = l2.Next
		} else {
			num2 = 0
		}

		sum := num2 + num1 + add

		if sum >= 10 {
			add = 1
		} else {
			add = 0
		}

		sum = sum % 10

		node := &ListNode{
			sum % 10,
			nil,
		}

		if nil == head {
			head = node
			tail = node
		} else {
			tail.Next = node
			tail = node
		}

	}

	return head

}
