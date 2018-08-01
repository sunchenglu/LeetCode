package algorithm

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {

	if head == nil {
		return head
	}

	slow := head
	fast := head
	i := 0 

	for i=0; i<n; i++ {
		if fast == nil {
			break
		}
		fast = fast.Next
	}

	if i < n {
		return head
	}

	if fast == nil {
		return head.Next
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return head
}