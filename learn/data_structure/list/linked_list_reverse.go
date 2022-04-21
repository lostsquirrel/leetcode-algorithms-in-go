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
	group := make([]*ListNode, k)
	for !finished && walker != nil {
		max_index := k - 1
		index := max_index

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
			if vHead.Next == nil {
				vTail = group[k-1]
			} else {
				vTail.Next = group[k-1]
			}

		} else {

			for index <= max_index-1 {
				index += 1
				e := group[index]
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
