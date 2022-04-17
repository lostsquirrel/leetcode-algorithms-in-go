package main

func swapPairs(head *ListNode) *ListNode {

	return _swap(head)

}

func _swap(head *ListNode) (left *ListNode) {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	l := head
	r := head.Next
	rright := r.Next
	r.Next = head
	l.Next = _swap(rright)
	return r
}
