package tree

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	walker := root
	for walker != nil {

		if walker.Val > p.Val && walker.Val > q.Val {
			walker = walker.Left
		} else if walker.Val < p.Val && walker.Val < q.Val {
			walker = walker.Right
		} else {
			return walker
		}
	}
	return walker

}
