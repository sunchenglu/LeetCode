package algorithm

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	var sortArr []int

	for l1 != nil {

		num1 := l1.Val

		if l2 == nil {
			sortArr = append(sortArr, num1)
			l1 = l1.Next
			continue
		}

		num2 := l2.Val

		if num1 <= num2{
			sortArr = append(sortArr, num1)
			l1 = l1.Next
			continue
		}

		for l2 != nil {
			num2 = l2.Val
			if num2 <= num1 {
				sortArr = append(sortArr, num2)
				l2 = l2.Next
			} else {
				sortArr = append(sortArr, num1)
				l1 = l1.Next
				break
			}
		}
	}

	for l2 != nil {
		num2 := l2.Val
		sortArr = append(sortArr, num2)
		l2 = l2.Next
	}

	length := len(sortArr)

	ret := &ListNode{
		Val:sortArr[length-1],
		Next:nil,
	}

	if length == 1 {
		return ret
	}

	for i:=length-2; i>=0; i-- {
		node := &ListNode{
			Val:sortArr[i],
			Next:ret,
		}
		ret = node
	}

	return ret
}
