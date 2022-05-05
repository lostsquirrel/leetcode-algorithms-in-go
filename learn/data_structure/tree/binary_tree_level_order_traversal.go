package tree

func levelOrder2(root *TreeNode) [][]int {
	r := make([][]int, 0)
	s := []*TreeNode{root}
	var l []int
	for len(s) > 0 {
		s, l = levelHelper(s)
		if len(l) > 0 {
			r = append(r, l)
		}

	}
	return r
}

func levelHelper(prev []*TreeNode) ([]*TreeNode, []int) {
	next := make([]*TreeNode, 0)
	l := make([]int, 0)
	if prev == nil {
		return next, l
	}
	for _, node := range prev {
		if node == nil {
			continue
		}
		l = append(l, node.Val)
		if node.Left != nil {
			next = append(next, node.Left)
		}
		if node.Right != nil {
			next = append(next, node.Right)
		}
	}
	return next, l
}
