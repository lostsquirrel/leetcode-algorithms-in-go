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

func detectCycle(head *ListNode) *ListNode {
	m := make(map[*ListNode]int)
	var pw *ListNode
	walker := head
	index := 0
	for walker != nil {
		i, hasKey := m[walker]
		if hasKey {
			return pw
		} else {
			m[walker] = index
			pw = walker
			walker = walker.Next
			index += 1
		}
	}
	return nil
}
