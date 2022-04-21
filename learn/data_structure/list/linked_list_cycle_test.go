package main

import "testing"

func TestHasCycle2(t *testing.T) {
	head := &ListNode{}
	if hasCycle2(head) {
		t.Error("null has no loop")
	}
}

func TestHasCycle2WithLoop(t *testing.T) {
	c := &ListNode{}
	b := &ListNode{
		Next: c,
	}
	head := &ListNode{
		Next: b,
	}

	c.Next = head
	if !hasCycle2(head) {
		t.Error("should has loop")
	}
}
