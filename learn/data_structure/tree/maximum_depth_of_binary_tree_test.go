package tree

import "testing"

func TestDepath(t *testing.T) {
	root := []*int{IntPtr(3), IntPtr(9), IntPtr(20), nil, nil, IntPtr(15), IntPtr(7)}
	tree := createTree(root)
	depth := maxDepth(tree)
	isEqualInt(3, depth, t)
}
func TestDepath02(t *testing.T) {
	root := []*int{IntPtr(1), nil, IntPtr(2)}
	tree := createTree(root)
	depth := maxDepth(tree)
	isEqualInt(2, depth, t)
}
