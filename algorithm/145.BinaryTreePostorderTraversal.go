package algorithm

func PostorderTraversal(root *TreeNode) []int {

	ret := make([]int, 0)
	stack := make([]*TreeNode, 0)

	p := root

	for p != nil || len(stack) != 0 {

		if p != nil {
			stack = append(stack, p)
			ret = append(ret, p.Val)
			p = p.Right
		} else {
			p = stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			p = p.Left
		}
	}

	newRet := make([]int, len(ret))
	for i := 0; i < len(ret); i++ {
		newRet[i] = ret[len(ret)-1-i]
	}

	return newRet
}