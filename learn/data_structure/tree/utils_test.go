package tree

import (
	"fmt"
	"testing"
)

func TestCreateTree(t *testing.T) {
	root := []*int{IntPtr(3), IntPtr(9), IntPtr(20), nil, nil, IntPtr(15), IntPtr(7)}
	tree := createTree(root)
	fmt.Println(levelOrder2(tree))
}
func TestCreateTree02(t *testing.T) {
	root := []*int{IntPtr(1), nil, IntPtr(2)}
	tree := createTree(root)
	fmt.Println(levelOrder2(tree))
}
