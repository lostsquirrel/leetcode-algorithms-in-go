package tree

import "testing"

func TestValidate(t *testing.T) {
	tree := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	check(!isValidBST(tree), t)
	tree = &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 6,
			},
		},
	}
	check(isValidBST(tree), t)

}

func TestValidate2(t *testing.T) {

	tree := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
		},
		Right: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	check(isValidBST(tree), t)
}

func check(r bool, t *testing.T) {
	if r {
		t.Error("expected false get true")
	}
}
