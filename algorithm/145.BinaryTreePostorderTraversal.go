package algorithm

func PostorderTraversal(root *TreeNode) []int {

	ret := make([]int, 0)

	if root == nil {
		return ret
	}

	stack := []*TreeNode{root}

	for len(stack) != 0 {
		p := stack[len(stack)-1]
		stack = stack[0: len(stack) -1]
		ret = append(ret, p.Val)

		if p.Left != nil {
			stack = append(stack, p.Left)
		}

		if p.Right != nil {
			stack = append(stack, p.Right)
		}
	}

	newRet := make([]int, len(ret))
	for i:=0; i<len(ret); i++ {
		newRet[i] = ret[len(ret)-1-i]
	}

	return newRet
}