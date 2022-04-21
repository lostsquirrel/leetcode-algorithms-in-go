package main

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

func reverseKGroup(head *ListNode, k int) *ListNode {
	vHead := &ListNode{}
	var vTail *ListNode
	walker := head
	finished := false
	for !finished && walker != nil {
		index := k - 1
		group := make([]*ListNode, k)
		for index > -1 {
			if walker == nil {
				finished = true
				break
			}
			group[index] = walker

			walker = walker.Next
			index -= 1
		}
		if finished {
			vTail.Next = group[k-1]
		} else {
			for _, e := range group {
				if e == nil {
					continue
				}

				e.Next = nil
				if vHead.Next == nil {
					vHead.Next = e
					vTail = e
				} else {
					vTail.Next = e
					vTail = e
				}
			}
		}

	}
	return vHead.Next

}
