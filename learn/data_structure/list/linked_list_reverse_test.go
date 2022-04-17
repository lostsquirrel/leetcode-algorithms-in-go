package main

import (
	"fmt"
	"testing"
)

func TestLinkedListReverseEmpty(t *testing.T) {
	var c *ListNode
	r := reverseList(c)
	fmt.Println(r)
}

func TestLinkedListReverseOne(t *testing.T) {
	c := &ListNode{
		Val:  1,
		Next: nil,
	}
	r := reverseList(c)
	fmt.Println(r)

}
func TestLinkedListReverseTwo(t *testing.T) {
	c := ListNodeBuilder(1, 2)
	r := reverseList(c)
	PrintListNode(r)

}
func TestLinkedListReverseFive(t *testing.T) {
	c := ListNodeBuilder(1, 2, 3, 4, 5)
	r := reverseList(c)
	PrintListNode(r)

}

func TestListNodeBuilder(t *testing.T) {
	var head *ListNode = ListNodeBuilder(1, 2, 3)
	PrintListNode(head)
}
