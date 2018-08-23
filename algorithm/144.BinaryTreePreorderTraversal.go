package algorithm

func PreorderTraversal(root *TreeNode) []int {

	ret := make([]int, 0)
	stack := make([]*TreeNode, 0)

	p := root

	for p != nil || len(stack) != 0 {

		if p != nil {
			stack = append(stack, p)
			p = p.Left
		} else {
			p = stack[len(stack)-1]
			stack = stack[0:len(stack)-1]
			ret = append(ret, p.Val)
			p = p.Right
		}
	}
	return ret
}