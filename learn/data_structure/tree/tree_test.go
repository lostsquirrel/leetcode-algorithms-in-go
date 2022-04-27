package tree

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPreOrderTraversal(t *testing.T) {
	c := createTestNode()
	result := preOrderTraversal(&c)
	fmt.Println(result)
	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(expected, result) {
		t.Fail()
	}
}

func createTestNode() TreeNode {
	a := TreeNode{3, nil, nil}
	b := TreeNode{2, &a, nil}
	c := TreeNode{1, nil, &b}
	return c
}

func TestInOrderTraversal(t *testing.T) {
	c := createTestNode()
	result := inOrderTraversal(&c)
	fmt.Println(result)
	expected := []int{1, 3, 2}
	if !reflect.DeepEqual(expected, result) {
		t.Fail()
	}
}

func TestPostOrderTraversal(t *testing.T) {
	c := createTestNode()
	result := postOrderTraversal(&c)
	fmt.Println(result)
	expected := []int{3, 2, 1}
	if !reflect.DeepEqual(expected, result) {
		t.Fail()
	}
}
