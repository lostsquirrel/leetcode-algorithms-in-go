package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {

	var prev *ListNode
	// var next *ListNode
	walker := head
	for walker != nil {
		next := walker.Next
		t := walker
		t.Next = prev
		prev = t
		walker = next
	}
	return prev
}

func PrintListNode(head *ListNode) {
	walker := head
	for walker != nil {
		fmt.Println(walker.Val)
		walker = walker.Next
	}
}

func ListNodeBuilder(data ...int) *ListNode {
	var node *ListNode
	var head *ListNode
	for i, n := range data {
		n := &ListNode{
			Val: n,
		}
		if i == 0 {
			head = n
			node = n
		} else {
			node.Next = n
			node = n
		}
	}
	return head
}
