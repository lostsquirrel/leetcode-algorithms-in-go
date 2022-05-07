package tree

func maxDepth(root *TreeNode) int {
	return walker(root, 0)
}

func walker(node *TreeNode, depth int) int {
	if node == nil {
		return depth
	}
	left := walker(node.Left, depth+1)
	right := walker(node.Right, depth+1)
	if left > right {
		return left
	}
	return right
}
