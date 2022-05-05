package tree

import (
	"fmt"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	tree := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	r := levelOrder2(tree)
	fmt.Println(r)
}

func TestLevelOrder2(t *testing.T) {
	tree := &TreeNode{}
	r := levelOrder2(tree)
	fmt.Println(r)
}
func TestLevelOrder3(t *testing.T) {
	var tree *TreeNode
	r := levelOrder2(tree)
	fmt.Println(r)
}
