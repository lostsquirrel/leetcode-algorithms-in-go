package tree

import "testing"

type IntPtrArrayInt struct {
	arr      []*int
	expected int
}

func TestMinDepth(t *testing.T) {

	tests := []IntPtrArrayInt{
		{[]*int{IntPtr(3), IntPtr(9), IntPtr(20), nil, nil, IntPtr(15), IntPtr(7)},
			2},
	}
	for _, test := range tests {
		tree := createTree(test.arr)
		isEqualInt(test.expected, minDepth(tree), t)
	}

}

func TestMinDepth02(t *testing.T) {
	tree := &TreeNode{
		Val: 2,
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 4,
				Right: &TreeNode{
					Val: 5,
					Right: &TreeNode{
						Val: 6,
					},
				},
			},
		},
	}
	isEqualInt(5, minDepth(tree), t)
}
