package main

import "fmt"

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
