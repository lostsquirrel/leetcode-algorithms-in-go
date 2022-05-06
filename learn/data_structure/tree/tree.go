package tree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preOrderTraversal(root *TreeNode) []int {
	r := make([]int, 0)
	if root != nil {

		r = append(r, root.Val)
		if root.Left != nil {
			r = append(r, preOrderTraversal(root.Left)...)
		}
		if root.Right != nil {
			r = append(r, preOrderTraversal(root.Right)...)
		}
	}
	return r
}

func inOrderTraversal(root *TreeNode) []int {
	r := make([]int, 0)
	inOrderTraversalHelp(root, &r)
	return r
}

func inOrderTraversalHelp(root *TreeNode, r *[]int) {
	if root != nil {
		inOrderTraversalHelp(root.Left, r)
		*r = append((*r), root.Val)
		inOrderTraversalHelp(root.Right, r)
	}
}

func postOrderTraversal(root *TreeNode) []int {
	r := make([]int, 0)
	postOrderTraversalHelper(root, &r)
	return r
}

func postOrderTraversalHelper(root *TreeNode, r *[]int) {
	if root != nil {
		postOrderTraversalHelper(root.Left, r)
		postOrderTraversalHelper(root.Right, r)
		*r = append((*r), root.Val)
	}
}

func levelOrder(root *TreeNode) [][]int {
	a := make([][]int, 0)
	queue := make(chan int, 10)
	queue <- root.Val
	x := <-queue
	for x > 0 {
		x = <-queue
	}
	return a
}

func level(queue chan int) ([]int, chan int) {
	row := make([]int, 0)
	c := make(chan int, 10)
	return row, c
}
