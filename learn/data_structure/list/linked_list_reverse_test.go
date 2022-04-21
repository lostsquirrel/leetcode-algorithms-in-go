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
func TestLinkedListReverseKGroupEmpty(t *testing.T) {
	var c *ListNode
	r := reverseKGroup(c, 2)
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
func TestLinkedListReverseKGroupOne(t *testing.T) {
	c := &ListNode{
		Val:  1,
		Next: nil,
	}
	r := reverseKGroup(c, 2)
	fmt.Println(r)
}

func TestLinkedListReverseTwo(t *testing.T) {
	c := ListNodeBuilder(1, 2)
	r := reverseList(c)
	PrintListNode(r)
}
func TestLinkedListReverseKGrouTwo(t *testing.T) {
	c := ListNodeBuilder(1, 2)
	r := reverseKGroup(c, 2)
	PrintListNode(r)
}
func TestLinkedListReverseFive(t *testing.T) {
	c := ListNodeBuilder(1, 2, 3, 4, 5)
	r := reverseList(c)
	PrintListNode(r)

}
func TestLinkedListReverseKGroupFive(t *testing.T) {
	c := ListNodeBuilder(1, 2, 3, 4, 5)
	r := reverseKGroup(c, 2)
	PrintListNode(r)

}
func TestLinkedListReverseKGroupFiveAndThree(t *testing.T) {
	c := ListNodeBuilder(1, 2, 3, 4, 5)
	r := reverseKGroup(c, 3)
	PrintListNode(r)

}

func TestListNodeBuilder(t *testing.T) {
	var head *ListNode = ListNodeBuilder(1, 2, 3)
	PrintListNode(head)
}
