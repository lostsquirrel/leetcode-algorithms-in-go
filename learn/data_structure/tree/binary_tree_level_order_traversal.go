package tree

func levelOrder2(root *TreeNode) [][]int {
	r := make([][]int, 0)
	for s := []*TreeNode{root}; len(s) > 0; s = levelHelper(s) {
		l := make([]int, 0)
		for _, n := range s {
			if n != nil {
				l = append(l, n.Val)
			}
		}
		if len(l) > 0 {
			r = append(r, l)
		}

	}
	return r
}

func levelHelper(prev []*TreeNode) []*TreeNode {
	next := make([]*TreeNode, 0)
	if prev == nil {
		return next
	}
	for _, node := range prev {
		if node == nil {
			continue
		}
		if node.Left != nil {
			next = append(next, node.Left)
		}
		if node.Right != nil {
			next = append(next, node.Right)
		}
	}
	return next
}
