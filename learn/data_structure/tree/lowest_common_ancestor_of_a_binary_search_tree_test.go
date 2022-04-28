package tree

import "testing"

func TestLowestCommonAncestor(t *testing.T) {
	p := &TreeNode{
		Val: 3,
	}
	q := &TreeNode{
		Val: 5,
	}
	tree := &TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val:   4,
				Left:  p,
				Right: q,
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 7,
			},
			Right: &TreeNode{
				Val: 9,
			},
		},
	}

	if lowestCommonAncestor(tree, p, q).Val != 4 {
		t.Error("error")
	}
}
