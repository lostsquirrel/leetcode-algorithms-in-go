package main

func hasCycle(head *ListNode) bool {
	m := make(map[*ListNode]bool)
	walker := head
	for walker != nil {
		_, hasKey := m[walker]
		if hasKey {
			return true
		} else {
			m[walker] = true
			walker = walker.Next

		}
	}
	return false
}

// 快慢指针

func hasCycle2(head *ListNode) bool {
	// walker := head
	fast := head
	slow := head
	for fast != nil {
		fast = fast.Next
		if fast == nil {
			return false
		}
		fast = fast.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

func detectCycle(head *ListNode) *ListNode {
	m := make(map[*ListNode]bool)

	walker := head
	for walker != nil {
		_, hasKey := m[walker]
		if hasKey {
			return walker
		} else {
			m[walker] = true
			walker = walker.Next
		}
	}
	return nil
}
