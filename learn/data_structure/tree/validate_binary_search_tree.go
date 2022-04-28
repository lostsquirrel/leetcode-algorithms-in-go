package tree

func isValidBST(root *TreeNode) bool {
	isValid, _ := isValid(root)
	return isValid
}

func isValid(root *TreeNode) (bool, []int) {
	if root == nil {
		return true, []int{}
	}
	var c []int
	c = append(c, root.Val)
	if root.Left != nil {
		isValidLeft, leftChildren := isValid(root.Left)
		if !isValidLeft {
			return false, []int{}
		}
		for _, child := range leftChildren {
			if child >= root.Val {
				return false, []int{}
			}
		}
		c = append(c, leftChildren...)
	}
	if root.Right != nil {
		isValidRight, rightChildren := isValid(root.Right)
		if !isValidRight {
			return false, []int{}
		}
		for _, child := range rightChildren {
			if child <= root.Val {
				return false, []int{}
			}
		}
		c = append(c, rightChildren...)
	}

	return true, c
}
