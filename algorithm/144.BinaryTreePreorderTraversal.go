package algorithm

func PreorderTraversal(root *TreeNode) []int {

	ret := make([]int, 0)
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

	return ret
}