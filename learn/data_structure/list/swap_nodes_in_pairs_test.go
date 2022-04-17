package main

import (
	"fmt"
	"testing"
)

func TestSwapEmpty(t *testing.T) {
	var head *ListNode
	r := swapPairs(head)
	if r != nil {
		t.Errorf("got %v, wanted %v", r, nil)
	}
}
func TestSwapOne(t *testing.T) {
	var head *ListNode = &ListNode{Val: 1}
	r := swapPairs(head)
	if r.Val != 1 {
		t.Errorf("got %v, wanted %v", r, 1)
	}
}
func TestSwapTwo(t *testing.T) {
	var head *ListNode = ListNodeBuilder(1, 2)
	r := swapPairs(head)
	if r.Val != 2 {
		t.Errorf("got %v, wanted %v", r, 2)
	}
}
func TestSwap3(t *testing.T) {
	var head *ListNode = ListNodeBuilder(1, 2, 3)
	r := swapPairs(head)

	ValidateListNode(t, r, 2, 1, 3)
}
func TestSwap4(t *testing.T) {
	var head *ListNode = ListNodeBuilder(1, 2, 3, 4)
	r := swapPairs(head)

	ValidateListNode(t, r, 2, 1, 4, 3)
}
func TestSwap5(t *testing.T) {
	var head *ListNode = ListNodeBuilder(1, 2, 3, 4, 5)
	r := swapPairs(head)

	ValidateListNode(t, r, 2, 1, 4, 3, 5)
}

func ValidateListNode(t *testing.T, head *ListNode, expected ...int) {
	walker := head
	for i, n := range expected {
		if walker.Val != n {
			t.Errorf("got %v at %d, wanted %v", walker, i, 2)
		}
		fmt.Printf(" %d, %d, %d\n", i, walker.Val, n)
		walker = walker.Next
	}
}
