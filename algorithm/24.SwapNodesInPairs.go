package algorithm

func SwapPairs(head *ListNode) *ListNode {

	newHead := &ListNode{}

	newHead.Next = head

	cur := head
	prev := newHead

	for cur != nil && cur.Next != nil {

		prev.Next = cur.Next
		temp := cur.Next.Next
		cur.Next.Next = cur
		cur.Next = temp
		prev = cur
		cur = temp
	}

	return newHead.Next
}