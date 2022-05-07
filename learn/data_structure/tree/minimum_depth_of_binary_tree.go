package tree

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	s := []*TreeNode{root}
	for i := 1; ; i++ {
		s = walkLevel(s)
		if len(s) == 0 {
			return i
		}
		if hasLeaf(s) {
			return i + 1
		}
	}
}

func walkLevel(prev []*TreeNode) []*TreeNode {
	next := make([]*TreeNode, 0)
	for _, node := range prev {
		if node != nil {
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}

		}
	}
	return next
}

func hasLeaf(level []*TreeNode) bool {
	for _, node := range level {
		if node.Left == nil && node.Right == nil {
			return true
		}
	}
	return false
}
