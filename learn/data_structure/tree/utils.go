package tree

import "testing"

func createTree(s []*int) *TreeNode {

	nodes := make([]*TreeNode, len(s))
	for i, v := range s {
		if v != nil {
			vn := &TreeNode{
				Val: *v,
			}
			nodes[i] = vn
			if i > 0 {
				j := i + 1
				parent := j/2 - 1
				if j%2 == 0 {

					nodes[parent].Left = vn
				} else {
					nodes[parent].Right = vn
				}
			}
		}
	}
	return nodes[0]
}

func IntPtr(x int) *int {
	return &x
}

func isFalse(r bool, t *testing.T) {
	if r {
		t.Error("expected false get true")
	}
}

func isEqualInt(expected, actual int, t *testing.T) {
	if expected != actual {
		t.Errorf("expected %d got %d", expected, actual)
	}
}
